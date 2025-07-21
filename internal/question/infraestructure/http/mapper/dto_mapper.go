package mapper

import (
	"fmt"

	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/entities"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/valueobjects"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/infraestructure/http/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapCreateQuestionRequest(req request.CreateQuestionRequest) (entities.Question, error) {
	questionType, err := valueobjects.NewQuestionType(req.Type)
	if err != nil {
		return entities.Question{}, fmt.Errorf("tipo de pregunta inválido: %w", err)
	}

	difficulty, err := valueobjects.NewDifficulty(req.Difficulty)
	if err != nil {
		return entities.Question{}, fmt.Errorf("dificultad inválida: %w", err)
	}

	id := req.ID
	if id == "" {
		id = primitive.NewObjectID().Hex()
	}

	isActive := req.IsActive
	if !req.IsActive {
		isActive = true // Por defecto activo
	}

	return entities.Question{
		ID:         id,
		StoryID:    req.StoryID,
		Question:   req.Question,
		Answer:     req.Answer,
		Type:       questionType,
		Difficulty: difficulty,
		Points:     req.Points,
		IsActive:   isActive,
		Options:    req.Options,
	}, nil
}

func MapUpdateQuestionRequest(req request.UpdateQuestionRequest) (entities.Question, error) {
	questionType, err := valueobjects.NewQuestionType(req.Type)
	if err != nil {
		return entities.Question{}, fmt.Errorf("tipo de pregunta inválido: %w", err)
	}

	difficulty, err := valueobjects.NewDifficulty(req.Difficulty)
	if err != nil {
		return entities.Question{}, fmt.Errorf("dificultad inválida: %w", err)
	}

	return entities.Question{
		Question:   req.Question,
		Answer:     req.Answer,
		Type:       questionType,
		Difficulty: difficulty,
		Points:     req.Points,
		IsActive:   req.IsActive,
		Options:    req.Options,
	}, nil
}
