package aggregates

import (
	"errors"
	"time"

	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/entities"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/events"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/valueobjects"
)

type StoryAggregate struct {
	entities.Story
	createdAt   time.Time
	updatedAt   time.Time
	publishedAt *time.Time
	events      []interface{}
}


func NewStoryAggregate(id, title, description string, language valueobjects.Language, category valueobjects.Category) (*StoryAggregate, error) {
	if title == "" {
		return nil, errors.New("el título es requerido")
	}
	if description == "" {
		return nil, errors.New("la descripción es requerida")
	}

	now := time.Now()
	aggregate := &StoryAggregate{
		Story: entities.Story{
			ID:          id,
			Title:       title,
			Description: description,
			Language:    language,
			Category:    category,
			ContentJSON: make(map[string]interface{}),
			Status:      valueobjects.Draft,
		},
		createdAt: now,
		updatedAt: now,
		events:    []interface{}{},
	}

	aggregate.events = append(aggregate.events, events.StoryCreated{
		StoryID:   id,
		Title:     title,
		Language:  language.String(),
		Category:  category.String(),
		CreatedAt: now,
	})

	return aggregate, nil
}

// Publish publica el cuento
func (s *StoryAggregate) Publish() error {
	if len(s.ContentJSON) == 0 {
		return errors.New("el cuento debe tener contenido")
	}

	if !s.Status.IsDraft() {
		return errors.New("solo cuentos en borrador pueden ser publicados")
	}

	now := time.Now()
	s.Status = valueobjects.Published
	s.publishedAt = &now
	s.updatedAt = now

	// Agregar evento
	s.events = append(s.events, events.StoryPublished{
		StoryID:     s.ID,
		Title:       s.Title,
		Language:    s.Language.String(),
		PublishedAt: now,
	})

	return nil
}

// Archive archiva el cuento
func (s *StoryAggregate) Archive() error {
	if s.Status.IsArchived() {
		return errors.New("el cuento ya está archivado")
	}

	now := time.Now()
	s.Status = valueobjects.Archived
	s.updatedAt = now

	// Agregar evento
	s.events = append(s.events, events.StoryArchived{
		StoryID:    s.ID,
		Title:      s.Title,
		ArchivedAt: now,
	})

	return nil
}

// UpdateContent actualiza el contenido del cuento
func (s *StoryAggregate) UpdateContent(content map[string]interface{}) error {
	if !s.Status.IsDraft() {
		return errors.New("solo se puede actualizar el contenido de cuentos en borrador")
	}

	if len(content) == 0 {
		return errors.New("el contenido no puede estar vacío")
	}

	s.ContentJSON = content
	s.updatedAt = time.Now()

	return nil
}

// GetEvents obtiene los eventos del agregado
func (s *StoryAggregate) GetEvents() []interface{} {
	return s.events
}

// ClearEvents limpia los eventos del agregado
func (s *StoryAggregate) ClearEvents() {
	s.events = []interface{}{}
}

// GetCreatedAt obtiene la fecha de creación
func (s *StoryAggregate) GetCreatedAt() time.Time {
	return s.createdAt
}

// GetUpdatedAt obtiene la fecha de actualización
func (s *StoryAggregate) GetUpdatedAt() time.Time {
	return s.updatedAt
}

// GetPublishedAt obtiene la fecha de publicación
func (s *StoryAggregate) GetPublishedAt() *time.Time {
	return s.publishedAt
}
