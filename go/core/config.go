package core

import (
	"sync/atomic"

	"github.com/toothrot/guff/go/generated"
)

type Config struct {
	AppStatus guff_proto.AppStatusCode
}

type ConfigStore struct {
	value *atomic.Value
}

func NewConfigStore(conf Config) *ConfigStore {
	cs := &ConfigStore{}
	cs.SetConfig(conf)
	return cs
}

func (c *ConfigStore) GetAppStatus() guff_proto.AppStatusCode {
	code := guff_proto.AppStatusCode_APP_STATUS_UNKNOWN
	if c.value.Load() == nil {
		return code
	}
	conf, ok := c.value.Load().(Config)
	if !ok {
		return code
	}
	return conf.AppStatus
}

func (c *ConfigStore) GetConfig() Config {
	conf, ok := c.value.Load().(Config)
	if !ok {
		return Config{}
	}
	return conf
}

func (c *ConfigStore) SetConfig(conf Config) {
	if c.value == nil {
		c.value = &atomic.Value{}
	}
	c.value.Store(conf)
}
