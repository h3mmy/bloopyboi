package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/cdfmlr/ellipsis"
	"github.com/h3mmy/bloopyboi/bot/internal/log"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"github.com/h3mmy/bloopyboi/bot/services"
	"go.uber.org/zap"
)

type BookCommand struct {
	meta        models.BloopyMeta
	Name        string
	Description string
	logger      *zap.Logger
	bookSvc     *services.BookService
}

func NewBookCommand(bookSvc *services.BookService) *BookCommand {
	return &BookCommand{
		meta:        models.NewBloopyMeta(),
		Name:        "book",
		Description: "(Xperimental) Get book info",
		bookSvc:     bookSvc,
		logger:      log.NewZapLogger().Named("book_command"),
	}
}

func (b *BookCommand) GetAppCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        strings.ToLower(b.Name),
		Description: b.Description,
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "isbn",
				Description: "The ISBN of the book",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "title",
				Description: "The title of the book",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "author",
				Description: "The author of the book",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "publisher",
				Description: "The publisher of the book",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "year",
				Description: "The year of the book",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "text_snippet",
				Description: "Any text snippet to search for",
				Required:    false,
			},
		},
	}
}

func (b *BookCommand) GetAppCommandHandler() func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		var isbn, title, author, publisher, textSnippet string
		for _, opt := range options {
			if opt.Name == "isbn" {
				isbn = opt.StringValue()
			}
			if opt.Name == "title" {
				title = opt.StringValue()
			}
			if opt.Name == "author" {
				author = opt.StringValue()
			}
			if opt.Name == "publisher" {
				publisher = opt.StringValue()
			}
			if opt.Name == "text_snippet" {
				textSnippet = opt.StringValue()
			}
		}
		booksvc := b.bookSvc
		volumes, err := booksvc.SearchBook(context.TODO(), &models.BookSearchRequest{
			ISBN:        isbn,
			Title:       title,
			Author:      author,
			Publisher:   publisher,
			TextSnippet: textSnippet,
		})
		if err != nil {
			b.logger.Error("error searching for book", zap.Error(err))
			err = s.InteractionRespond(i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "I was getting the books, and I don't know what happened. Sorry I failed you. I wrote the reason down in the logs",
					},
				},
			)
			if err != nil {
				b.logger.Error("error responding to interaction", zap.Error(err))
			}
		} else {
			buttonOpts := []discordgo.MessageComponent{}
			selectOpts := []discordgo.SelectMenuOption{}

			for _, volume := range volumes.Items {
				var blabel string
				etal := len(volume.VolumeInfo.Authors)
				switch {
				case etal == 0:
					blabel = fmt.Sprintf("%s by Unknown", volume.VolumeInfo.Title)
				case etal >= 2:
					blabel = fmt.Sprintf("%s by %s, %s et al.", volume.VolumeInfo.Title, volume.VolumeInfo.Authors[0], volume.VolumeInfo.Authors[1])
				default:
					blabel = fmt.Sprintf("%s by %s", volume.VolumeInfo.Title, volume.VolumeInfo.Authors[0])
				}
				blabel = ellipsis.Ending(blabel, 100)
				buttonOpts = append(buttonOpts,
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.Button{
								Label: blabel,
								Style: discordgo.PrimaryButton,
								Emoji: discordgo.ComponentEmoji{
									Name: "📖",
								},
								Disabled: false,
								CustomID: fmt.Sprintf("select_book_%s", volume.Id),
							}},
					},
				)
				selectOpts = append(selectOpts,
					discordgo.SelectMenuOption{
						Value: volume.Id,
						Emoji: discordgo.ComponentEmoji{
							Name: "📖",
						},
						Label:       blabel,
						Description: volume.VolumeInfo.Publisher,
						Default:     false,
					})
			}
			selectMenu := discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.SelectMenu{
						CustomID:    "select_book",
						Placeholder: "Select an option",
						Options:     selectOpts,
					},
				},
			}
			b.logger.Debug(fmt.Sprintf("made %d buttons", len(buttonOpts)))
			b.logger.Debug(fmt.Sprintf("made %d select menu options", len(selectOpts)))
			if len(selectOpts) == 0 {
				err = s.InteractionRespond(i.Interaction,
					&discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "I couldn't find any books. Sorry I failed you. I wrote the reason down in the logs",
						},
					},
				)
				if err != nil {
					b.logger.Error("error responding to interaction", zap.Error(err))
				}
				return
			}
			err = s.InteractionRespond(i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						CustomID: "select_book",
						Title:    "Book Search Results",
						Flags:    discordgo.MessageFlagsEphemeral,
						Components: []discordgo.MessageComponent{
							selectMenu,
						},
					},
				},
			)
			if err != nil {
				b.logger.Error("failed to respond to interaction", zap.Error(err), zap.String("command", "book"))
			}
		}
	}
}

