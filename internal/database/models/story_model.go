package models

import (
	"time"

	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/aggregates"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/entities"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/valueobjects"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// StoryModel representa la estructura de un cuento en MongoDB
type StoryModel struct {
	ID          primitive.ObjectID     `bson:"_id,omitempty"`
	Title       string                 `bson:"title"`
	Description string                 `bson:"description"`
	Language    string                 `bson:"language"`
	Category    string                 `bson:"category"`
	ContentJSON map[string]interface{} `bson:"content_json"`
	Status      string                 `bson:"status"`
	CreatedAt   time.Time              `bson:"created_at"`
	UpdatedAt   time.Time              `bson:"updated_at"`
	PublishedAt *time.Time             `bson:"published_at,omitempty"`
}

// ToDomainStoryAggregate convierte el modelo MongoDB a agregado de dominio
func (sm *StoryModel) ToDomainStoryAggregate() (*aggregates.StoryAggregate, error) {
	// Convertir strings a value objects
	language, err := valueobjects.NewLanguage(sm.Language)
	if err != nil {
		return nil, err
	}

	category, err := valueobjects.NewCategory(sm.Category)
	if err != nil {
		return nil, err
	}

	// Crear el agregado usando el constructor
	aggregate, err := aggregates.NewStoryAggregate(
		sm.ID.Hex(),
		sm.Title,
		sm.Description,
		language,
		category,
	)
	if err != nil {
		return nil, err
	}

	// Actualizar el contenido si existe
	if len(sm.ContentJSON) > 0 {
		err = aggregate.UpdateContent(sm.ContentJSON)
		if err != nil {
			return nil, err
		}
	}

	// Si el estado no es Draft, actualizarlo según corresponda
	if sm.Status == string(valueobjects.Published) && sm.PublishedAt != nil {
		err = aggregate.Publish()
		if err != nil {
			return nil, err
		}
	} else if sm.Status == string(valueobjects.Archived) {
		err = aggregate.Archive()
		if err != nil {
			return nil, err
		}
	}

	// Limpiar eventos ya que estamos reconstruyendo desde la base de datos
	aggregate.ClearEvents()

	return aggregate, nil
}

// FromDomainStoryAggregate convierte agregado de dominio a modelo MongoDB
func FromDomainStoryAggregate(aggregate *aggregates.StoryAggregate) *StoryModel {
	var objectID primitive.ObjectID

	// Si tiene ID, convertir de string a ObjectID
	if aggregate.ID != "" {
		if oid, err := primitive.ObjectIDFromHex(aggregate.ID); err == nil {
			objectID = oid
		}
	}

	model := &StoryModel{
		ID:          objectID,
		Title:       aggregate.Title,
		Description: aggregate.Description,
		Language:    aggregate.Language.String(),
		Category:    aggregate.Category.String(),
		ContentJSON: aggregate.ContentJSON,
		Status:      aggregate.Status.String(),
		CreatedAt:   aggregate.GetCreatedAt(),
		UpdatedAt:   aggregate.GetUpdatedAt(),
		PublishedAt: aggregate.GetPublishedAt(),
	}

	return model
}

// NewStoryModel crea un nuevo modelo de cuento para MongoDB
func NewStoryModel(title, description, language, category string, content map[string]interface{}) *StoryModel {
	return &StoryModel{
		Title:       title,
		Description: description,
		Language:    language,
		Category:    category,
		ContentJSON: content,
		Status:      string(valueobjects.Draft),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// UpdateContent actualiza el contenido del modelo
func (sm *StoryModel) UpdateContent(content map[string]interface{}) {
	sm.ContentJSON = content
	sm.UpdatedAt = time.Now()
}

// Publish marca el cuento como publicado
func (sm *StoryModel) Publish() {
	now := time.Now()
	sm.Status = string(valueobjects.Published)
	sm.PublishedAt = &now
	sm.UpdatedAt = now
}

// Archive marca el cuento como archivado
func (sm *StoryModel) Archive() {
	sm.Status = string(valueobjects.Archived)
	sm.UpdatedAt = time.Now()
}

// ToDomainStory convierte el modelo MongoDB a entidad de dominio
func (sm *StoryModel) ToDomainStory() (*entities.Story, error) {
	// Convertir strings a value objects
	language, err := valueobjects.NewLanguage(sm.Language)
	if err != nil {
		return nil, err
	}

	category, err := valueobjects.NewCategory(sm.Category)
	if err != nil {
		return nil, err
	}

	status, err := valueobjects.NewStatus(sm.Status)
	if err != nil {
		return nil, err
	}

	return &entities.Story{
		ID:          sm.ID.Hex(),
		Title:       sm.Title,
		Description: sm.Description,
		Language:    language,
		Category:    category,
		ContentJSON: sm.ContentJSON,
		Status:      status,
	}, nil
}

// FromDomainStory convierte entidad de dominio a modelo MongoDB
func FromDomainStory(story *entities.Story) *StoryModel {
	var objectID primitive.ObjectID

	// Si tiene ID, convertir de string a ObjectID
	if story.ID != "" {
		if oid, err := primitive.ObjectIDFromHex(story.ID); err == nil {
			objectID = oid
		}
	}

	return &StoryModel{
		ID:          objectID,
		Title:       story.Title,
		Description: story.Description,
		Language:    story.Language.String(),
		Category:    story.Category.String(),
		ContentJSON: story.ContentJSON,
		Status:      story.Status.String(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// ToDomainStory función de conveniencia a nivel de paquete
func ToDomainStory(model *StoryModel) (*entities.Story, error) {
	return model.ToDomainStory()
}
