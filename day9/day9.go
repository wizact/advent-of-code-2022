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
	drawMap(ins)
	trackMovement(ins)
}

type Postion struct {
	x int
	y int
}

func trackMovement(ins []Instruction) {
	arr := []Postion{}
	headPosition := Postion{x: 0, y: 0}
	tailPosition := Postion{x: 0, y: 0}

	for i := 0; i < len(ins); i++ {
		arr = move(ins[i], &headPosition, &tailPosition, arr)
	}

	fmt.Println("Number of unique positions tail has been at least once:", len(arr))
}

func move(ins Instruction, hp *Postion, tp *Postion, arr []Postion) []Postion {
	return stepThroughInstruction(ins, hp, tp, arr)
}

func stepThroughInstruction(ins Instruction, hp *Postion, tp *Postion, arr []Postion) []Postion {
	for i := 0; i < ins.steps; i++ {
		adjustPosition(Instruction{direction: ins.direction, steps: 1}, hp)
		calcTailRelativePostion(hp, tp, ins)
		if !tailAlreadyBeenInPosition(arr, tp) {
			arr = append(arr, Postion{x: tp.x, y: tp.y})
		}
	}
	return arr
}

func adjustPosition(ins Instruction, p *Postion) {
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

func calcTailRelativePostion(hp, tp *Postion, ins Instruction) {
	newTailInstruction := &Instruction{}
	xDistance := hp.x - tp.x
	yDistance := hp.y - tp.y
	absXDistance := math.Abs(float64(xDistance))
	absYDistance := math.Abs(float64(yDistance))

	if xDistance == 0 && yDistance == 0 {
		// Overlapping, no adjustment is required
		return
	}

	touchingX := absXDistance == 1
	touchingY := absYDistance == 1

	if ((touchingX || touchingY) && !(touchingX && touchingY)) && absXDistance < 2 && absYDistance < 2 {
		// Touching adjacent, no adjustment is required
		return
	}

	if touchingX && touchingY {
		// Touching diagonally, no adjustment is required
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
		adjustPosition(*newTailInstruction, tp)
		return
	}

	// Distant diagonally. Adjustment is required. Move tail to the same direction as head has moved by x steps
	if ins.direction == UP || ins.direction == DOWN {
		newTailInstruction.steps = int(math.Abs(float64(hp.y-tp.y)) - 1)
		newTailInstruction.direction = ins.direction
		tp.x = hp.x
	} else {
		newTailInstruction.steps = int(math.Abs(float64(hp.x-tp.x)) - 1)
		newTailInstruction.direction = ins.direction
		tp.y = hp.y
	}

	adjustPosition(*newTailInstruction, tp)
}

func tailAlreadyBeenInPosition(arr []Postion, tp *Postion) bool {
	for _, v := range arr {
		if v.x == tp.x && v.y == tp.y {
			return true
		}
	}
	return false
}

func drawMap(ins []Instruction) {
	var u, d, l, r int

	for i := 0; i < len(ins); i++ {
		switch ins[i].direction {
		case UP:
			u = u + ins[i].steps
		case DOWN:
			d = d - ins[i].steps
		case LEFT:
			l = l - ins[i].steps
		case RIGHT:
			r = r + ins[i].steps
		}
	}

	fmt.Println("Total movement up:", u, "down:", d, "left:", l, "right:", r)

	x := r + l
	y := u + d
	fmt.Println("Total movement of head horizontally:", x, "and vertically", y)
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
