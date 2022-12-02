package day2

import (
	"fmt"
	"strings"

	"github.com/wizact/advent-of-code-2022/file"
)

func Day2a() {
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
		cr := transformInput(t[0], t[1])

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

func transformInput(i, j string) Round {
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
	case "X":
		round.myHand = rockString
	case "Y":
		round.myHand = paperString
	case "Z":
		round.myHand = scissorsString
	}

	return round
}
