package db

import (
	"step1_simple_api/internal/types"
	"sync"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	mu *sync.Mutex
	db *gorm.DB
}

func InitDB() (*DB, error) {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Fatalf("Could not connect to database, %v", err)
		return nil, err
	}

	if err := gormDB.AutoMigrate(&types.Task{}); err != nil {
		logrus.Fatalf("Could not migrate, %v", err)
		return nil, err
	}

	return &DB{db: gormDB}, nil
}

func (db *DB) ObjectsTasks() ([]types.Task, error) {
	tasks := []types.Task{}

	if err := db.db.Find(&tasks).Error; err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (db *DB) ObjectTask(objectID string) (types.Task, error) {
	task := types.Task{}

	if err := db.db.First(&task, "uuid = ?", objectID).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (db *DB) CreateTask(task types.Task) (types.Task, error) {
	if err := db.db.Create(&task).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (db *DB) UpdateTask(task types.Task) error {
	if err := db.db.Updates(&task).Error; err != nil {
		return err
	}

	return nil
}

func (db *DB) DestroyTask(objectID string) error {
	task := types.Task{UUID: objectID}

	if err := db.db.Delete(&task).Error; err != nil {
		return err
	}

	return nil
}
