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

	bloopyCommands "gitlab.com/h3mmy/bloopyboi/bot/commands"
	"gitlab.com/h3mmy/bloopyboi/bot/providers"
	"gitlab.com/h3mmy/bloopyboi/bot/util"

	"github.com/alexliesenfeld/health"
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
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
	s.AddHandler(messageCreate)
	s.AddHandler(bloopyCommands.DirectMessageCreate)
}

// Where the magic happens
func main() {

	ctx := context.Background()

	// Get token
	Token = providers.GetBotToken()

	// Create a new Discord session using the provided bot token.
	s, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Just like the ping pong example, we only care about receiving message
	// events in this example.
	s.Identify.Intents = discordgo.IntentsGuildMessages

	addHandlers(s)

	// Open a websocket connection to Discord and begin listening.
	err = s.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	ctx = context.WithValue(ctx, "customReady", "true")


	RegisteredCommands = make([]*discordgo.ApplicationCommand, len(bloopyCommands.Commands))
	for i, v := range bloopyCommands.Commands {
		// Leaving GuildId empty
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		RegisteredCommands[i] = cmd
	}

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

	if RemoveCommands {
		log.Println("Removing commands...")
		// // We need to fetch the commands, since deleting requires the command ID.
		// // We are doing this from the returned commands on line 375, because using
		// // this will delete all the commands, which might not be desirable, so we
		// // are deleting only the commands that we added.
		// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
		// if err != nil {
		// 	log.Fatalf("Could not fetch registered commands: %v", err)
		// }

		for _, v := range RegisteredCommands {
			err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}
}

// Temporary method to check if command is relevant
func isNotDMCommand(q string) bool {
	switch q {
	case "ping":
		return false
	}
	return true
}

// Temporary method to check if command is relevant
func isNotChannelCommand(q string) bool {
	switch q {
	case "inspire":
		return false
	}
	return true
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == "inspire" {
		bttp := util.NewBloopyClient()
		embed := &discordgo.MessageEmbed{
			Author: &discordgo.MessageEmbedAuthor{},
			Image: &discordgo.MessageEmbedImage{
				URL: bttp.Inspiro_api.GetInspiro(),
			},
		}
		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}

	if m.Content == "Pong!" {
		s.ChannelMessageSend(m.ChannelID, "-_-")
	}
}
