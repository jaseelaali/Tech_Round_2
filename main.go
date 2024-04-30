package main

import (
	"fmt"
)

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	var arr []string
	mymap := make(map[string]bool)
	for _, v := range developers {
		if _, found := mymap[v.Name]; found {
			continue
		} else {
			mymap[v.Name] = true
			arr = append(arr, v.Name)
		}
	}
	return arr
}

func main() {
	fmt.Println("Filter Unique Challenge")
	dev := []Developer{
		{Name: "alin"},
		{Name: "alin"},
		{Name: "anu"},
		{Name: "anu"},
		{Name: "vinu"},

	}
	fmt.Println(FilterUnique(dev))
}
