package calculator

import "sync"

func (m *Calculator) Contains(key int) bool {
	m.lock.RLock()
	_, exists := m.cache[key]
	m.lock.RUnlock()
	return exists
}

func (m *Calculator) Set(key int, value int) {
	m.lock.Lock()
	m.cache[key] = value
	m.lock.Unlock()
}

func (m *Calculator) Read(key int) int {
	m.lock.RLock()
	result, _ := m.cache[key]
	m.lock.RUnlock()
	return result
}

func (m *Calculator) Get(key int) int {
	if !m.Contains(key) {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			value := m.fibonacci(key)
			m.Set(key, value)
			wg.Done()
		}()
		wg.Wait()
	}
	return m.Read(key)
}
