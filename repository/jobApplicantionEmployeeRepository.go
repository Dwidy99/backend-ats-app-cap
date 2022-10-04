package repository

import (
	"github.com/PutraFajarF/backend-ats-app-cap/entity"

	"github.com/jinzhu/gorm"
)

type JobApplicationEmployeeRepository interface {
	UpdateProgress(apply entity.Jobapplication) (entity.Jobapplication, error)
	FindUserByID(userID int) (entity.User, error)
	FindApplicantByID(userID int) (entity.Applicant, error)
	FindJobApplicantByID(ID int) (entity.Jobapplication, error)
}

type jobApplicationEmployeeConnection struct {
	connection *gorm.DB
}

func NewjobApplicationEmployee(db *gorm.DB) JobApplicationEmployeeRepository {
	return &jobApplicationEmployeeConnection{
		connection: db,
	}
}

func (db *jobApplicationEmployeeConnection) UpdateProgress(apply entity.Jobapplication) (entity.Jobapplication, error) {
	err := db.connection.Save(&apply).Error
	if err != nil {
		return apply, err
	}

	return apply, nil
}

func (db *jobApplicationEmployeeConnection) FindUserByID(userID int) (entity.User, error) {
	var user entity.User

	err := db.connection.Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (db *jobApplicationEmployeeConnection) FindApplicantByID(userID int) (entity.Applicant, error) {
	var applicant entity.Applicant

	err := db.connection.Where("user_id = ?", userID).Find(&applicant).Error
	if err != nil {
		return applicant, err
	}

	return applicant, nil
}

func (db *jobApplicationEmployeeConnection) FindJobApplicantByID(ID int) (entity.Jobapplication, error) {
	var jobapplication entity.Jobapplication

	err := db.connection.Where("id = ?", ID).Find(&jobapplication).Error
	if err != nil {
		return jobapplication, err
	}

	return jobapplication, nil
}

