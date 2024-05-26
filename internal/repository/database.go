package repository

import (
	"context"
	"database/sql"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	Connect() error
}

type MongoDBConfig struct {
	URI      string
	DBName   string
	Username string
	Password string
}

type PostgreSQLConfig struct {
	URI      string
	DBName   string
	Username string
	Password string
}

type MongoDBConnection struct {
	config MongoDBConfig
	client *mongo.Client
}

type PostgreSQLConnection struct {
	config PostgreSQLConfig
	db     *sql.DB
}

func (m *MongoDBConnection) Connect() error {
	clientOptions := options.Client().ApplyURI(m.config.URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	m.client = client
	fmt.Println("Connected to MongoDB successfully!")
	return nil
}

func (p *PostgreSQLConnection) Connect() error {
	db, err := sql.Open("postgres", p.config.URI)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	p.db = db
	fmt.Println("Connected to PostgreSQL successfully!")
	return nil
}
