package database

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	instance *MongoDB
	once     sync.Once
)

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func Connect() (*MongoDB, error) {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}

		// Construir URI de MongoDB desde variables de entorno
		uri := os.Getenv("MONGO_URI")
		dbName := os.Getenv("MONGO_DATABASE")

		// Configurar opciones del cliente
		clientOptions := options.Client().ApplyURI(uri)

		// Configuraciones de conexión (equivalente a MySQL)
		clientOptions.SetMaxPoolSize(25)                  // max_open_conns
		clientOptions.SetMinPoolSize(5)                   // min_idle_conns
		clientOptions.SetMaxConnIdleTime(1 * time.Minute) // conn_max_lifetime
		clientOptions.SetConnectTimeout(10 * time.Second) // timeout de conexión
		clientOptions.SetSocketTimeout(30 * time.Second)  // timeout de socket

		// Crear contexto con timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Conectar a MongoDB
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatalf("Error connecting to MongoDB: %v", err)
		}

		// Verificar la conexión con ping (equivalente a db.Ping())
		if err := client.Ping(ctx, nil); err != nil {
			log.Fatalf("Error pinging MongoDB: %v", err)
		}

		// Obtener referencia a la base de datos
		database := client.Database(dbName)

		// Crear la colección "stories" si no existe
		collections, err := database.ListCollectionNames(ctx, bson.M{"name": "stories"})
		if err != nil {
			log.Fatalf("Error listing collections: %v", err)
		}
		if len(collections) == 0 {
			if err := database.CreateCollection(ctx, "stories"); err != nil {
				log.Fatalf("Error creating collection 'stories': %v", err)
			}
			log.Println("Collection 'stories' created successfully")
		}

		// Crear la colección "questions" si no existe
		collections, err = database.ListCollectionNames(ctx, bson.M{"name": "questions"})
		if err != nil {
			log.Fatalf("Error listing collections: %v", err)
		}
		if len(collections) == 0 {
			if err := database.CreateCollection(ctx, "questions"); err != nil {
				log.Fatalf("Error creating collection 'questions': %v", err)
			}
			log.Println("Collection 'questions' created successfully")
		}

		instance = &MongoDB{
			Client:   client,
			Database: database,
		}

		log.Println("Connected to MongoDB successfully")
	})

	return instance, nil
}
