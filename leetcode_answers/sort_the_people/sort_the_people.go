package main

import "sort"

func main() {}

type Person struct {
	name   string
	height int
}

func sortPeople(names []string, heights []int) []string {
	p := []Person{}
	for i := 0; i < len(names); i++ {
		p = append(p, Person{names[i], heights[i]})
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].height >= p[j].height
	})

	result := make([]string, len(p))
	for i := 0; i < len(names); i++ {
		result[i] = p[i].name
	}
	return result
}
