package day6

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

func Day6a() {
	//processStream(4)
	processStream(14)
}

func processStream(buffSize int) {
	f, e := os.Open("./day6/day6.txt")

	if e != nil {
		panic(e)
	}

	br := bufio.NewReader(f)

	a := make([]string, buffSize)
	c := 0
	// infinite loop
	for {

		b, err := br.ReadByte()

		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Println(err)
			break
		}

		if c < buffSize {
			a[c] = string(b)
		} else {
			a = append(a[1:], string(b))
		}
		c = c + 1
		fmt.Println(a)

		// check the elements for being identical from each other
		hs := hasDuplicates(a)

		if !hs && c > buffSize {
			fmt.Println("Found the element at:", c, a)
			break
		}

		if err != nil {
			// end of file
			break
		}
	}
}

func hasDuplicates(arr []string) bool {
	m := make(map[string]bool)
	for _, v := range arr {
		if m[v] {
			return true
		} else {
			m[v] = true
		}
	}
	return false
}
