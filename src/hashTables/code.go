package hashtables

import (
	"math"
)

type Song struct {
	name            string
	artist          string
	album           string
	durationMinutes uint32
	year            uint16
}

type HashTable struct {
	length   int
	capacity int
	items    []*Song
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

func (table HashTable) Insert(key string, value Song) {
	if table.length == table.capacity {
		panic("The HashTable is full! Cannot any more elements to it")
	}
	table.length += 1 // TODO: Bug, length does not increases
	var keyValue int = getStringKeyValue(key)
	var song Song = Song{
		name:            value.name,
		artist:          value.artist,
		album:           value.album,
		durationMinutes: value.durationMinutes,
		year:            value.year,
	}
	var pos int = hashByMultiplication(float64(keyValue), float64(table.capacity))
	if table.items[pos] == nil {
		table.items[pos] = &song
		return
	}
	for i := 0; i < table.capacity; i++ {
		newPos := doubleHash(pos, float64(keyValue), i, float64(table.capacity))
		if table.items[newPos] == nil {
			table.items[newPos] = &song
			return
		}
	}
}
