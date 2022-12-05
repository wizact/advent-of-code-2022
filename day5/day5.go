package day5

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wizact/advent-of-code-2022/file"
)

func getFile() []byte {
	b, e := file.ReadFile("./day5/day5.txt")

	if e != nil {
		panic(e)
	}
	return b
}

func Day5a() {

	ss := getStartingStack()
	m := getMoves()

	for j := 0; j < len(m); j++ {
		from := ss.items[m[j].from-1]
		to := ss.items[m[j].to-1]
		for i := 0; i < m[j].quantity; i++ {
			f := from.Pop()
			to.Push(f)
		}
	}

	for i := 0; i < 9; i++ {
		fmt.Println(ss.items[i].crates)
	}
}

func Day5b() {

	ss := getStartingStack()
	m := getMoves()

	for j := 0; j < len(m); j++ {
		from := ss.items[m[j].from-1]
		to := ss.items[m[j].to-1]

		interStack := []*Crate{}
		for i := 0; i < m[j].quantity; i++ {
			f := from.Pop()
			interStack = append(interStack, f)
		}

		for i := len(interStack); i > 0; i-- {
			to.Push(interStack[i-1])
		}

	}

	for i := 0; i < 9; i++ {
		fmt.Println(ss.items[i].crates)
	}
}

func getMoves() map[int]Moves {
	b := getFile()
	c := string(b)
	ms := strings.Split(c, "\n")
	noMs := len(ms)
	m := make(map[int]Moves, noMs)

	for k, v := range ms {
		ins := strings.Replace(v, "move ", "", 1)
		ins = strings.Replace(ins, " from ", ",", 1)
		ins = strings.Replace(ins, " to ", ",", 1)

		insArray := strings.Split(ins, ",")
		q, _ := strconv.Atoi(insArray[0])
		f, _ := strconv.Atoi(insArray[1])
		t, _ := strconv.Atoi(insArray[2])

		m[k] = Moves{quantity: q, from: f, to: t}
	}
	return m
}

func getStartingStack() Stacks {
	ss := Stacks{}

	s1lit := []string{"L", "N", "W", "T", "D"}
	s2lit := []string{"C", "P", "H"}
	s3lit := []string{"W", "P", "H", "N", "D", "G", "M", "J"}
	s4lit := []string{"C", "W", "S", "N", "T", "Q", "L"}
	s5lit := []string{"P", "H", "C", "N"}
	s6lit := []string{"T", "H", "N", "D", "M", "W", "Q", "B"}
	s7lit := []string{"M", "B", "R", "J", "G", "S", "L"}
	s8lit := []string{"Z", "N", "W", "G", "V", "B", "R", "T"}
	s9lit := []string{"W", "G", "D", "N", "P", "L"}

	sss := [][]string{s1lit, s2lit, s3lit, s4lit, s5lit, s6lit, s7lit, s8lit, s9lit}

	for j := 0; j < len(sss); j++ {
		s := NewStack()
		for i := 0; i < len(sss[j]); i++ {

			s.Push(&Crate{sss[j][i]})
		}
		ss.items = append(ss.items, s)
	}
	return ss
}

type Stacks struct {
	items []*Stack
}

type Moves struct {
	quantity int
	from     int
	to       int
}

type Crate struct {
	Value string
}

func (c *Crate) String() string {
	return fmt.Sprint(c.Value)
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	crates []*Crate
	count  int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Crate) {
	s.crates = append(s.crates[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Crate {
	if s.count == 0 {
		return nil
	}
	s.count--
	crateToPop := s.crates[s.count]
	s.crates = s.crates[:s.count]
	return crateToPop
}
