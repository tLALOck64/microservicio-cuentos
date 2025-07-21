package models

import (
	"time"

	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/entities"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/valueobjects"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuestionModel struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	StoryID    string             `bson:"story_id"`
	Question   string             `bson:"question"`
	Answer     string             `bson:"answer"`
	Type       string             `bson:"type"`
	Difficulty string             `bson:"difficulty"`
	Points     int                `bson:"points"`
	IsActive   bool               `bson:"is_active"`
	Options    []string           `bson:"options"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}

func (qm *QuestionModel) ToDomainQuestion() (*entities.Question, error) {
	typeVO, err := valueobjects.NewQuestionType(qm.Type)
	if err != nil {
		return nil, err
	}
	difficultyVO, err := valueobjects.NewDifficulty(qm.Difficulty)
	if err != nil {
		return nil, err
	}
	return &entities.Question{
		ID:         qm.ID.Hex(),
		StoryID:    qm.StoryID,
		Question:   qm.Question,
		Answer:     qm.Answer,
		Type:       typeVO,
		Difficulty: difficultyVO,
		Points:     qm.Points,
		IsActive:   qm.IsActive,
		Options:    qm.Options,
	}, nil
}

func FromDomainQuestion(question *entities.Question) *QuestionModel {
	var objectID primitive.ObjectID

	if question.ID != "" {
		if oid, err := primitive.ObjectIDFromHex(question.ID); err == nil {
			objectID = oid
		}
	}

	return &QuestionModel{
		ID:         objectID,
		StoryID:    question.StoryID,
		Question:   question.Question,
		Answer:     question.Answer,
		Type:       question.Type.String(),
		Difficulty: question.Difficulty.String(),
		Points:     question.Points,
		IsActive:   question.IsActive,
		Options:    question.Options,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func NewQuestionModel(storyID, question, answer, questionType, difficulty string, points int) *QuestionModel {
	return &QuestionModel{
		StoryID:    storyID,
		Question:   question,
		Answer:     answer,
		Type:       questionType,
		Difficulty: difficulty,
		Points:     points,
		IsActive:   true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
