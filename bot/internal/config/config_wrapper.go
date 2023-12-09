package config

import (
	"sync"
	"time"
)

type AppConfig struct {
	m sync.RWMutex
	bloopyConfig *BotConfig
	lastUpdated time.Time
	revision int
}

func NewAppConfig(botConfig *BotConfig) *AppConfig {
	return &AppConfig{
		bloopyConfig: botConfig,
		lastUpdated: time.Now(),
		revision: 1,
	}
}

func (c *AppConfig) GetRevision() int {
	return c.revision
}

func (c *AppConfig) UpdateConfig(botConfig *BotConfig) {
	c.m.Lock()
	c.bloopyConfig = botConfig
	c.lastUpdated = time.Now()
	c.revision++
	c.m.Unlock()
}

func (c *AppConfig) GetConfig() *BotConfig {
	return c.bloopyConfig
}