func (b *BookCommand) GetMessageComponentHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"request_book": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			b.logger.Debug(fmt.Sprintf("received book request with %v", i.Data), zap.String(string(models.CtxKeyMessageID), i.Message.ID))
			ctx := context.WithValue(context.TODO(), models.CtxKeyMessageID, i.Message.ID)
			fields := i.Message.Embeds[0].Fields
			for _, field := range fields {
				if field.Name == "Volume ID" {
					b.logger.Info(fmt.Sprintf("Received Request for volumeId %s", field.Value))
					var discordUser *discordgo.User
					if i.User != nil {
						discordUser = i.User
					} else {
						discordUser = i.Member.User
					}
					err := b.bookSvc.SubmitBookRequest(ctx, discordUser, field.Value)
					if err != nil {
						b.logger.Error("error submitting book request", zap.Error(err))
					}
				}
			}
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content: "Request Received. My request box doesn't have a bottom yet so they do tend to disappear. Still working on that",
					Embeds:  i.Message.Embeds,
				},
			})
			if err != nil {
				b.logger.Error("error responding to book request", zap.Error(err))
			}
		},
		"ignore_book": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content: "No Worries. I'll go read it myself",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
			if err != nil {
				b.logger.Error("error responding to book ignore")
			}
			_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
				Content: "lol, jk. I can't read!",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
			)
			if err != nil {
				b.logger.Error("error with follow up to book ignore")
			}
		},
		"select_book": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var response *discordgo.InteractionResponse

			data := i.MessageComponentData()
			selectionInfo := data.Values[0]
			b.logger.Debug(fmt.Sprintf("messageData: %v", data))
			selectedVol, err := b.bookSvc.GetVolume(selectionInfo)
			if err != nil {
				b.logger.Error("error getting volume", zap.String("volume", selectionInfo))
			}
			volumeEmbed := &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: selectedVol.VolumeInfo.ImageLinks.Thumbnail,
				},
				Title: fmt.Sprintf("%s by %s", selectedVol.VolumeInfo.Title, strings.Join(selectedVol.VolumeInfo.Authors, "")),
				URL:   selectedVol.VolumeInfo.InfoLink,

				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "Publisher",
						Value: selectedVol.VolumeInfo.Publisher,
					},
					{
						Name:  "Published Date",
						Value: selectedVol.VolumeInfo.PublishedDate,
					},
					{
						Name:  "Volume ID",
						Value: selectionInfo,
					},
				},
			}
			for _, identifier := range selectedVol.VolumeInfo.IndustryIdentifiers {
				volumeEmbed.Fields = append(volumeEmbed.Fields, &discordgo.MessageEmbedField{
					Name:  identifier.Type,
					Value: identifier.Identifier,
				})
			}

			response = &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("You selected %s by %s", selectedVol.VolumeInfo.Title, strings.Join(selectedVol.VolumeInfo.Authors, "")),
					Embeds: []*discordgo.MessageEmbed{
						volumeEmbed,
					},
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Label:    "Request",
									Style:    discordgo.SuccessButton,
									CustomID: "request_book",
									Emoji: discordgo.ComponentEmoji{
										Name: "✅",
									},
								},
								discordgo.Button{
									Label:    "Ignore",
									Style:    discordgo.SecondaryButton,
									CustomID: "ignore_book",
									Emoji: discordgo.ComponentEmoji{
										Name: "⭕",
									},
								},
							},
						},
					},
					Flags: discordgo.MessageFlagsEphemeral,
				},
			}

			err = s.InteractionRespond(i.Interaction, response)
			if err != nil {
				b.logger.Error("failed responding with book selection", zap.Error(err))
			}
		},
	}
}
