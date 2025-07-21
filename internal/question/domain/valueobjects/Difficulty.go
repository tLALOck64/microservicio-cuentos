package valueobjects

import (
	"fmt"
)

type Difficulty string

const (
	Easy   Difficulty = "easy"
	Medium Difficulty = "medium"
	Hard   Difficulty = "hard"
)

func NewDifficulty(difficulty string) (Difficulty, error) {
	switch Difficulty(difficulty) {
	case Easy, Medium, Hard:
		return Difficulty(difficulty), nil
	default:
		return "", fmt.Errorf("dificultad inválida: %s", difficulty)
	}
}

func (d Difficulty) String() string {
	return string(d)
}

func (d Difficulty) IsEasy() bool {
	return d == Easy
}

func (d Difficulty) IsMedium() bool {
	return d == Medium
}

func (d Difficulty) IsHard() bool {
	return d == Hard
}
