package asynchandlers

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

func TestFindSimilarEmoji(t *testing.T) {
	mr := NewMessageReactor()
	t.Run("should return emojis that meet the dynamic similarity threshold", func(t *testing.T) {
		message := &discordgo.Message{
			Content: "I saw a happy cat today",
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
		assert.Len(t, result, 2)
		var resultNames []string
		for _, e := range result {
			resultNames = append(resultNames, e.Name)
		}
		assert.Contains(t, resultNames, "cat")
		assert.Contains(t, resultNames, "happy")
	})
	t.Run("should return original pool when no keywords are extracted", func(t *testing.T) {
		message := &discordgo.Message{
			Content: "...",
		}
		emojiPool := []*discordgo.Emoji{
			{Name: "dog"},
			{Name: "cat"},
		}

		result := mr.FindSimilarEmoji(message, emojiPool)

		assert.NotNil(t, result)
		assert.Equal(t, emojiPool, result)
	})

	t.Run("should return an empty slice when no emojis meet the threshold", func(t *testing.T) {
		message := &discordgo.Message{
			Content: "this is a test message about a happy cat",
		}
		emojiPool := []*discordgo.Emoji{
			{Name: "xyz"},
			{Name: "abc"},
		}

		result := mr.FindSimilarEmoji(message, emojiPool)

		assert.NotNil(t, result)
		assert.Empty(t, result)
	})
}
