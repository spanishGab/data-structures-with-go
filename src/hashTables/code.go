package hashtables

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
