// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/h3mmy/bloopyboi/bot"
	"github.com/h3mmy/bloopyboi/bot/providers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	"github.com/alexliesenfeld/health"
)

const (
	botLogFieldKey = "bot"
)

// main is the entry point of the application.
// It sets up the signal handling, logger, bot, gateway, and health checks.
// It also starts the bot and gateway in separate goroutines.
func main() {
	// Set up a context that is canceled when the application receives a signal.
	ctx := signals.SetupSignalHandler()
	ctx, cancelCtxFn := context.WithCancel(ctx)

	// Start a goroutine to listen for signals and cancel the context when a signal is received.
	go func() {
		sCh := make(chan os.Signal, 3)
		signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sCh
		cancelCtxFn()
	}()

	// Create a new logger.
	logger := providers.NewZapLogger()
	commonLogger := logger.With(zapcore.Field{
		Key:    "group",
		Type:   zapcore.StringType,
		String: "common",
	})

	// Create a new bot instance.
	boi := bot.New()
	boi.WithLogger(commonLogger.With(zapcore.Field{
		Key:    botLogFieldKey,
		Type:   zapcore.StringType,
		String: "BloopyBoi",
	}))

	// Create a new errgroup to manage the bot's goroutines.
	errGroup, ctx := errgroup.WithContext(ctx)

	// Start the bot in a separate goroutine.
	errGroup.Go(func() error {
		return boi.Run(ctx)
	})

	// Create a new gateway and start it in a separate goroutine.
	gateway := bot.NewDefaultGateway().WithBotInstance(boi)

	// Start server
	errGroup.Go(func() error {
		if err := gateway.Start(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	// Create a new liveness checker.
	// Liveness check should mostly contain checks that identify if the service is locked up or in a state that it
	// cannot recover from (deadlocks, etc.). In most cases it should just respond with 200 OK to avoid unexpected
	// restarts.
	livenessChecker := health.NewChecker(
		health.WithCheck(health.Check{
			Name: "boi",
			Check: func(ctx context.Context) error {
				return boi.Ping(ctx)
			},
		}),
	)

	// Get the readiness checker from the bot.
	readinessChecker := boi.GetReadinessChecker()

	// Create a new health check http.Handler that returns the health status
	// serialized as a JSON string. You can pass pass further configuration
	// options to NewHandler to modify default configuration.
	http.Handle("/healthz", health.NewHandler(livenessChecker))
	http.Handle("/ready", health.NewHandler(readinessChecker))

	// Start the HTTP server.
	log.Fatalln(http.ListenAndServe(":3000", nil))

	// Wait for the bot to finish.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	// Wait for the errgroup to finish.
	err := errGroup.Wait()
	if err != nil {
		logger.Error("error", zapcore.Field{
			Key:    "error",
			Type:   zapcore.ErrorType,
			String: err.Error(),
		})
	} else {
		logger.Info("main exited")
	}

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := gateway.Shutdown(ctx); err != nil {
		commonLogger.Error("error shutting down gateway", zap.Error(err))
	}

}
