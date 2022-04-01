package modules

import (
	"context"

	"github.com/fdurand/freeradius-go"
)

// Memory struct
type CacheEap struct {
	ModuleName string
	radlog     freeradius.Log
}

func NewCacheEap(ctx context.Context, name string) (freeradius.Module, error) {

	Module := &CacheEap{}
	Module.ModuleName = name
	err := Module.NewModule(ctx, name)

	return Module, err
}

func (m *CacheEap) NewModule(ctx context.Context, name string) error {
	return nil
}

func (m *CacheEap) Init(logger freeradius.Log) error {
	m.radlog = logger
	m.radlog.Radlog(freeradius.LogTypeInfo, "Init from go plugin")
	return nil
}

func (m *CacheEap) Authorize(req freeradius.Request) freeradius.RlmCode {
	m.radlog.Info("Authorize in example module called")
	return freeradius.RlmCodeNoop
}
