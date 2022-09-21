package repository

import (
	"mini-project/entity"

	"github.com/jinzhu/gorm"
)

type JobsRepository interface {
	FindUserByID(userID int) (entity.User, error)
	InsertJobs(j entity.Jobs) (entity.Jobs, error)
}

type jobsConnection struct {
	connection *gorm.DB
}

func NewJobsRepository(db *gorm.DB) JobsRepository {
	return &jobsConnection{
		connection: db,
	}
}

func (db *jobsConnection) FindUserByID(userID int) (entity.User, error) {
	var user entity.User

	err := db.connection.Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (db *jobsConnection) InsertJobs(j entity.Jobs) (entity.Jobs, error) {
	err := db.connection.Save(&j).Error
	if err != nil {
		return j, err
	}
	return j, nil
}
