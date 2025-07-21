package aggregates

import (
	"errors"
	"time"

	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/entities"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/events"
	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/valueobjects"
)

type QuestionAggregate struct {
	entities.Question
	createdAt time.Time
	updatedAt time.Time
	events    []interface{}
}

func NewQuestionAggregate(id, storyID, question, answer string, questionType valueobjects.QuestionType, difficulty valueobjects.Difficulty, points int) (*QuestionAggregate, error) {
	if question == "" {
		return nil, errors.New("la pregunta es requerida")
	}
	if answer == "" {
		return nil, errors.New("la respuesta es requerida")
	}
	if storyID == "" {
		return nil, errors.New("el ID de la historia es requerido")
	}
	if points < 0 {
		return nil, errors.New("los puntos no pueden ser negativos")
	}

	now := time.Now()
	aggregate := &QuestionAggregate{
		Question: entities.Question{
			ID:         id,
			StoryID:    storyID,
			Question:   question,
			Answer:     answer,
			Type:       questionType,
			Difficulty: difficulty,
			Points:     points,
			IsActive:   true,
		},
		createdAt: now,
		updatedAt: now,
		events:    []interface{}{},
	}

	aggregate.events = append(aggregate.events, events.QuestionCreated{
		QuestionID: id,
		StoryID:    storyID,
		Type:       questionType.String(),
		Difficulty: difficulty.String(),
		CreatedAt:  now,
	})

	return aggregate, nil
}

// Deactivate desactiva la pregunta
func (q *QuestionAggregate) Deactivate() error {
	if !q.IsActive {
		return errors.New("la pregunta ya está desactivada")
	}

	now := time.Now()
	q.IsActive = false
	q.updatedAt = now

	q.events = append(q.events, events.QuestionDeactivated{
		QuestionID:    q.ID,
		StoryID:       q.StoryID,
		DeactivatedAt: now,
	})

	return nil
}

// Activate activa la pregunta
func (q *QuestionAggregate) Activate() error {
	if q.IsActive {
		return errors.New("la pregunta ya está activa")
	}

	now := time.Now()
	q.IsActive = true
	q.updatedAt = now

	q.events = append(q.events, events.QuestionActivated{
		QuestionID:  q.ID,
		StoryID:     q.StoryID,
		ActivatedAt: now,
	})

	return nil
}

// UpdateQuestion actualiza la pregunta
func (q *QuestionAggregate) UpdateQuestion(question, answer string) error {
	if question == "" {
		return errors.New("la pregunta no puede estar vacía")
	}
	if answer == "" {
		return errors.New("la respuesta no puede estar vacía")
	}

	q.Question.Question = question
	q.Question.Answer = answer
	q.updatedAt = time.Now()

	return nil
}

// UpdatePoints actualiza los puntos de la pregunta
func (q *QuestionAggregate) UpdatePoints(points int) error {
	if points < 0 {
		return errors.New("los puntos no pueden ser negativos")
	}

	q.Points = points
	q.updatedAt = time.Now()

	return nil
}

// GetEvents obtiene los eventos del agregado
func (q *QuestionAggregate) GetEvents() []interface{} {
	return q.events
}

// ClearEvents limpia los eventos del agregado
func (q *QuestionAggregate) ClearEvents() {
	q.events = []interface{}{}
}

// GetCreatedAt obtiene la fecha de creación
func (q *QuestionAggregate) GetCreatedAt() time.Time {
	return q.createdAt
}

// GetUpdatedAt obtiene la fecha de actualización
func (q *QuestionAggregate) GetUpdatedAt() time.Time {
	return q.updatedAt
}
