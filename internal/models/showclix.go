package models

const ShowclixAPIURL = "https://api.showclix.com"
const ShowclixAPIEventPrefix = "/Event"

type PriceLevel struct {
	LevelID string `json:"level_id"` //int
	EventID string `json:"event_id"` //int
	Price string `json:"price"` // float64
	MinPrice string `json:"min_price,omitempty"` // float64
	Level string `json:"level,omitempty"` // string
	Active string `json:"active"` // int bool
  Description string `json:"description,omitempty"` // string
	Subheading string `json:"subheading,omitempty"` // string
	ParentID string `json:"parent_id,omitempty"` // int
	Position string `json:"position,omitempty"` // int
	IncrementBy string `json:"increment_by,omitempty"` // int
	TransactionLimit string `json:"transaction_limit,omitempty"` // int
	TicketLayoutID string `json:"ticket_layout_id,omitempty"` // int
	UpsellPrice string `json:"upsell_price,omitempty"` // float
	DiscountType string `json:"discount_type,omitempty"` // string
	DiscountVal string `json:"discount_val,omitempty"` // float64
}
