package emoji

type EmojiPool interface {
	Size() int
	All() []interface{}
	Add(...interface{}) error
}
