package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service interface {
	Health() map[string]string
	Close() error
	GetDbORM() *gorm.DB
}

type service struct {
	db     *sql.DB
	gormDB *gorm.DB
}

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
	schema   = os.Getenv("DB_SCHEMA")

	dbInstance     *service
	gormInstance   *gorm.DB
	once           sync.Once
	connectionLock sync.Mutex
)

// New initializes and returns a singleton instance of the database service.
func New() Service {
	if dbInstance != nil {
		return dbInstance
	}

	connectionLock.Lock()
	defer connectionLock.Unlock()

	// Re-check inside the lock to avoid race conditions.
	if dbInstance == nil {
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
		db, err := sql.Open("pgx", connStr)
		if err != nil {
			log.Fatal(err)
		}

		dbInstance = &service{
			db: db,
		}

		once.Do(func() {
			var err error
			gormInstance, err = gorm.Open(postgres.New(postgres.Config{
				Conn: dbInstance.db,
			}), &gorm.Config{})
			if err != nil {
				log.Fatal(err)
			}

			// Automigrate database schema (if needed)
			err = gormInstance.AutoMigrate(&User{})
			if err != nil {
				log.Fatal(err)
			}

			dbInstance.gormDB = gormInstance
		})
	}

	return dbInstance
}

// Health checks the database health and returns stats.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err)
		return stats
	}

	stats["status"] = "up"
	stats["message"] = "It's healthy"

	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	return stats
}

// Close disconnects the database connection.
func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}

// GetDbORM returns the singleton instance of the gorm.DB object.
func (s *service) GetDbORM() *gorm.DB {
	return gormInstance
}
