package day9

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/wizact/advent-of-code-2022/file"
)

const UP string = "U"
const DOWN string = "D"
const LEFT string = "L"
const RIGHT string = "R"

func Day9() {
	ins := getInstructions()
	trackMovement(ins)
}

type Position struct {
	x int
	y int
}

func trackMovement(ins []Instruction) {
	noOfKnots := 2
	arr := []Position{}

	knots := make(map[int]Position, noOfKnots)
	for i := 0; i < noOfKnots; i++ {
		knots[i] = Position{x: 0, y: 0}
	}

	for i := 0; i < len(ins); i++ {
		arr = move(ins[i], knots, arr)
	}

	fmt.Println("Number of unique positions tail has been at least once:", len(arr))
}

func move(ins Instruction, knots map[int]Position, arr []Position) []Position {
	return stepThroughInstruction(ins, knots, arr)
}

func stepThroughInstruction(ins Instruction, knots map[int]Position, arr []Position) []Position {
	for i := 0; i < ins.steps; i++ {
		head := knots[0]
		adjustPosition(Instruction{direction: ins.direction, steps: 1}, &head)
		knots[0] = head
		for j := 0; j < len(knots)-1; j++ {

			k1 := knots[j]
			k2 := knots[j+1]

			calcTailRelativePostion(&k1, &k2, ins)

			if j+2 == len(knots) {
				if !tailAlreadyBeenInPosition(arr, &k2) {
					arr = append(arr, Position{x: k2.x, y: k2.y})
				}
			}

			knots[j+1] = k2
		}
	}

	return arr
}

func adjustPosition(ins Instruction, p *Position) {
	switch ins.direction {
	case UP:
		p.y = p.y + ins.steps
	case DOWN:
		p.y = p.y - ins.steps
	case LEFT:
		p.x = p.x - ins.steps
	case RIGHT:
		p.x = p.x + ins.steps
	}
}

func calcTailRelativePostion(hp, tp *Position, ins Instruction) {
	newTailInstruction := &Instruction{}
	xDistance := hp.x - tp.x
	yDistance := hp.y - tp.y
	absXDistance := math.Abs(float64(xDistance))
	absYDistance := math.Abs(float64(yDistance))

	if absXDistance < 2 && absYDistance < 2 {
		return
	}

	// Tail is two or more steps away from head either up, down, left, or right. Tail must also move x steps in that direction
	spaceX := absXDistance >= 2 // there is distance between them on x axis.
	spaceY := absYDistance >= 2 // there is distance between them on y axis.

	if ((spaceX || spaceY) && !(spaceX && spaceY)) && absXDistance == 0 || absYDistance == 0 {
		// Distant but on the same row or column.
		// Adjustment is required. Move tail to the same direction as head has moved by x steps
		if spaceX {
			// Adjust X
			if xDistance > 0 {
				newTailInstruction.direction = RIGHT
			} else if xDistance < 0 {
				newTailInstruction.direction = LEFT
			}
			newTailInstruction.steps = int(absXDistance) - 1

		} else if spaceY {
			// Adjust Y
			if yDistance > 0 {
				newTailInstruction.direction = UP
			} else if yDistance < 0 {
				newTailInstruction.direction = DOWN
			}
			newTailInstruction.steps = int(absYDistance) - 1
		}
	} else {
		// Distant diagonally. Adjustment is required. Move tail to the same direction as head has moved by x steps
		newTailInstruction.steps = 1
		newTailInstruction.direction = ins.direction
		if ins.direction == UP || ins.direction == DOWN {
			tp.x = hp.x
		} else {
			tp.y = hp.y
		}
	}
	adjustPosition(*newTailInstruction, tp)
}

func tailAlreadyBeenInPosition(arr []Position, tp *Position) bool {
	for _, v := range arr {
		if v.x == tp.x && v.y == tp.y {
			return true
		}
	}
	return false
}

func getInstructions() []Instruction {
	b := getFile()
	c := string(b)
	ms := strings.Split(c, "\n")

	ins := []Instruction{}
	for i := 0; i < len(ms); i++ {
		inss := strings.Split(ms[i], " ")

		d := inss[0]
		s, e := strconv.Atoi(inss[1])

		if e != nil {
			panic(e)
		}

		ins = append(ins, Instruction{direction: d, steps: s})
	}

	return ins
}

func getFile() []byte {
	b, e := file.ReadFile("./day9/day9.txt")

	if e != nil {
		panic(e)
	}
	return b
}

type Instruction struct {
	direction string
	steps     int
}
