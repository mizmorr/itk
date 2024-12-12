package main

import (
	"fmt"
	"sync"
)

var hashSize = 7

const needToRehashValue = 6

// HashTable структура для хеш-таблицы
type HashTable struct {
	slice []*bucket
}

func newHashTable(size int) *HashTable {
	table := &HashTable{slice: make([]*bucket, size)}
	for i := 0; i < size; i++ {
		table.slice[i] = &bucket{}
	}
	return table
}

// bucket структура, представляющая список для разрешения коллизий
type bucket struct {
	head *bucketNode
}

// bucketNode структура узла в списке
type bucketNode struct {
	key   string
	value int
	next  *bucketNode
}

// Insert добавляет пару ключ-значение в хеш-таблицу
func (h *HashTable) Insert(key string, value int) {
	index := hash(len(h.slice), key)
	if h.slice[index].size() > needToRehashValue {
		if h.needToRehash() {
			h.rehash()
			return
		}
	}
	h.slice[index].insert(key, value)
}

// Search возвращает значение по ключу в хеш-таблице
func (h *HashTable) Search(key string) (int, bool) {
	index := hash(len(h.slice), key)
	return h.slice[index].search(key)
}

// Delete удаляет элемент по ключу из хеш-таблицы
func (h *HashTable) Delete(key string) {
	index := hash(len(h.slice), key)
	h.slice[index].delete(key)
}

func (h *HashTable) needToRehash() bool {
	chanWithSize := make(chan int)
	go func() {
		defer close(chanWithSize)
		for k := range h.slice {
			chanWithSize <- h.slice[k].size()
		}
	}()
	var bucketSizes int
	for size := range chanWithSize {
		bucketSizes += size
	}
	return bucketSizes/hashSize > needToRehashValue
}

func (h *HashTable) rehash() {
	newSize := len(h.slice) + 1
	newHash := newHashTable(newSize)

	mutex := &sync.Mutex{}
	for i := range h.slice {
		go func() {
			cur := h.slice[i].head
			for cur != nil {
				index := hash(newSize, cur.key)
				mutex.Lock()
				newHash.slice[index].insert(cur.key, cur.value)
				mutex.Unlock()
				cur = cur.next
			}
		}()
	}
	*h = *newHash
}

func (b *bucket) size() (count int) {
	curElem := (b.head)
	for curElem != nil {
		count++
		curElem = curElem.next
	}
	return
}

// hash простая хеш-функция
func hash(hashTableSize int, key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % hashTableSize
}

// insert добавляет элемент в список
func (b *bucket) insert(key string, value int) {
	if !b.searchKey(key) {
		newNode := &bucketNode{key: key, value: value}
		newNode.next = b.head
		b.head = newNode
	} else {
		// Обновляем значение, если ключ уже существует
		currNode := b.head
		for currNode.key != key {
			currNode = currNode.next
		}
		currNode.value = value
	}
}

// search ищет ключ в списке
func (b *bucket) search(key string) (int, bool) {
	currNode := b.head
	for currNode != nil {
		if currNode.key == key {
			return currNode.value, true
		}
		currNode = currNode.next
	}
	return 0, false
}

func (b *bucket) searchKey(key string) bool {
	_, isFound := b.search(key)
	return isFound
}

// delete удаляет ключ из списка
func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	prevNode := b.head
	for prevNode.next.key != key {
		prevNode = prevNode.next
	}
	prevNode.next = prevNode.next.next
}

func main() {
	hashTable := newHashTable(hashSize)
	fmt.Println("Hash Table Created")
	hashTable.Insert("name", 123)
	hashTable.Insert("age", 456)
	val, ok := hashTable.Search("name")

	fmt.Println(ok, val)

	hashTable.rehash()
}
