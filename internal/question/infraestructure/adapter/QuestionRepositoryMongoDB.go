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
	"github.com/tLALOck64/microservicio-cuentos/internal/question/domain/entities"
)

type QuestionRepositoryMongoDB struct {
	DB *database.MongoDB
}

// NewQuestionRepositoryMongoDB crea una nueva instancia del repositorio MongoDB
func NewQuestionRepositoryMongoDB() (*QuestionRepositoryMongoDB, error) {
	db, err := database.Connect()
	if err != nil {
		panic("Error connecting to MongoDB: " + err.Error())
	}

	return &QuestionRepositoryMongoDB{DB: db}, nil
}

// Create crea una nueva pregunta en la base de datos
func (r *QuestionRepositoryMongoDB) Create(question *entities.Question) (*entities.Question, error) {
	collection := r.DB.Database.Collection("questions")

	model := models.FromDomainQuestion(question)

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
		return nil, fmt.Errorf("error al crear pregunta: %w", err)
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		model.ID = oid
	}

	question.ID = model.ID.Hex()

	return question, nil
}

// GetById obtiene una pregunta por su ID
func (r *QuestionRepositoryMongoDB) GetById(id string) (*entities.Question, error) {
	collection := r.DB.Database.Collection("questions")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("ID inválido: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var model models.QuestionModel
	filter := bson.M{"_id": objectID}

	err = collection.FindOne(ctx, filter).Decode(&model)
	if err != nil {
		return nil, fmt.Errorf("error al obtener pregunta: %w", err)
	}

	questionEntity, err := model.ToDomainQuestion()
	if err != nil {
		return nil, err
	}

	return questionEntity, nil
}

// Get obtiene todas las preguntas activas
func (r *QuestionRepositoryMongoDB) Get() ([]*entities.Question, error) {
	collection := r.DB.Database.Collection("questions")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filtro para obtener solo preguntas activas
	filter := bson.M{"is_active": true}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error al obtener preguntas: %w", err)
	}
	defer cursor.Close(ctx)

	var modelos []models.QuestionModel
	if err = cursor.All(ctx, &modelos); err != nil {
		return nil, fmt.Errorf("error al decodificar preguntas: %w", err)
	}

	// Convertir modelos a entidades de dominio
	questions := make([]*entities.Question, 0, len(modelos))
	for _, model := range modelos {
		question, err := model.ToDomainQuestion()
		if err != nil {
			log.Printf("Error al convertir pregunta %s: %v", model.ID.Hex(), err)
			continue
		}
		questions = append(questions, question)
	}

	return questions, nil
}

// GetByStoryId obtiene todas las preguntas de una historia específica
func (r *QuestionRepositoryMongoDB) GetByStoryId(storyID string) ([]*entities.Question, error) {
	collection := r.DB.Database.Collection("questions")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Filtro para obtener preguntas de una historia específica y activas
	filter := bson.M{
		"story_id":  storyID,
		"is_active": true,
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("error al obtener preguntas de la historia: %w", err)
	}
	defer cursor.Close(ctx)

	var modelos []models.QuestionModel
	if err = cursor.All(ctx, &modelos); err != nil {
		return nil, fmt.Errorf("error al decodificar preguntas: %w", err)
	}

	// Convertir modelos a entidades de dominio
	questions := make([]*entities.Question, 0, len(modelos))
	for _, model := range modelos {
		question, err := model.ToDomainQuestion()
		if err != nil {
			log.Printf("Error al convertir pregunta %s: %v", model.ID.Hex(), err)
			continue
		}
		questions = append(questions, question)
	}

	return questions, nil
}

// Update actualiza una pregunta existente
func (r *QuestionRepositoryMongoDB) Update(question *entities.Question) (*entities.Question, error) {
	collection := r.DB.Database.Collection("questions")

	objectID, err := primitive.ObjectIDFromHex(question.ID)
	if err != nil {
		return nil, fmt.Errorf("ID inválido: %w", err)
	}

	model := models.FromDomainQuestion(question)
	model.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": model}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("error al actualizar pregunta: %w", err)
	}

	return question, nil
}

// Delete elimina una pregunta (soft delete cambiando is_active a false)
func (r *QuestionRepositoryMongoDB) Delete(id string) error {
	collection := r.DB.Database.Collection("questions")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("ID inválido: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"is_active": false, "updated_at": time.Now()}}

	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error al eliminar pregunta: %w", err)
	}

	return nil
}
