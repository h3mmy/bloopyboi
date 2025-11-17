package asynchandlers

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestFindSimilarEmoji(t *testing.T) {
	mr := NewMessageReactor()
	message := &discordgo.Message{
		Content: "this is a test message about a happy cat",
	}
	emojiPool := []*discordgo.Emoji{
		{Name: "dog"},
		{Name: "cat"},
		{Name: "happy"},
		{Name: "sad"},
	}

	result := mr.FindSimilarEmoji(message, emojiPool)

	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
	assert.Equal(t, "cat", result[0].Name)
	assert.Equal(t, "happy", result[1].Name)
}
