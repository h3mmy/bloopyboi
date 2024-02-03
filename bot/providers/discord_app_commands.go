package providers

import (
	"context"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
	"github.com/h3mmy/bloopyboi/bot/internal/models"
	"go.uber.org/zap"
	"github.com/cdfmlr/ellipsis"
)

var (
	AppCommands = []*discordgo.ApplicationCommand{
		{
			Name:        "inspire",
			Description: "Summons Inspiration",
		},
		{
			Name:        "blissfest",
			Description: "Gets blissfest related information",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "lineup",
					Description: "Try to fetch lineup image",
					Required:    true,
				},
			},
		},
		{
			Name:        "book",
			Description: "Actions relating to a book",
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
		},
	}
	AppCommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"inspire": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			bttp := GetInspiroClient()

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Author: &discordgo.MessageEmbedAuthor{},
							Image: &discordgo.MessageEmbedImage{
								URL: bttp.GetInspiroImageURL(),
							},
						},
					},
				},
			})
			if err != nil {
				logger.Error("Failed to respond to interaction", zap.Error(err), zap.String("command", "inspire"))
			}
		},
		"book": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
			booksvc := GetBookService()
			volumes, err := booksvc.SearchBook(context.TODO(), &models.BookSearchRequest{
				ISBN:        isbn,
				Title:       title,
				Author:      author,
				Publisher:   publisher,
				TextSnippet: textSnippet,
			})
			if err != nil {
				logger.Error("error searching for book", zap.Error(err))
				err = s.InteractionRespond(i.Interaction,
					&discordgo.InteractionResponse{
						Type: discordgo.InteractionResponseChannelMessageWithSource,
						Data: &discordgo.InteractionResponseData{
							Content: "I was getting the books, and I don't know what happened. Sorry I failed you. I wrote the reason down in the logs",
						},
					},
				)
				if err != nil {
					logger.Error("error responding to interaction", zap.Error(err))
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
				logger.Debug(fmt.Sprintf("made %d buttons", len(buttonOpts)))
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
					logger.Error("failed to respond to interaction", zap.Error(err), zap.String("command", "book"))
				}
			}
		},
		"blissfest": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			getLineUp := false
			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options
			for _, opt := range options {
				if opt.Name == "lineup" {
					getLineUp = opt.BoolValue()
				}
			}

			bsvc := GetBlissfestService()
			var resData discordgo.InteractionResponseData

			if getLineUp {
				resData = discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{
						{
							Author: &discordgo.MessageEmbedAuthor{},
							Image: &discordgo.MessageEmbedImage{
								URL: bsvc.GetLineupImageURI(),
							},
						},
					},
					Title: "Blissfest",
					// pending https://github.com/dustin/go-humanize/pull/92
					// Content: fmt.Sprintf("%s left", humanize.Time(bsvc.GetTimeUntilStart(nil))),
					Content: fmt.Sprintf("%s left", humanize.Time(*bsvc.GetStartTime())),
				}

			} else {
				resData = discordgo.InteractionResponseData{
					Title:   "Blissfest",
					Content: fmt.Sprintf("%s left", humanize.Time(*bsvc.GetStartTime())),
				}
			}

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &resData,
			})
			if err != nil {
				logger.Error("error responding to interaction", zap.Error(err))
			}
		},
	}
	MessageComponentHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"request_book": func(s *discordgo.Session, i *discordgo.InteractionCreate){
			logger.Debug(fmt.Sprintf("received book request with %v", i.Data), zap.String("message_id", i.Message.ID))
			fields := i.Message.Embeds[0].Fields
			for _, field := range fields {
				if field.Name == "Volume ID" {
					logger.Info(fmt.Sprintf("Received Request for volumeId %s", field.Value))
				}
			}
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content: "Request Received. My request box doesn't have a bottom yet so they do tend to disappear. Still working on that",
					Embeds: i.Message.Embeds,
				},
			})
			if err != nil {
				logger.Error("error responding to book request")
			}
		},
		"ignore_book": func(s *discordgo.Session, i *discordgo.InteractionCreate){

			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseUpdateMessage,
				Data: &discordgo.InteractionResponseData{
					Content: "No Worries. I'll go read it myself",
					Flags: discordgo.MessageFlagsEphemeral,
				},
			})
			if err != nil {
				logger.Error("error responding to book ignore")
			}
			_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
					Content: "lol, jk. I can't read!",
					Flags: discordgo.MessageFlagsEphemeral,
				},
			)
			if err != nil {
				logger.Error("error with follow up to book ignore")
			}
		},
		"select_book": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			var response *discordgo.InteractionResponse

			data := i.MessageComponentData()
			selectionInfo := data.Values[0]
			logger.Debug(fmt.Sprintf("messageData: %v", data))
			selectedVol, err := GetBookService().GetVolume(selectionInfo)
			if err != nil {
				logger.Error("error getting volume", zap.String("volume", selectionInfo))
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
						Name: "Volume ID",
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
				logger.Error("failed responding with book selection", zap.Error(err))
			}
		},
	}
	ModalSubmitHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){}
)
