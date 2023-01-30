package config

import (
	"sync"
)

func init() {
	storage = &Cfg[any]{
		h: make(map[string]any),
		m: sync.RWMutex{},
	}
}

type (
	Cfg[V any] struct {
		h map[string]V
		m sync.RWMutex
	}
	Value struct {
		any
	}
)

var storage *Cfg[any]

func GetValue(key string) Value {
	storage.m.RLock()
	defer storage.m.RUnlock()
	if value, ok := storage.h[key]; ok {
		return Value{value}
	}
	return Value{newConfigError("value %+v doesn't exists", ErrNotFound)}
}

func (v Value) String() string {
	if j, ok := v.any.(string); ok {
		return j
	}
	return ""
}

func (v Value) Int() int {
	if j, ok := v.any.(int); ok {
		return j
	}
	return 0
}
