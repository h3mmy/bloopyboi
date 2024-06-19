package config

import (
	"sync"
	"time"
)

// AppConfigLoader acts as a config store for the bloopyboi instance
type AppConfigLoader struct {
	m            sync.RWMutex
	bloopyConfig *AppConfig
	lastUpdated  time.Time
	revision     int
}

func NewAppConfigLoader(AppConfig *AppConfig) *AppConfigLoader {
	return &AppConfigLoader{
		bloopyConfig: AppConfig,
		lastUpdated:  time.Now(),
		revision:     1,
	}
}

func (c *AppConfigLoader) GetRevision() int {
	return c.revision
}

func (c *AppConfigLoader) UpdateConfig(AppConfig *AppConfig) {
	c.m.Lock()
	c.bloopyConfig = AppConfig
	c.lastUpdated = time.Now()
	c.revision++
	c.m.Unlock()
}

func (c *AppConfigLoader) GetConfig() *AppConfig {
	return c.bloopyConfig
}
