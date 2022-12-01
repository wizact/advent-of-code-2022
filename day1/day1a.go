package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func Day1a() {
	d, e := readFile()

	if e != nil {
		panic(e)
	}

	current_max := 0
	current_elf_calorie := 0
	calories := strings.Split(string(d), "\n")

	for _, v := range calories {
		if v == "" {
			if current_elf_calorie > current_max {
				current_max = current_elf_calorie
			}
			current_elf_calorie = 0
		} else {
			c, e := strconv.Atoi(v)

			if e != nil {
				panic(e)
			}

			current_elf_calorie += c
		}
	}

	fmt.Printf("The elf with the most calories has %v calories\n", current_max)
}
