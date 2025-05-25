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

	gormDB = gormDB.Debug()

	logrus.Info("Database connection established successfully")
	return &DB{db: gormDB}, nil
}

func (db *DB) GetTasks() ([]types.Task, error) {
	tasks := []types.Task{}

	if err := db.db.Find(&tasks).Error; err != nil {
		logrus.WithError(err).Error("Failed to fetch tasks")
		return tasks, err
	}

	logrus.WithField("count", len(tasks)).Info("Tasks fetched successfully")
	return tasks, nil
}

func (db *DB) GetTask(objectID string) (types.Task, error) {
	task := types.Task{}

	if err := db.db.First(&task, "uuid = ?", objectID).Error; err != nil {
		logrus.WithError(err).WithField("uuid", objectID).Warn("Task not found")
		return task, err
	}

	logrus.WithField("uuid", objectID).Info("Task fetched successfully")
	return task, nil
}

func (db *DB) CreateTask(task types.Task) (types.Task, error) {
	if err := db.db.Create(&task).Error; err != nil {
		logrus.WithError(err).WithField("uuid", task.UUID).Error("Failed to create task")
		return task, err
	}

	logrus.WithField("uuid", task.UUID).Info("Task created successfully")
	return task, nil
}

func (db *DB) UpdateTask(task types.Task) error {
	if err := db.db.Updates(&task).Error; err != nil {
		logrus.WithError(err).WithField("uuid", task.UUID).Error("Failed to update task")
		return err
	}

	logrus.WithField("uuid", task.UUID).Info("Task updated successfully")
	return nil
}

func (db *DB) DestroyTask(objectID string) error {
	task := types.Task{UUID: objectID}

	if err := db.db.Delete(&task).Error; err != nil {
		logrus.WithError(err).WithField("uuid", objectID).Error("Failed to delete task")
		return err
	}

	logrus.WithField("uuid", objectID).Info("Task deleted successfully")
	return nil
}
