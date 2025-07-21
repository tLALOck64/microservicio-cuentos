package valueobjects

import (
	"errors"
	"fmt"
)

type QuestionType string

const (
	MultipleChoice QuestionType = "multiple_choice"
	TrueFalse      QuestionType = "true_false"
	OpenEnded      QuestionType = "open_ended"
	FillBlank      QuestionType = "fill_blank"
)

func NewQuestionType(questionType string) (QuestionType, error) {
	switch QuestionType(questionType) {
	case MultipleChoice, TrueFalse, OpenEnded, FillBlank:
		return QuestionType(questionType), nil
	default:
		return "", errors.New(fmt.Sprintf("tipo de pregunta inválido: %s", questionType))
	}
}

func (qt QuestionType) String() string {
	return string(qt)
}

func (qt QuestionType) IsMultipleChoice() bool {
	return qt == MultipleChoice
}

func (qt QuestionType) IsTrueFalse() bool {
	return qt == TrueFalse
}

func (qt QuestionType) IsOpenEnded() bool {
	return qt == OpenEnded
}

func (qt QuestionType) IsFillBlank() bool {
	return qt == FillBlank
}
