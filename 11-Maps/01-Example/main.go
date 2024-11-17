package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Gounder Thottam"] = Vertex{
		10.764817, 76.997213,
	}
	m["Other Coordinate"] = Vertex{
		5.334817, 33.447213,
	}
	// fmt.Println(m["Bell Labs"])
	// for range m {
	// 	delete(m, "Gounder Thottam")
	// }
	// fmt.Println(m)
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Printf("m: %v\n", m)
	fmt.Printf("k: %v\n", keys)

	key, ok := m["Gounder Thottam"]
	if ok {
		fmt.Println(key)
	}

	key1 := m["Gounder Thottam"]
	fmt.Println(key1)
}
