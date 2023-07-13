package hashtables

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"
)

type Song struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Artist          string `json:"artist"`
	Album           string `json:"album"`
	DurationMinutes uint32 `json:"durationMinutes"`
	Year            uint16 `json:"year"`
}

type HashTable struct {
	length   int
	capacity int
	items    []*Song
	keys     []int
}

const hashMultiplicationFactor float64 = 0.618

// Remember to use a prime number to capacity,
// it reduces the chances of colisions
func New(capacity int) HashTable {
	var hash HashTable = HashTable{
		length:   0,
		capacity: capacity,
		items:    make([]*Song, capacity),
	}
	for i := 0; i < hash.capacity; i++ {
		hash.items[i] = nil
	}
	return hash
}

// Utilitary function to convert keys in integers
// It can be used in case you want to store string keys
func getStringKeyValue(key string) int {
	keyValue := 7
	keyCharacters := []rune(key)
	for i := 0; i < len(key); i++ {
		keyValue = 31*keyValue + int(keyCharacters[i])
	}
	return keyValue
}

func hashByMultiplication(keyValue float64, tableSize float64) int {
	keyRemainder := math.Mod(keyValue*hashMultiplicationFactor, 1)
	return int(tableSize * keyRemainder)
}

func doubleHash(hash1 int, keyValue float64, currentPos int, tableSize float64) int {
	hash2 := hashByMultiplication(keyValue, tableSize-1) + 1
	doubleHashedPos := math.Abs(float64(hash1) + float64(currentPos*hash2))
	return int(doubleHashedPos) % int(tableSize)
}

func (table *HashTable) checkEmptyPosition(pos int) bool {
	if table.items[pos] == nil {
		return true
	}
	return false
}

func (table *HashTable) Insert(key int, value Song) {
	if table.length == table.capacity {
		panic("The HashTable is full! Cannot any more elements to it")
	}
	table.length += 1 // TODO: Bug, length does not increases
	table.keys = append(table.keys, key)
	var song Song = Song{
		Id:              value.Id,
		Name:            value.Name,
		Artist:          value.Artist,
		Album:           value.Album,
		DurationMinutes: value.DurationMinutes,
		Year:            value.Year,
	}
	var pos int = hashByMultiplication(float64(key), float64(table.capacity))
	var newPos int
	for i := 0; i < table.capacity; i++ {
		newPos = doubleHash(pos, float64(key), i, float64(table.capacity))
		fmt.Println(pos, newPos, i)
		if table.checkEmptyPosition(newPos) {
			fmt.Println("Empty pos", newPos, table.items[newPos])
			table.items[newPos] = &song
			return
		}
	}
}

func (table *HashTable) Get(key int) *Song {
	var pos int = hashByMultiplication(float64(key), float64(table.capacity))
	var newPos int
	for i := 0; i < table.capacity; i++ {
		newPos = doubleHash(pos, float64(key), i, float64(table.capacity))
		if table.checkEmptyPosition(newPos) {
			return nil
		}
		if table.items[newPos].Id == key {
			return table.items[newPos]
		}
	}
	return nil
}

func (table *HashTable) Repr() string {
	var tableElements []string
	var jsonSong []byte
	for _, key := range table.keys {
		jsonSong, _ = json.Marshal(table.Get(key))
		tableElements = append(tableElements, fmt.Sprintf("\"%d\": %s", key, string(jsonSong)))
	}
	return fmt.Sprintf(
		"table length: %d\ntable elements: %s",
		table.length,
		"{"+strings.Join(tableElements, ",")+"}",
	)
}
