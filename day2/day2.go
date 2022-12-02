package day2

const (
	lost = 0
	draw = 3
	won  = 6

	rockMultiplier     = 1
	paperMultiplier    = 2
	scissorsMultiplier = 3

	rockString     = "rock"
	paperString    = "paper"
	scissorsString = "scissors"

	lostString = "lost"
	drawString = "draw"
	wonString  = "won"
)

type Round struct {
	rawMyHand         string
	rawYourHand       string
	myHand            string
	yourHand          string
	winningResult     int
	winningMultiplier int
}

func calcMultiplier(round *Round) {
	switch round.myHand {
	case rockString:
		round.winningMultiplier = rockMultiplier
	case paperString:
		round.winningMultiplier = paperMultiplier
	case scissorsString:
		round.winningMultiplier = scissorsMultiplier
	}
}

func calcWinningResult(round *Round) {
	m := round.myHand
	y := round.yourHand

	if m == y {
		round.winningResult = draw
	}

	// rock scissors = won
	// scissors paper = won
	// paper rock = won
	// rock paper = lost
	// paper scissors = lost
	// scissors rock = lost

	if (m == rockString && y == scissorsString) ||
		(m == scissorsString && y == paperString) ||
		(m == paperString && y == rockString) {
		round.winningResult = won
	}

	if (m == rockString && y == paperString) ||
		(m == paperString && y == scissorsString) ||
		(m == scissorsString && y == rockString) {
		round.winningResult = lost
	}
}
