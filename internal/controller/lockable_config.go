package controller

import (
	"sync"

	traefikconfig "github.com/traefik/traefik/v3/pkg/config/dynamic"
)

type LockableConfig struct {
	Mutex  sync.Mutex
	Config *traefikconfig.Configuration
}

func (LC *LockableConfig) GetConfig() *traefikconfig.Configuration {
	LC.Mutex.Lock()
	defer LC.Mutex.Unlock()
	return LC.Config
}
func (LC *LockableConfig) UpdateConfig(config *traefikconfig.Configuration) {
	LC.Mutex.Lock()
	defer LC.Mutex.Unlock()
	LC.Config = config
}
