package hashtables

import "fmt"

func Example1() {
	var table HashTable = New(7)

	table.Insert("He is", Song{
		name:            "He is",
		artist:          "Ghost",
		album:           "Meliora",
		durationMinutes: 253,
		year:            2015,
	})
	table.Insert("Spirit", Song{
		name:            "Spirit",
		artist:          "Ghost",
		album:           "Meliora",
		durationMinutes: 315,
		year:            2015,
	})
	fmt.Println(table, table.length)

	// queue.Enqueue(101)
	// fmt.Println(queue.Repr())

	// queue.Enqueue(102)
	// fmt.Println(queue.Repr())

	// queue.Enqueue(103)
	// fmt.Println(queue.Repr())

	// fmt.Println("Peek:", queue.Peek())

	// for !queue.isEmpty() {
	// 	fmt.Println("Popping", queue.Dequeue())
	// 	fmt.Println(queue.Repr())
	// }
}
