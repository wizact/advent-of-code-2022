package day2

import (
	"fmt"
	"strings"

	"github.com/wizact/advent-of-code-2022/file"
)

func Day2b() {
	b, e := file.ReadFile("./day2/day2a.txt")

	if e != nil {
		panic(e)
	}
	i := string(b)
	rounds := strings.Split(i, "\n")
	noOfRounds := len(rounds)
	m := make(map[int]Round, noOfRounds)

	for k, v := range rounds {
		t := strings.Split(v, " ")
		cr := transformNewInput(t[0], t[1])

		calcWinningResult(&cr)
		calcMultiplier(&cr)
		m[k] = cr

		fmt.Printf("My hand is %v (%v), You hand is %v (%v). I %v, ml %v\n", m[k].rawMyHand, m[k].myHand, m[k].rawYourHand, m[k].yourHand, m[k].winningResult, m[k].winningMultiplier)
	}

	totalScore := 0
	for _, v := range m {
		totalScore += v.winningMultiplier + v.winningResult
	}

	fmt.Println(totalScore)
}

func transformNewInput(i, j string) Round {
	round := Round{rawMyHand: i, rawYourHand: j}
	/*
		A Rock
		B Paper
		C Scissors

		X Rock
		Y Paper
		Z Scissors
	*/

	switch i {
	case "A":
		round.yourHand = rockString
	case "B":
		round.yourHand = paperString
	case "C":
		round.yourHand = scissorsString
	}

	switch j {
	case "X": // need to lose
		round.myHand = chooseMyHand(round.yourHand, lostString)
	case "Y": // need to draw
		round.myHand = chooseMyHand(round.yourHand, drawString)
	case "Z": // need to win
		round.myHand = chooseMyHand(round.yourHand, wonString)
	}

	return round
}

func chooseMyHand(yourHand string, desiredResult string) string {
	switch desiredResult {
	case wonString:
		return mustWin(yourHand)
	case drawString:
		return mustDraw(yourHand)
	case lostString:
		return mustLose(yourHand)
	}

	panic("Input is inconclusive")
}

func mustWin(yourHand string) string {
	// scissor > rock
	// paper > scissor
	// rock > paper
	switch yourHand {
	case scissorsString:
		return rockString
	case paperString:
		return scissorsString
	case rockString:
		return paperString
	default:
		return ""
	}
}

func mustLose(yourHand string) string {
	// paper > rock
	// scissors > paper
	// rock > scissors
	switch yourHand {
	case paperString:
		return rockString
	case scissorsString:
		return paperString
	case rockString:
		return scissorsString
	default:
		return ""
	}
}

func mustDraw(yourHand string) string {
	// scissor > scissor
	// paper > paper
	// rock > rock
	return yourHand
}
