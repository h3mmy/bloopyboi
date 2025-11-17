package wherewithal

// This is only a rudimentary oversimplified example intended to inspire better ways to analyze context.

type ContextAnalyzer struct {
    recentMessages map[string][]*discordgo.MessageCreate // channel ID -> messages
}

func NewContextAnalyzer() *ContextAnalyzer {
    return &ContextAnalyzer{
        recentMessages: make(map[string][]*discordgo.MessageCreate),
    }
}

func NewContextAnalyzer() *ContextAnalyzer {
    return &ContextAnalyzer{
        recentMessages: make(map[string][]*discordgo.MessageCreate),
    }
}

func (c *ContextAnalyzer) Analyze(m *discordgo.MessageCreate) *MessageContext {
    ctx := &MessageContext{
        TimeOfDay:    c.getTimeOfDay(),
        Tone:         c.detectTone(m.Content),
        RecentTopics: c.extractTopics(m),
        Tags:         []string{},
    }
    
    // Track conversation
    c.trackMessage(m)
    ctx.ConversationLength = len(c.recentMessages[m.ChannelID])
    
    // Build tags
    ctx.Tags = append(ctx.Tags, ctx.TimeOfDay, ctx.Tone)
    ctx.Tags = append(ctx.Tags, ctx.RecentTopics...)
    
    return ctx
}

func (c *ContextAnalyzer) getTimeOfDay() string {
    hour := time.Now().Hour()
    
    switch {
    case hour >= 0 && hour < 6:
        return "late_night"
    case hour >= 6 && hour < 12:
        return "morning"
    case hour >= 12 && hour < 18:
        return "afternoon"
    case hour >= 18 && hour < 22:
        return "evening"
    default:
        return "late_night"
    }
}

func (c *ContextAnalyzer) detectTone(content string) string {
    contentLower := strings.ToLower(content)
    
    // Excited indicators
    excitedPatterns := []string{"!", "!!", "omg", "hype", "let's go", "amazing", "awesome", "wow"}
    for _, pattern := range excitedPatterns {
        if strings.Contains(contentLower, pattern) {
            return "excited"
        }
    }
    
    // Frustrated indicators
    frustratedPatterns := []string{"ugh", "wtf", "annoying", "broken", "doesn't work", "hate"}
    for _, pattern := range frustratedPatterns {
        if strings.Contains(contentLower, pattern) {
            return "frustrated"
        }
    }
    
    // Questioning indicators
    if strings.Contains(content, "?") || strings.HasPrefix(contentLower, "how") || 
       strings.HasPrefix(contentLower, "why") || strings.HasPrefix(contentLower, "what") {
        return "questioning"
    }
    
    // Default to casual
    return "casual"
}

func (c *ContextAnalyzer) extractTopics(m *discordgo.MessageCreate) []string {
    var topics []string
    contentLower := strings.ToLower(m.Content)
    
    // Define topic keywords
    topicMap := map[string][]string{
        "gaming":  {"game", "gaming", "play", "steam", "xbox", "ps5"},
        "music":   {"song", "music", "album", "concert", "band", "listen"},
        "event":   {"event", "festival", "blissfest", "concert", "party"},
        "tech":    {"code", "programming", "bug", "server", "kubernetes", "deploy"},
        "food":    {"food", "eat", "dinner", "lunch", "hungry", "restaurant"},
        "workout": {"gym", "workout", "exercise", "run", "fitness"},
    }
    
    for topic, keywords := range topicMap {
        for _, keyword := range keywords {
            if strings.Contains(contentLower, keyword) {
                topics = append(topics, topic)
                break
            }
        }
    }
    
    return topics
}

func (c *ContextAnalyzer) trackMessage(m *discordgo.MessageCreate) {
    messages := c.recentMessages[m.ChannelID]
    messages = append(messages, m)
    
    // Keep only last 20 messages
    if len(messages) > 20 {
        messages = messages[len(messages)-20:]
    }
    
    c.recentMessages[m.ChannelID] = messages
}