package entities

import "github.com/tLALOck64/microservicio-cuentos/internal/question/domain/valueobjects"

type Question struct {
	ID         string
	StoryID    string
	Question   string
	Answer     string
	Type       valueobjects.QuestionType
	Difficulty valueobjects.Difficulty
	Points     int
	IsActive   bool
	Options    []string
}
