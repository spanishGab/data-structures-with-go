package hashtables

import "math"

type Student struct {
	id     int
	name   string
	grade1 float32
	grade2 float32
	grade3 float32
}

type HashTable struct {
	length   int
	capacity int
	items    []*Student
}

const hashMultiplicationFactor float64 = 0.618

// Remember to use a prime number to capacity,
// it reduces the chances of colisions
func New(capacity int) HashTable {
	var hash HashTable = HashTable{
		length:   0,
		capacity: capacity,
		items:    make([]*Student, capacity),
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
    keyValue = 31 * keyValue + int(keyCharacters[i])
  }
  return keyValue
}


func (table HashTable) hashByMultiplication(keyValue int) int {
  keyRemainder := math.Mod(float64(keyValue) * hashMultiplicationFactor, 1)
  return int(float64(table.capacity) * keyRemainder)
}
