package repository

import (
	"mini-project/entity"

	"github.com/jinzhu/gorm"
)

type JobApplicationRepository interface {
	CreateApply(apply entity.Jobapplication) (entity.Jobapplication, error)
	FindUserByID(userID int) (entity.User, error)
	FindApplicantByID(userID int) (entity.Applicant, error)
}

type jobApplicationConnection struct {
	connection *gorm.DB
}

func NewjobApplication(db *gorm.DB) JobApplicationRepository{
	return &jobApplicationConnection{
		connection: db,
	}
}

func (db *jobApplicationConnection) CreateApply(apply entity.Jobapplication) (entity.Jobapplication, error) {
	err := db.connection.Save(&apply).Error
	if err != nil {
		return apply, err
	}

	return apply, nil
}

func (db *jobApplicationConnection) FindUserByID(userID int) (entity.User, error) {
	var user entity.User

	err := db.connection.Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (db *jobApplicationConnection) FindApplicantByID(userID int) (entity.Applicant, error) {
	var applicant entity.Applicant

	err := db.connection.Where("user_id = ?", userID).Find(&applicant).Error
	if err != nil {
		return applicant, err
	}

	return applicant, nil
}