package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"github.com/h3mmy/bloopyboi/bot/services"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	textResponseMap = map[string]string{
		"pong": "Ping!",
		"Pong!": "-_-",
		"!bliss": "I use slash commands now. Try using /bliss",
	}
)

type MessageChanBlooper struct {
	msgCreateChan *chan *discordgo.MessageCreate
	msgReactAChan *chan *discordgo.MessageReactionAdd
	msgReactRChan *chan *discordgo.MessageReactionRemove
	msgSendChan   *chan *models.DiscordMessageSendRequest
	logger        *zap.Logger
	msgRegistry   map[string]*discordgo.Message
	inspiroSvc    *services.InspiroService
}

func NewMessageChanBlooper(
	insproSvc *services.InspiroService,
	createCh *chan *discordgo.MessageCreate,
	reactACh *chan *discordgo.MessageReactionAdd,
	reactRCh *chan *discordgo.MessageReactionRemove,
	msgSendChan *chan *models.DiscordMessageSendRequest,
) *MessageChanBlooper {

	lgr := log.NewZapLogger().With(zapcore.Field{
		Key:    HandlerLoggerFieldKey,
		Type:   zapcore.StringType,
		String: "messageChan",
	})

	return &MessageChanBlooper{
		inspiroSvc:    insproSvc,
		msgCreateChan: createCh,
		msgReactAChan: reactACh,
		msgReactRChan: reactRCh,
		msgSendChan:   msgSendChan,
		logger:        lgr,
		msgRegistry:   make(map[string]*discordgo.Message),
	}
}

func (mcb *MessageChanBlooper) Start(ctx context.Context) error {
	for {
		mcb.logger.Debug("Listening to channels")
		select {
		case msg := <-*mcb.msgCreateChan:
			mcb.logger.Debug("Received new msgC via channel")
			mcb.processIncomingMessage(msg)
			mcb.logger.Debug("Finished processIncomingMessage")
		case msgRAdd := <-*mcb.msgReactAChan:
			mcb.logger.Debug("Received new reactionA via channel")
			mcb.processReactionAdd(msgRAdd)
			mcb.logger.Debug("Finished processReactionAdd")
		case msgRMinus := <-*mcb.msgReactRChan:
			mcb.logger.Debug("Received new reactionA via channel")
			mcb.processReactionRemove(msgRMinus)
			mcb.logger.Debug("Finished processReactionRemove")
		case <-ctx.Done():
			mcb.logger.Info("Exit Received. Closing Channels")
			close(*mcb.msgCreateChan)
			close(*mcb.msgReactAChan)
			close(*mcb.msgReactRChan)
			return nil
		}
	}
}

func (mcb *MessageChanBlooper) processIncomingMessage(msg *discordgo.MessageCreate) {
	logger := mcb.logger.With(zapcore.Field{Key: "method", Type: zapcore.StringType, String: "processIncomingMessage"})
	mcb.logger.Debug(fmt.Sprintf("processing new message with ID %s from user %s", msg.ID, msg.Author.Username))

	// Check for test message
	if msg.Content == "test reaction thingy" {
		logger.Debug(fmt.Sprintf("Adding msg %s from user %s to registry", msg.ID, msg.Author.Username))
		mcb.msgRegistry[msg.ID] = msg.Message
	}

	// Check for inspiro request
	if strings.ToLower(msg.Content) == "inspire" {
		logger.Debug(
			fmt.Sprintf(
				"Received Inspiration Request from %s with ID %s",
				msg.Author.Username,
				msg.Author.ID),
		)

		bttp := mcb.inspiroSvc
		embed := bttp.CreateInsprioEmbed()
		inspRes := &models.DiscordMessageSendRequest{
			ChannelID: msg.ChannelID,
			MessageComplex: &discordgo.MessageSend{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		}
		*mcb.msgSendChan <- inspRes
	}

	resp, ok := textResponseMap[msg.Content]
	if !ok {
		// Means nothing stored for canned Response
		return
	}
	logger.Debug(
		fmt.Sprintf(
			"Received Message from %s with ID %s",
			msg.Author.Username,
			msg.Author.ID),
	)

	*mcb.msgSendChan <- &models.DiscordMessageSendRequest{
		ChannelID: msg.ChannelID,
		MessageComplex: &discordgo.MessageSend{
			Content: resp,
			Reference: &discordgo.MessageReference{
				MessageID: msg.ID,
				ChannelID: msg.ChannelID,
				GuildID:   msg.GuildID,
			},
		},
	}
}

func (mcb *MessageChanBlooper) processReactionAdd(msgRAdd *discordgo.MessageReactionAdd) {
	mcb.logger.Debug(fmt.Sprintf("processing new reaction add on messageID %s with Emoji %v",
		msgRAdd.MessageID,
		msgRAdd.Emoji))

	if smsg, ok := mcb.msgRegistry[msgRAdd.MessageID]; ok {
		mcb.logger.Debug(fmt.Sprintf("found message id %s in registry", smsg.ID))
	}
}

func (mcb *MessageChanBlooper) processReactionRemove(msgRMinus *discordgo.MessageReactionRemove) {
	mcb.logger.Debug(fmt.Sprintf("processing new reaction add on messageID %s with Emoji %v",
		msgRMinus.MessageID,
		msgRMinus.Emoji))

	if smsg, ok := mcb.msgRegistry[msgRMinus.MessageID]; ok {
		mcb.logger.Debug(fmt.Sprintf("found message id %s in registry", smsg.ID))
	}
}
