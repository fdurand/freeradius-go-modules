package modules

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/fdurand/freeradius-go"
)

// Memory struct
type CacheEap struct {
	ModuleName string
	radlog     freeradius.Log
}

func NewCacheEap(name string) (freeradius.Module, error) {

	Module := &CacheEap{}
	Module.ModuleName = name
	err := Module.NewModule(name)

	return Module, err
}

func (m *CacheEap) NewModule(name string) error {
	return nil
}

func (m *CacheEap) Init(logger freeradius.Log) error {
	m.radlog = logger
	m.radlog.Radlog(freeradius.LogTypeInfo, "Init from go plugin")
	return nil
}

func (m *CacheEap) Authorize(req freeradius.Request) freeradius.RlmCode {
	m.radlog.Info("Authorize in cache_eap module called")
	spew.Dump(req.Packet())

	return freeradius.RlmCodeNoop
}
