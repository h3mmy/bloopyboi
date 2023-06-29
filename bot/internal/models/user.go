package models

type User struct {
	ID				string		`json:"id,omitempty"`
	DiscordId		string		`json:"discordid"`
	PlexId			string		`json:"plexid,omitempty"`
	BloopnetId		string		`json:"bloopnetid"`
	AuthentikPKID	string		`json:"authentikid"`
}
