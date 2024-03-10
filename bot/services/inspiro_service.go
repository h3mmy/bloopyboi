package services

import (
	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type InspiroService struct {
	inspiroClient *InspiroClient
	logger         *zap.Logger
}

// Creates New InspiroService with specified Service
func NewInspiroService(inspiro *InspiroClient) *InspiroService {
	lgr := log.NewZapLogger().With(zapcore.Field{
		Key:    ServiceLoggerFieldKey,
		Type:   zapcore.StringType,
		String: "inspiro",
	})
	return &InspiroService{
		inspiroClient: inspiro,
		logger:         lgr,
	}
}

func (ic *InspiroService) GetInspiroImageURL() string {
	ic.logger.Debug("Getting Inspiro Image")
	return ic.inspiroClient.GetInspiro()
}

func (ic *InspiroService) GetClient() *InspiroClient {
	return ic.inspiroClient
}

// Creates the discord.Embed object to be used in a message
func (ic *InspiroService) CreateInsprioEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Image: &discordgo.MessageEmbedImage{
			URL: ic.GetInspiroImageURL(),
		},
	}
}

// Creates InteractionResponseData with to be used with custom flags, etc
func (ic *InspiroService) CreateInteractionResponseData() *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{
			ic.CreateInsprioEmbed(),
		},
	}
}

// Creates InteractionResponse with default Type and Flags, etc.
func (ic *InspiroService) CreateInteractionResponse() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: ic.CreateInteractionResponseData(),
	}
}
