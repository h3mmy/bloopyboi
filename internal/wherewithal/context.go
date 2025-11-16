package wherewithal

import "github.com/bwmarrin/discordgo"

type MessageContext struct {
    TimeOfDay          string   // morning, afternoon, evening, late_night
    Tone               string   // excited, casual, frustrated, questioning, serious
    RecentTopics       []string // Topics from recent conversation
    ConversationLength int      // Number of recent messages
    Tags               []string // Combined context tags
}

type DiscordMessageContextAnalyzer interface {
    Analyze(msg *discordgo.Message) *MessageContext
}
