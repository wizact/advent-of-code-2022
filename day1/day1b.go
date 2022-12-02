package day1

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/wizact/advent-of-code-2022/file"
)

type Calorie struct {
	Key   int
	Value int
}

type CalorieList []Calorie

func (p CalorieList) Len() int           { return len(p) }
func (p CalorieList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p CalorieList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func Day1b() {
	d, e := file.ReadFile("./day1/day1a.txt")

	if e != nil {
		panic(e)
	}

	calories := strings.Split(string(d), "\n")
	noOfElves := strings.Split(string(d), "\n\n")
	elvs_calories := make(CalorieList, len(noOfElves))
	elf_index := 0
	current_elf_calorie := 0
	for _, v := range calories {
		if v == "" {
			elvs_calories[elf_index] = Calorie{elf_index, current_elf_calorie}
			current_elf_calorie = 0
			elf_index++
		} else {
			c, e := strconv.Atoi(v)

			if e != nil {
				panic(e)
			}

			current_elf_calorie += c
		}
	}

	sort.Sort(sort.Reverse(elvs_calories))

	fmt.Printf("The top three elves with the most calories have %v calories in total\n", elvs_calories[0].Value+elvs_calories[1].Value+elvs_calories[2].Value)
}
