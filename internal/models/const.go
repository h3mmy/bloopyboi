package models

type CtxKey string
type ColorCode int
type MediaRequestStatus string
type MediaRequestStatusCategory string

const (
	MediaRequestStatusCategoryPending  MediaRequestStatusCategory = "pending"
	MediaRequestStatusCategoryComplete MediaRequestStatusCategory = "complete"
	MediaRequestStatusCategoryBlocked  MediaRequestStatusCategory = "blocked"
)

var (
	MediaRequestStatusCategoryMap = map[MediaRequestStatus]MediaRequestStatusCategory{
		MediaRequestStatusRequested: MediaRequestStatusCategoryPending,
		MediaRequestStatusApproved:  MediaRequestStatusCategoryPending,
		MediaRequestStatusCancelled: MediaRequestStatusCategoryComplete,
		MediaRequestStatusRejected:  MediaRequestStatusCategoryComplete,
		MediaRequestStatusSuccess:   MediaRequestStatusCategoryComplete,
		MediaRequestStatusError:     MediaRequestStatusCategoryBlocked,
	}
)

const (
	MediaRequestStatusRequested MediaRequestStatus = "requested"
	MediaRequestStatusApproved  MediaRequestStatus = "approved"
	MediaRequestStatusError     MediaRequestStatus = "error"
	MediaRequestStatusSuccess   MediaRequestStatus = "success"
	MediaRequestStatusCancelled MediaRequestStatus = "cancelled"
	MediaRequestStatusRejected  MediaRequestStatus = "rejected"
)

const (
	CtxKeyInteraction        CtxKey = "discord_interaction_id"
	CtxKeyRequestID          CtxKey = "media_request_id"
	CtxKeyMediaRequestStatus CtxKey = "media_request_status"
	CtxKeyMessageID          CtxKey = "discord_message_id"
	CtxChannelID             CtxKey = "discord_channel_id"
	CtxDiscordGuildID        CtxKey = "discord_guild_id"
)

const (
	ColorCodeDefault ColorCode = iota
	ColorCodeSuccess           = 0x00ff00
	ColorCodeInfo              = 0x0000ff
	ColorCodeWarning           = 0xffa500
	ColorCodeDanger            = 0xff0000
)

const (
	ColorCodeDefaultStr = "default"
	ColorCodeSuccessStr = "success"
	ColorCodeInfoStr    = "info"
	ColorCodeWarningStr = "warning"
	ColorCodeDangerStr  = "danger"
)

var (
	ColorCodeMap = map[string]ColorCode{
		ColorCodeDefaultStr: ColorCodeDefault,
		ColorCodeSuccessStr: ColorCodeSuccess,
		ColorCodeInfoStr:    ColorCodeInfo,
		ColorCodeWarningStr: ColorCodeWarning,
		ColorCodeDangerStr:  ColorCodeDanger,
	}
)

// Values provides list valid values for Enum.
func (MediaRequestStatus) Values() (kinds []string) {
	for _, s := range []MediaRequestStatus{
		MediaRequestStatusRequested,
		MediaRequestStatusApproved,
		MediaRequestStatusSuccess,
		MediaRequestStatusRejected,
		MediaRequestStatusCancelled,
		MediaRequestStatusError,
	} {
		kinds = append(kinds, string(s))
	}
	return
}
