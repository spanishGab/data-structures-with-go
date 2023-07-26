package hashtables

import (
	"encoding/json"
	"fmt"
)

func Example1() {
	var table HashTable = New(7)

	table.Insert(666, Song{
		Id:              666,
		Name:            "He is",
		Artist:          "Ghost",
		Album:           "Meliora",
		DurationMinutes: 253,
		Year:            2015,
	})
	table.Insert(667, Song{
		Id:              667,
		Name:            "Spirit",
		Artist:          "Ghost",
		Album:           "Meliora",
		DurationMinutes: 315,
		Year:            2015,
	})
	fmt.Println(table.Repr())

	heIsSong, _ := json.MarshalIndent(table.Get(666), "", "  ")
	fmt.Println(string(heIsSong))
}
