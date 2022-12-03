package day3

import (
	"bytes"
	"fmt"
)

func Day3a() {
	b := getFile()
	rsl := createRucksack(b)

	for k, v := range rsl {
		findSharedItem(&v)
		rsl[k] = v
	}

	ts := 0
	for _, v := range rsl {
		ts = ts + v.priorityScore
	}

	fmt.Println("Total priority score is", ts)
}

func findSharedItem(r *Rucksack) {
	pl := createPriorityList()
	bc1 := []byte(r.compartment1)
	bc2 := []byte(r.compartment2)

	for _, v := range bc1 {
		if bytes.Index(bc2, []byte{v}) > -1 {
			r.sharedElement = string(v)
			r.priorityScore = pl[string(v)]
			return
		}
	}

	panic("error in input")
}
