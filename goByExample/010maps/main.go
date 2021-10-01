package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 8
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	// 2nd arg returned when getting value indicates if key present in map
	_, prs := m["k2"]
	fmt.Println("present:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
