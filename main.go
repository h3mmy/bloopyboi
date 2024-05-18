// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

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

// Where the magic happens
func main() {

	ctx := signals.SetupSignalHandler()
	ctx, cancelCtxFn := context.WithCancel(ctx)

	go func() {
		sCh := make(chan os.Signal, 3)
		signal.Notify(sCh, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		<-sCh
		cancelCtxFn()
	}()

	logger := providers.NewZapLogger()
	commonLogger := logger.With(zapcore.Field{
		Key:    "group",
		Type:   zapcore.StringType,
		String: "common",
	})

	boi := bot.New()
	boi.WithLogger(commonLogger.With(zapcore.Field{
		Key:    botLogFieldKey,
		Type:   zapcore.StringType,
		String: "BloopyBoi",
	}))

	errGroup, ctx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		return boi.Run(ctx)
	})

	gateway := bot.NewDefaultGateway().WithBotInstance(boi)

	// Start server
	errGroup.Go(func() error {
		if err := gateway.Start(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

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

	readinessChecker := boi.GetReadinessChecker()

	// Create a new health check http.Handler that returns the health status
	// serialized as a JSON string. You can pass pass further configuration
	// options to NewHandler to modify default configuration.
	http.Handle("/healthz", health.NewHandler(livenessChecker))
	http.Handle("/ready", health.NewHandler(readinessChecker))

	// Start the HTTP server
	log.Fatalln(http.ListenAndServe(":3000", nil))

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")

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
