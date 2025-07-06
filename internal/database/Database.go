package database

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
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
		uri := "mongodb://" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT")
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

		instance = &MongoDB{
			Client:   client,
			Database: database,
		}

		log.Println("Connected to MongoDB successfully")
	})

	return instance, nil
}
