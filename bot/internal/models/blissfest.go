package models

import "time"

type BlissfestConfig struct {
	// DateTime when gates open
	Start time.Time
	// DateTime when you have to be out by
	End      time.Time
	Homepage string
}
