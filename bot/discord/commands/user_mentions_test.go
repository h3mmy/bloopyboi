package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBotReferenced(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected bool
	}{
		{
			name:     "Exact match",
			content:  "bloopyboi",
			expected: true,
		},
		{
			name:     "Case insensitive exact match",
			content:  "BloopyBoi",
			expected: true,
		},
		{
			name:     "Contains bloopyboi",
			content:  "hey bloopyboi how are you",
			expected: true,
		},
		{
			name:     "Bloopy boi with space",
			content:  "hey bloopy boi",
			expected: true,
		},
		{
			name:     "Bloopy boi with multiple spaces",
			content:  "hey bloopy   boi",
			expected: true,
		},
		{
			name:     "Bloop boi",
			content:  "bloop boi is here",
			expected: true,
		},
		{
			name:     "The boi",
			content:  "referencing the boi",
			expected: true,
		},
		{
			name:     "The bot",
			content:  "where is the bot?",
			expected: true,
		},
		{
			name:     "Bloopy",
			content:  "hey bloopy!",
			expected: true,
		},
		{
			name:     "Fuzzy match typo bloopyboii",
			content:  "bloopyboii is cool",
			expected: true,
		},
		{
			name:     "Fuzzy match typo bloopyboy",
			content:  "hey bloopyboy",
			expected: true,
		},
		{
			name:     "False positive the bottom",
			content:  "at the bottom of the sea",
			expected: false,
		},
		{
			name:     "False positive blooming",
			content:  "the flowers are blooming",
			expected: false,
		},
		{
			name:     "False positive boil",
			content:  "water will boil",
			expected: false,
		},
		{
			name:     "Random message",
			content:  "hello world",
			expected: false,
		},
		{
			name:     "Short word no match",
			content:  "the",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsBotReferenced(tt.content))
		})
	}
}
