package handlers

// HandlerLoggerFieldKey is the key used for the logger field in the handler.
const HandlerLoggerFieldKey = "bloopyHandler"

type AppCommandName string
const (
	Blissfest AppCommandName = "blissfest"
	Book AppCommandName = "book"
	Requests AppCommandName = "requests"
	AnalyzeEmoji AppCommandName = "analyze"
	Inspiro AppCommandName = "inspire"
)
