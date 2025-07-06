package adapters

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/tLALOck64/microservicio-cuentos/internal/database"
	"github.com/tLALOck64/microservicio-cuentos/internal/database/models"
	"github.com/tLALOck64/microservicio-cuentos/internal/story/domain/entities"
)

type StoryRepositoryMongoDB struct {
	DB *database.MongoDB
}

// NewStoryRepositoryMongoDB crea una nueva instancia del repositorio MongoDB
func NewStoryRepositoryMongoDB() (*StoryRepositoryMongoDB, error) {
	db, err := database.Connect()
	if err != nil {
		panic("Error connecting to MongoDB: " + err.Error())
	}

	return &StoryRepositoryMongoDB{DB: db}, nil
}

// Create crea un nuevo cuento en la base de datos
func (r *StoryRepositoryMongoDB) Create(story *entities.Story) (*entities.Story, error) {
	collection := r.DB.Database.Collection("stories")

	model := models.FromDomainStory(story)

	if model.ID.IsZero() {
		model.ID = primitive.NewObjectID()
	}

	now := time.Now()
	model.CreatedAt = now
	model.UpdatedAt = now

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, model)
	if err != nil {
		return nil, fmt.Errorf("error al crear cuento: %w", err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		model.ID = oid
	}

	story.ID = model.ID.Hex()

	return story, nil
}

// GetById obtiene un cuento por su ID
func (r *StoryRepositoryMongoDB) GetById(id string) (*entities.Story, error) {
	collection := r.DB.Database.Collection("stories")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("ID inválido: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var model models.StoryModel
	filter := bson.M{"_id": objectID}

	err = collection.FindOne(ctx, filter).Decode(&model)
	if err != nil {
		return nil, fmt.Errorf("error al obtener cuento: %w", err)
	}

	return models.ToDomainStory(&model)
}

// Get obtiene todos los cuentos activos
func (r *StoryRepositoryMongoDB) Get() ([]*entities.Story, error) {
	collection := r.DB.Database.Collection("stories")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filtro para obtener solo cuentos que no estén en estado inactivo
	filter := bson.M{
		"status": bson.M{
			"$ne": "inactive", // No incluir cuentos inactivos
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error al obtener cuentos: %w", err)
	}
	defer cursor.Close(ctx)

	var modelos []models.StoryModel
	if err = cursor.All(ctx, &modelos); err != nil {
		return nil, fmt.Errorf("error al decodificar cuentos: %w", err)
	}

	// Convertir modelos a entidades de dominio
	stories := make([]*entities.Story, len(modelos))
	for i, model := range modelos {
		story, err := models.ToDomainStory(&model)
		if err != nil {
			log.Printf("Error al convertir cuento %s: %v", model.ID.Hex(), err)
			continue
		}
		stories[i] = story
	}

	return stories, nil
}
