package task3

type StringIntMap struct {
	innerMap map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{innerMap: map[string]int{}}
}

func (m *StringIntMap) Add(key string, value int) {
	m.innerMap[key] = value
}

func (m *StringIntMap) Get(key string) (int, bool) {
	value, exists := m.innerMap[key]
	return value, exists
}

func (m *StringIntMap) Remove(key string) {
	delete(m.innerMap, key)
}

func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.innerMap[key]
	return exists
}

func (m *StringIntMap) Copy() map[string]int {
	copyMap := make(map[string]int)
	for key, value := range m.innerMap {
		copyMap[key] = value
	}
	return copyMap
}
