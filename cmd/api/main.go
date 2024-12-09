package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/romulosm/go-modular-backend-template/internal/infrastructure/database"
	"github.com/romulosm/go-modular-backend-template/internal/infrastructure/messaging"
	"github.com/romulosm/go-modular-backend-template/internal/user"
	"github.com/romulosm/go-modular-backend-template/pkg/logger"
)

func main() {
	// Carregar variáveis de ambiente
	if err := loadEnv(); err != nil {
		log.Fatalf("Erro ao carregar variáveis de ambiente: %v", err)
	}

	// Configurar o nível de log
	logger.SetLogLevel(os.Getenv("LOG_LEVEL"))

	// Criar um contexto base
	ctx := context.Background()

	// Inicializar conexões
	pgDB, err := database.NewPostgresConnection()
	if err != nil {
		logger.Log.Fatalf("Falha ao conectar ao PostgreSQL: %v", err)
	}
	defer pgDB.Close()

	mongoDB, err := database.NewMongoDBConnection()
	if err != nil {
		logger.Log.Fatalf("Falha ao conectar ao MongoDB: %v", err)
	}
	defer func() {
		if err := mongoDB.Disconnect(ctx); err != nil {
			logger.Log.Errorf("Erro ao desconectar do MongoDB: %v", err)
		}
	}()

	rabbitMQ, err := messaging.NewRabbitMQConnection()
	if err != nil {
		logger.Log.Fatalf("Falha ao conectar ao RabbitMQ: %v", err)
	}
	defer rabbitMQ.Close()

	// Configurar Gin
	r := gin.Default()

	// Inicializar e configurar módulo de usuário
	user.InitModule(r, pgDB)

	// Iniciar servidor
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	logger.Log.Infof("Servidor iniciado na porta %s", port)
	r.Run(":" + port)
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("erro ao carregar .env file: %w", err)
	}
	return nil
}
