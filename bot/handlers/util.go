package handlers

// typeInChannel sets the typing indicator for a channel. The indicator is cleared
// when a message is sent.
// func typeInChannel(channel chan bool, s *discordgo.Session, channelID string) {
// 	for {
// 		select {
// 		case <-channel:
// 			return
// 		default:
// 			if err := s.ChannelTyping(channelID); err != nil {
// 				fmt.Println("unable to set typing indicator: ", err)
// 			}
// 			time.Sleep(time.Second * 5)
// 		}
// 	}
// }
