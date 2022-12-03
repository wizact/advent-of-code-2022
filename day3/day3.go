package day3

import (
	"strings"

	"github.com/wizact/advent-of-code-2022/file"
)

func getFile() []byte {
	b, e := file.ReadFile("./day3/day3a.txt")

	if e != nil {
		panic(e)
	}
	return b
}

func createRucksack(b []byte) map[int]Rucksack {
	i := string(b)
	rss := strings.Split(i, "\n")
	noRss := len(rss)
	m := make(map[int]Rucksack, noRss)

	for k, v := range rss {
		t := len(v)
		c1 := v[0:(t / 2)]
		c2 := v[(t / 2):t]

		rs := Rucksack{compartment1: c1, compartment2: c2, totalPackage: v}
		m[k] = rs
	}
	return m
}

func createPriorityList() map[string]int {
	m := make(map[string]int, 52)
	index := 0
	for i := 'a'; i <= 'z'; i++ {
		index = index + 1
		m[string(i)] = index
	}

	for i := 'A'; i <= 'Z'; i++ {
		index = index + 1
		m[string(i)] = index
	}

	return m
}

type Rucksack struct {
	totalPackage  string
	compartment1  string
	compartment2  string
	sharedElement string
	priorityScore int
}

type Items struct {
	letter   string
	priority int
}
