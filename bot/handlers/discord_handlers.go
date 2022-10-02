package handlers

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"gitlab.com/h3mmy/bloopyboi/bot/providers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MessageChanBlooper struct {
	msgCreateChan *chan *discordgo.MessageCreate
	msgReactAChan *chan *discordgo.MessageReactionAdd
	msgReactRChan *chan *discordgo.MessageReactionRemove
	logger        *zap.Logger
	msgRegistry   map[string]*discordgo.Message
}

func NewMessageChanBlooper(
	createCh *chan *discordgo.MessageCreate,
	reactACh *chan *discordgo.MessageReactionAdd,
	reactRCh *chan *discordgo.MessageReactionRemove,
) *MessageChanBlooper {

	lgr := providers.NewZapLogger().With(zapcore.Field{
		Key:    HandlerLoggerFieldKey,
		Type:   zapcore.StringType,
		String: "messageChan",
	})

	return &MessageChanBlooper{
		msgCreateChan: createCh,
		msgReactAChan: reactACh,
		msgReactRChan: reactRCh,
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
	mcb.logger.Debug(fmt.Sprintf("processing new message with ID %s from user %s", msg.ID, msg.Author.Username))
	if msg.Content == "test reaction thingy" {
		mcb.logger.Debug(fmt.Sprintf("Adding msg %s from user %s to registry", msg.ID, msg.Author.Username))
		mcb.msgRegistry[msg.ID] = msg.Message
	}
	return
}

func (mcb *MessageChanBlooper) processReactionAdd(msgRAdd *discordgo.MessageReactionAdd) {
	mcb.logger.Debug(fmt.Sprintf("processing new reaction add on messageID %s with Emoji %v",
		msgRAdd.MessageID,
		msgRAdd.Emoji))

	if smsg, ok := mcb.msgRegistry[msgRAdd.MessageID]; ok {
		mcb.logger.Debug(fmt.Sprintf("found message id %s in registry", smsg.ID))
	}
	return
}

func (mcb *MessageChanBlooper) processReactionRemove(msgRMinus *discordgo.MessageReactionRemove) {
	mcb.logger.Debug(fmt.Sprintf("processing new reaction add on messageID %s with Emoji %v",
		msgRMinus.MessageID,
		msgRMinus.Emoji))

	if smsg, ok := mcb.msgRegistry[msgRMinus.MessageID]; ok {
		mcb.logger.Debug(fmt.Sprintf("found message id %s in registry", smsg.ID))
	}
	return
}
