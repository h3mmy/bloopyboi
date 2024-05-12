package discord

import (
	"strconv"
	"time"
)

// SnowflakeTimestamp returns the creation time of a Snowflake ID relative to the creation of Discord.
// https://discord.com/developers/docs/reference#snowflakes
//
// 111111111111111111111111111111111111111111 11111 11111 111111111111
// 64                                         22    17    12          0
//
// Field	Bits	Number of bits	Description	Retrieval
// Timestamp	63 to 22	42 bits	Milliseconds since Discord Epoch, the first second of 2015 or 1420070400000.	(snowflake >> 22) + 1420070400000
// Internal worker ID	21 to 17	5 bits		(snowflake & 0x3E0000) >> 17
// Internal process ID	16 to 12	5 bits		(snowflake & 0x1F000) >> 12
// Increment	11 to 0	12 bits	For every ID that is generated on that process, this number is incremented	snowflake & 0xFFF
func SnowflakeTimestamp(ID string) (t time.Time, err error) {
	i, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		return
	}
	timestamp := (i >> 22) + DiscordEpoch
	t = time.Unix(0, timestamp*1000000)
	return
}
