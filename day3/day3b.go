package day3

import (
	"bytes"
	"fmt"
)

func Day3b() {
	b := getFile()
	rsl := createRucksack(b)

	gm := make(map[int]ElvesGroup)

	for i := 0; i < len(rsl); i = i + 3 {
		gm[i] = ElvesGroup{elfOne: rsl[i], elfTwo: rsl[i+1], elfThree: rsl[i+2]}
	}

	for k, v := range gm {
		findSharedItemBetweenGroup(&v)
		gm[k] = v
	}

	ts := 0
	for _, v := range gm {
		ts = ts + v.groupScore
	}

	fmt.Println("Total group score is", ts)
}

func findSharedItemBetweenGroup(e *ElvesGroup) {
	pl := createPriorityList()
	bc1 := []byte(e.elfOne.totalPackage)
	bc2 := []byte(e.elfTwo.totalPackage)
	bc3 := []byte(e.elfThree.totalPackage)

	for _, v := range bc1 {
		if bytes.Index(bc2, []byte{v}) > -1 && bytes.Index(bc3, []byte{v}) > -1 {
			e.sharedElement = string(v)
			e.groupScore = pl[string(v)]
			return
		}
	}

	panic("error in input")
}

type ElvesGroup struct {
	elfOne   Rucksack
	elfTwo   Rucksack
	elfThree Rucksack

	sharedElement string
	groupScore    int
}
