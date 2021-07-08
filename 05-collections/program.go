package main

import "fmt"

func main() {
	//Array
	//var nos [5]int
	//var nos [5]int = [5]int{2, 4, 12, 4, 2}
	var nos [5]int = [...]int{2, 4, 12, 4, 2}
	fmt.Println(nos)

	// iterating using the indexer
	for i := 0; i < len(nos); i++ {
		fmt.Println(nos[i])
	}

	//iterating using the range construct
	for idx, no := range nos {
		fmt.Printf("Value at [%d] is %d\n", idx, no)
	}

	// creating copy of the array
	newnos := nos
	fmt.Println(newnos)
	newnos[0] = 23
	fmt.Println(newnos)
	fmt.Println(nos)

	//slice
	var names []string = []string{"Vittal", "Vitral"}
	fmt.Println(names)
	fmt.Println(names, len(names))

	//adding new values
	names = append(names, "pai")
	fmt.Println(names)

	names = append(names, "matral", "pitral")
	fmt.Println(names)

	newNames := []string{"john", "bond"}
	names = append(names, newNames...)
	fmt.Println(names)

	//slicing
	/*
		[lo:hi] => from lo to hi-1
		[lo:] => from lo to last
		[:hi] => from 0 to  hi-1
		[lo:lo] => empty slice
		[lo:lo+1] => item at index lo
	*/
	fmt.Println(names[2:5])
	fmt.Println(names[2:])
	fmt.Println(names[:3])
	fmt.Println(names[4:4])
	fmt.Println(names[4:5])

	dupNames := names[:]
	dupNames[0] = "Ram"
	fmt.Println(dupNames, names)

	dupNames = append(dupNames, "smith")
	fmt.Println(dupNames, names)

	data := make([]int, 10)
	fmt.Println(data)
	data = append(data, 100)
	fmt.Println(data)

	//maps
	fmt.Println()
	fmt.Println("Maps")
	var cityRanks map[string]int = map[string]int{
		"Bengaluru": 5,
		"Udupi":     1,
		"Mangaluru": 3,
		"Mysuru":    2,
	}
	fmt.Println(cityRanks)
	fmt.Println("Rank of Mysuru =>", cityRanks["Mysuru"])

	//adding a new entry
	cityRanks["Chennai"] = 6
	fmt.Println(cityRanks)

	// checking if a key exists
	if rank, exists := cityRanks["Tumukur"]; exists {
		fmt.Println("Rank of tumkuru is =>", rank)
	} else {
		fmt.Println("Tumkur is not ranked yet!")
	}

	//removing an item
	delete(cityRanks, "Chennai")
	fmt.Println("After renoving chennai, cityrank => ", cityRanks)

	//iterate through the map
	for key, value := range cityRanks {
		fmt.Printf("Value of [%s] is %d\n", key, value)
	}
}
