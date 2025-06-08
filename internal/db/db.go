package db

import (
	"step1_simple_api/internal/types"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/rs/zerolog/log"
)

type DB struct {
	db *gorm.DB
}

func InitDB() (*DB, error) {
	dsn := "host=localhost user=postgres password=password dbname=postgres port=5432 sslmode=disable"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal().Err(err).Msg("Could not connect to database")
		return nil, err
	}

	gormDB = gormDB.Debug()

	log.Info().Msg("Database connection established successfully")
	return &DB{db: gormDB}, nil
}

func (db *DB) GetTasks() ([]types.Task, error) {
	tasks := []types.Task{}

	if err := db.db.Find(&tasks).Error; err != nil {
		log.Fatal().Err(err).Msg("Failed to fetch tasks")
		return tasks, err
	}

	log.Info().Int("count", len(tasks)).Msg("Tasks fetched successfully")
	return tasks, nil
}

func (db *DB) GetTask(objectID string) (types.Task, error) {
	task := types.Task{}

	if err := db.db.First(&task, "uuid = ?", objectID).Error; err != nil {
		log.Warn().Err(err).Str("uuid", objectID).Msg("Task not found")
		return task, err
	}

	log.Info().Str("count", objectID).Msg("Tasks fetched successfully")
	return task, nil
}

func (db *DB) CreateTask(task types.Task) (types.Task, error) {
	if err := db.db.Create(&task).Error; err != nil {
		log.Warn().Err(err).Str("uuid", task.UUID).Msg("Failed to create task")
		return task, err
	}

	log.Info().Str("count", task.UUID).Msg("Task created successfully")
	return task, nil
}

func (db *DB) UpdateTask(task types.Task) error {
	if err := db.db.Updates(&task).Error; err != nil {
		log.Warn().Err(err).Str("uuid", task.UUID).Msg("Failed to update task")
		return err
	}

	log.Info().Str("count", task.UUID).Msg("Task updated successfully")
	return nil
}

func (db *DB) DestroyTask(objectID string) error {
	task := types.Task{UUID: objectID}

	if err := db.db.Delete(&task).Error; err != nil {
		log.Warn().Err(err).Str("uuid", task.UUID).Msg("Failed to delete task")
		return err
	}

	log.Info().Str("count", task.UUID).Msg("Task deleted successfully")
	return nil
}

func (db *DB) GetUsers() ([]types.User, error) {
	users := []types.User{}

	if err := db.db.Find(&users).Error; err != nil {
		log.Fatal().Err(err).Msg("Failed to fetch users")
		return users, err
	}

	log.Info().Int("count", len(users)).Msg("Users fetched successfully")
	return users, nil
}

func (db *DB) GetUser(objectID int) (types.User, error) {
	user := types.User{}

	if err := db.db.First(&user, "id = ?", objectID).Error; err != nil {
		log.Warn().Err(err).Int("id", objectID).Msg("User not found")
		return user, err
	}

	log.Info().Int("count", objectID).Msg("User fetched successfully")
	return user, nil
}

func (db *DB) CreateUser(user types.User) (types.User, error) {
	if err := db.db.Create(&user).Error; err != nil {
		log.Warn().Err(err).Int("id", user.ID).Msg("Failed to create user")
		return user, err
	}

	log.Info().Int("count", user.ID).Msg("User created successfully")
	return user, nil
}

func (db *DB) UpdateUser(user types.User) error {
	if err := db.db.Updates(&user).Error; err != nil {
		log.Warn().Err(err).Int("id", user.ID).Msg("Failed to update user")
		return err
	}

	log.Info().Int("count", user.ID).Msg("User updated successfully")
	return nil
}

func (db *DB) DestroyUser(objectID int) error {
	user := types.User{ID: objectID}

	if err := db.db.Delete(&user).Error; err != nil {
		log.Warn().Err(err).Int("id", user.ID).Msg("Failed to delete user")
		return err
	}

	log.Info().Int("count", user.ID).Msg("User deleted successfully")
	return nil
}
