package mapper

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/entities"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/valueobjects"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/infraestructure/http/request"
)

func MapCreateStoryRequest(req request.CreateStoryRequest) (entities.Story, error) {
	language, err := valueobjects.NewLanguage(req.Language)

	if err != nil {
		return entities.Story{}, fmt.Errorf("tipo inválido: %w", err)
	}

	category, err := valueobjects.NewCategory(req.Category)

	if err != nil {
		return entities.Story{}, fmt.Errorf("tipo inválido: %w", err)
	}

	status, err := valueobjects.NewStatus(req.Status)
	if err != nil {
		return entities.Story{}, fmt.Errorf("tipo inválido: %w", err)
	}

	id := req.ID
	if id == "" {
		id = primitive.NewObjectID().Hex()
	}

	return entities.Story{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		Language:    language,
		Category:    category,
		ContentJSON: req.ContentJSON,
		Status:      status,
	}, nil
}
