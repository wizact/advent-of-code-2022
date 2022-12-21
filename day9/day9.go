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
		move(ins[i], &headPosition, &tailPosition)
		if !tailAlreadyBeenInPosition(arr, &tailPosition) {
			arr = append(arr, Postion{x: tailPosition.x, y: tailPosition.y})
		}
	}

	fmt.Println(len(arr))
}

func move(ins Instruction, hp *Postion, tp *Postion) {
	fmt.Println(ins.direction, ins.steps)
	adjustPosition(ins, hp)

	calcTailRelativePostion(hp, tp, ins)
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
		// No adjustment is required
		fmt.Println("overlapping")
		return
	}

	touchingX := absXDistance == 1
	touchingY := absYDistance == 1

	if ((touchingX || touchingY) && !(touchingX && touchingY)) && absXDistance < 2 && absYDistance < 2 {
		// No adjustment is required
		fmt.Println("touching adjacent", hp.x, hp.y, tp.x, tp.y, xDistance, yDistance, touchingX, touchingY)
		return
	}

	if touchingX && touchingY {
		// No adjustment is required
		fmt.Println("touching diagonally", hp.x, hp.y, tp.x, tp.y, xDistance, yDistance, touchingX, touchingY)
		return
	}

	// Tail is two or more steps away from head either up, down, left, or right. Tail must also move x steps in that direction
	spaceX := absXDistance >= 2 // not on the same col
	spaceY := absYDistance >= 2 // not on the same row

	if ((spaceX || spaceY) && !(spaceX && spaceY)) && absXDistance == 0 || absYDistance == 0 {
		// Adjustment is required. Move tail to the same direction as head has moved by x steps
		fmt.Println("Distant but on the same row or column", hp.x, hp.y, tp.x, tp.y, xDistance, yDistance, spaceX, spaceY)
		if spaceX {
			// Adjust X
			if xDistance > 0 {
				newTailInstruction.direction = RIGHT
			} else if xDistance < 0 {
				newTailInstruction.direction = LEFT
			} else {
				panic("should not happen")
			}
			newTailInstruction.steps = int(absXDistance) - 1

		} else if spaceY {
			// Adjust Y
			if yDistance > 0 {
				newTailInstruction.direction = UP
			} else if yDistance < 0 {
				newTailInstruction.direction = DOWN
			} else {
				panic("should not happen")
			}
			newTailInstruction.steps = int(absYDistance) - 1
		} else {
			panic("should not happen")
		}
		adjustPosition(*newTailInstruction, tp)
		fmt.Println("New  tail poisition", tp.x, tp.y)
		return
	}

	// Adjustment is required. Move tail to the same direction as head has moved by x steps
	fmt.Println("Distant diagonally", hp.x, hp.y, tp.x, tp.y, xDistance, yDistance, touchingX, touchingY)

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

	fmt.Println("New  tail poisition", tp.x, tp.y)

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
