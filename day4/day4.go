package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wizact/advent-of-code-2022/file"
)

func getFile() []byte {
	b, e := file.ReadFile("./day4/day4a.txt")

	if e != nil {
		panic(e)
	}
	return b
}

func Day4() {
	m := GetRawSections()

	for k, v := range m {
		rangesOverlapCompletely(&v)
		rangesOverlap(&v)
		m[k] = v
	}

	totalCompleteOverlaps := 0
	totalPartialOverlaps := 0
	for _, v := range m {
		if v.completeOverlap {
			totalCompleteOverlaps = totalCompleteOverlaps + 1
		}

		if v.partiallyOverlaps {
			totalPartialOverlaps = totalPartialOverlaps + 1
		}
	}

	fmt.Println(totalCompleteOverlaps, totalPartialOverlaps)

}

func rangesOverlapCompletely(s *Sections) {
	if s.pair1Lower <= s.pair2Lower && s.pair1Upper >= s.pair2Upper ||
		s.pair2Lower <= s.pair1Lower && s.pair2Upper >= s.pair1Upper {
		s.completeOverlap = true
	}
}

func rangesOverlap(s *Sections) {
	if rangeOverlap(s.pair1Lower, s.pair1Upper, s.pair2Lower, s.pair2Upper) {
		s.partiallyOverlaps = true
	}
}

func rangeOverlap(l1, u1, l2, u2 int) bool {
	return l2 >= l1 && l2 <= u1 || u2 >= l2 && u2 <= u1
}

func GetRawSections() map[int]Sections {
	b := getFile()
	c := string(b)
	sids := strings.Split(c, "\n")
	noSids := len(sids)
	m := make(map[int]Sections, noSids)

	for k, v := range sids {
		secs := strings.Split(v, ",")
		pair1Raw := secs[0]
		pair2Raw := secs[1]
		p1L, p1U := calcLowerUpper(pair1Raw)
		p2L, p2U := calcLowerUpper(pair2Raw)

		rs := Sections{pair1Raw: pair1Raw, pair2Raw: pair2Raw, pair1Lower: p1L, pair1Upper: p1U, pair2Lower: p2L, pair2Upper: p2U}
		m[k] = rs
	}
	return m
}

func calcLowerUpper(sr string) (int, int) {
	ps := strings.Split(sr, "-")
	p1, _ := strconv.Atoi(ps[0])
	p2, _ := strconv.Atoi(ps[1])
	return p1, p2
}

type Sections struct {
	pair1Raw string
	pair2Raw string

	pair1Lower int
	pair1Upper int

	pair2Lower int
	pair2Upper int

	completeOverlap   bool
	partiallyOverlaps bool
}
