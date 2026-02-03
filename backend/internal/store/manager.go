package store

import (
	"sync"

	"github.com/getkin/kin-openapi/openapi3"
)

type SpecManager struct {
	mu     sync.RWMutex
	specs  map[string]*openapi3.T
	active string
}

var Manager = &SpecManager{
	specs: make(map[string]*openapi3.T),
}

func (m *SpecManager) Set(env string, spec *openapi3.T) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.specs[env] = spec
}

func (m *SpecManager) GetActive() *openapi3.T {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.specs[m.active]
}

func (m *SpecManager) SetActive(env string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.active = env
}

func (m *SpecManager) List() []string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	envs := make([]string, 0, len(m.specs))
	for k := range m.specs {
		envs = append(envs, k)
	}
	return envs
}
