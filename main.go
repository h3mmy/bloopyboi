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
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gitlab.com/h3mmy/bloopyboi/bot/discord"
	bloopyCommands "gitlab.com/h3mmy/bloopyboi/bot/discord/commands"
	"gitlab.com/h3mmy/bloopyboi/bot/providers"
	"golang.org/x/sync/errgroup"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"

	"github.com/alexliesenfeld/health"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	botLogFieldKey = "bot"
)

// Variables
var (
	Token              string
	RegisteredCommands []*discordgo.ApplicationCommand
	RemoveCommands     = true
	s *discordgo.Session

)

func init() {
	viper.SetConfigName("config")           // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/config")          // path to look for the config file in
	viper.AddConfigPath("$HOME/.bloopyboi") // call multiple times to add many search paths
	viper.AddConfigPath(".")                // optionally look for config in the working directory
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(errors.New("Fatal error config file: " + err.Error()))
	}
}

// Register functions as callbacks for varios signatures
func addHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := bloopyCommands.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
	s.AddHandler(bloopyCommands.MessageCreate)
	s.AddHandler(bloopyCommands.DirectMessageCreate)
}

// Where the magic happens
func main() {

	ctx := signals.SetupSignalHandler()
	ctx, cancelCtxFn := context.WithCancel(ctx)

	defer cancelCtxFn()

	errGroup, ctx := errgroup.WithContext(ctx)

	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: false}

	commonLogger := logger.WithField("common", "group")

	discordClient, err := discord.NewDiscordClient(commonLogger.WithField(botLogFieldKey, "Discord"))
	if err != nil {
		logger.Panicf("Error Creating Discord Client %v", err)
		return
	}

	errGroup.Go(func() error {
		return discordClient.Start(ctx)
	})

	ctx = context.WithValue(ctx, "customReady", "true")

	readinessChecker := providers.NewReadinessChecker()

	// Liveness check should mostly contain checks that identify if the service is locked up or in a state that it
	// cannot recover from (deadlocks, etc.). In most cases it should just respond with 200 OK to avoid unexpected
	// restarts.
	livenessChecker := health.NewChecker()

	// Create a new health check http.Handler that returns the health status
	// serialized as a JSON string. You can pass pass further configuration
	// options to NewHandler to modify default configuration.
	http.Handle("/healthz", health.NewHandler(livenessChecker))
	http.Handle("/ready", health.NewHandler(readinessChecker))

	// Start the HTTP server
	log.Fatalln(http.ListenAndServe(":3000", nil))

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")


	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-stop

	// Cleanly close down the Discord session.
	defer s.Close()

}
