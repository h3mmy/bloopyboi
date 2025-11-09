package services

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/h3mmy/bloopyboi/pkg/logs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InspiroService is a service that provides inspirational images.
type InspiroService struct {
	inspiroClient *InspiroClient
	logger        *zap.Logger
}

// NewInspiroService creates a new InspiroService.
func NewInspiroService(inspiro *InspiroClient) *InspiroService {
	lgr := log.NewZapLogger().With(zapcore.Field{
		Key:    ServiceLoggerFieldKey,
		Type:   zapcore.StringType,
		String: "inspiro",
	})
	return &InspiroService{
		inspiroClient: inspiro,
		logger:        lgr,
	}
}

// GetInspiroImageURL returns the URL of an inspirational image.
func (ic *InspiroService) GetInspiroImageURL() string {
	ic.logger.Debug("Getting Inspiro Image")
	return ic.inspiroClient.GetInspiro()
}

// GetClient returns the InspiroClient.
func (ic *InspiroService) GetClient() *InspiroClient {
	return ic.inspiroClient
}

// CreateInsprioEmbed creates a Discord embed for an inspirational image.
func (ic *InspiroService) CreateInsprioEmbed() *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{},
		Image: &discordgo.MessageEmbedImage{
			URL: ic.GetInspiroImageURL(),
		},
	}
}

// CreateInteractionResponseData creates a Discord interaction response data for an inspirational image.
func (ic *InspiroService) CreateInteractionResponseData() *discordgo.InteractionResponseData {
	return &discordgo.InteractionResponseData{
		Embeds: []*discordgo.MessageEmbed{
			ic.CreateInsprioEmbed(),
		},
	}
}

// CreateInteractionResponse creates a Discord interaction response for an inspirational image.
func (ic *InspiroService) CreateInteractionResponse() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: ic.CreateInteractionResponseData(),
	}
}
