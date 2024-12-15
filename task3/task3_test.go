package task3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	// Создаем новый объект StringIntMap
	m := &StringIntMap{innerMap: make(map[string]int)}

	// Добавляем элемент
	m.Add("apple", 5)

	// Проверяем, что элемент был добавлен правильно
	value, exists := m.Get("apple")
	assert.True(t, exists, "Apple should exist in the map")
	assert.Equal(t, 5, value, "Value for apple should be 5")
}

func TestGet(t *testing.T) {
	m := &StringIntMap{innerMap: make(map[string]int)}
	m.Add("orange", 3)

	// Проверяем, что можно получить значение по ключу
	value, exists := m.Get("orange")
	assert.True(t, exists, "Orange should exist in the map")
	assert.Equal(t, 3, value, "Value for orange should be 3")

	// Проверяем несуществующий ключ
	_, exists = m.Get("banana")
	assert.False(t, exists, "Banana should not exist in the map")
}

func TestDelete(t *testing.T) {
	m := &StringIntMap{innerMap: make(map[string]int)}
	m.Add("banana", 7)

	// Удаляем элемент
	m.Delete("banana")

	// Проверяем, что элемент был удалён
	_, exists := m.Get("banana")
	assert.False(t, exists, "Banana should be deleted from the map")
}

func TestExists(t *testing.T) {
	m := &StringIntMap{innerMap: make(map[string]int)}
	m.Add("grape", 10)

	// Проверяем существование ключа
	assert.True(t, m.Exists("grape"), "Grape should exist in the map")

	// Проверяем несуществующий ключ
	assert.False(t, m.Exists("apple"), "Apple should not exist in the map")
}

func TestCopy(t *testing.T) {
	m := &StringIntMap{innerMap: make(map[string]int)}
	m.Add("kiwi", 12)
	m.Add("mango", 8)

	// Получаем копию карты
	copyMap := m.Copy()

	// Проверяем, что копия содержит правильные данные
	assert.Equal(t, 12, copyMap["kiwi"], "Copy should have kiwi with value 12")
	assert.Equal(t, 8, copyMap["mango"], "Copy should have mango with value 8")

	// Проверяем, что копия независима
	m.Add("kiwi", 15) // изменим оригинальную карту
	assert.NotEqual(t, copyMap["kiwi"], m.innerMap["kiwi"], "The copied map should be independent of the original map")
}
