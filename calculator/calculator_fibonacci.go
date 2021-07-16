package calculator

func (m *Calculator) fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	if m.Contains(n) {
		return m.Read(n)
	}

	return m.Get(n-1) + m.Get(n-2)
}
