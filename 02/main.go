package main

import (
	"aoc2022/utils"
	"fmt"
	"strings"
)

func main() {
	lines, err := utils.ReadFile("02", false)
	if err != nil {
		return
	}

	ROCK := 0
	PAPER := 1
	SCISSORS := 2

	LOSE := 0
	DRAW := 3
	WIN := 6

	opponentMappings := map[string]int{"A": ROCK, "B": PAPER, "C": SCISSORS}
	ownMappings := map[string]int{"X": ROCK, "Y": PAPER, "Z": SCISSORS}
	part2ScoreMapping := map[string]int{"X": LOSE, "Y": DRAW, "Z": WIN}

	score1 := 0
	score2 := 0
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " ")

		opponentGame := opponentMappings[parts[0]]
		ownGame := ownMappings[parts[1]]

		score1 += ownGame + 1
		if opponentGame == ownGame {
			score1 += DRAW
		} else if (opponentGame == ROCK && ownGame == PAPER) || (opponentGame == PAPER && ownGame == SCISSORS) || (opponentGame == SCISSORS && ownGame == ROCK) {
			score1 += WIN
		}

		part2GameScore := part2ScoreMapping[parts[1]]
		score2 += part2GameScore
		ownGame2 := 0
		if part2GameScore == LOSE {
			ownGame2 = opponentGame - 1
			if opponentGame == ROCK {
				ownGame2 = SCISSORS
			}
		} else if part2GameScore == DRAW {
			ownGame2 = opponentGame
		} else {
			ownGame2 = (opponentGame + 1) % 3
		}

		score2 += ownGame2 + 1
	}

	fmt.Println("Part 1: ", score1)
	fmt.Println("Part 2: ", score2)
}
