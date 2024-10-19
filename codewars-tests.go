package main

import (
	"fmt"
	"strings"
)

// Kata: Grading

// 90 <= score <= 100	'A'
// 80 <= score < 90	'B'
// 70 <= score < 80	'C'
// 60 <= score < 70	'D'
// 0 <= score < 60	'F'
func GetGrade(a, b, c int) rune {
	switch mean := float32(a+b+c) / 3; {
	case mean < 60:
		return 'F'
	case mean < 70:
		return 'D'
	case mean < 80:
		return 'C'
	case mean < 90:
		return 'B'
	default:
		return 'A'
	}
}

// Kata Rock Paper Scissors
const (
	Player1Won = "Player 1 won!"
	Player2Won = "Player 2 won!"
	Draw       = "Draw!"
)

var RpsPlays = map[string]string{
	"scissors - paper": Player1Won,
	"scissors - rock":  Player2Won,
	"rock - scissors":  Player1Won,
	"rock - paper":     Player2Won,
	"paper - scissors": Player2Won,
	"paper - rock":     Player1Won,
}

func Rps(p1, p2 string) string {
	player1 := strings.Trim(strings.ToLower(p1), " ")
	player2 := strings.Trim(strings.ToLower(p2), " ")
	if player1 == player2 {
		return Draw
	}

	play := fmt.Sprintf("%s - %s", player1, player2)
	result, ok := RpsPlays[play]
	if ok {
		return result
	}

	return fmt.Sprintf("UnKnown combo %s - %s", p1, p2)
}

var m = map[string]string{"rock": "paper", "paper": "scissors", "scissors": "rock"}

// Rps2 provides a more elegant solution
func Rps2(a, b string) string {
	if a == b {
		return "Draw!"
	}
	if m[a] == b {
		return "Player 2 won!"
	}
	return "Player 1 won!"
}

func main() {
	//fmt.Println(string(GetGrade(5, 20, 30)))
	//fmt.Println(Rps2("rock", "scissors"))
}
