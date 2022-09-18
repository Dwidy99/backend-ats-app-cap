package repository

import (
	"mini-project/entity"

	"github.com/jinzhu/gorm"
)

type ExperienceRepository interface {
	FindApplicantByID(userID int) (entity.Applicant, error)
	FindUserByID(userID int) (entity.User, error)
	FindExperienceByIdApplicant(idApplicant int) (entity.Jobexperience, error)
	FindExperienceByID(inputID int) (entity.Jobexperience, error)
	GetAllExperienceByID(applicantID int) ([]entity.Jobexperience, error)
	GetAllExperience() ([]entity.Jobexperience, error)
	InsertExperience(experience entity.Jobexperience) (entity.Jobexperience, error)
	Update(experience entity.Jobexperience) (entity.Jobexperience, error)
	DeleteExperience(inputID int) (entity.Jobexperience, error)
}

type experienceConnection struct {
	connection *gorm.DB
}

func NewExperienceRepository(db *gorm.DB) ExperienceRepository {
	return &experienceConnection{
		connection: db,
	}
}

func (db *experienceConnection) FindExperienceByIdApplicant(idApplicant int) (entity.Jobexperience, error) {
	var experience entity.Jobexperience

	err := db.connection.Where("applicant_id = ?", idApplicant).Find(&experience).Error
	if err != nil {
		return experience, err
	}

	return experience, nil
}

func (db *experienceConnection) FindExperienceByID(inputID int) (entity.Jobexperience, error) {
	var experience entity.Jobexperience

	err := db.connection.Where("id = ?", inputID).Find(&experience).Error
	if err != nil {
		return experience, err
	}

	return experience, nil
}

func (db *experienceConnection) FindApplicantByID(userID int) (entity.Applicant, error) {
	var applicant entity.Applicant

	err := db.connection.Where("user_id = ?", userID).Find(&applicant).Error
	if err != nil {
		return applicant, err
	}

	return applicant, nil
}

func (db *experienceConnection) FindUserByID(userID int) (entity.User, error) {
	var user entity.User

	err := db.connection.Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (db *experienceConnection) GetAllExperienceByID(applicantID int) ([]entity.Jobexperience, error) {
	var experience []entity.Jobexperience

	err := db.connection.Raw("SELECT id, applicant_id, company_name, role, description, date_start, date_end, status FROM jobexperiences WHERE applicant_id = ?", applicantID).Scan(&experience).Error
	if err != nil {
		return experience, err
	}

	return experience, nil
}

func (db *experienceConnection) GetAllExperience() ([]entity.Jobexperience, error) {
	var experience []entity.Jobexperience

	err := db.connection.Raw("SELECT id, applicant_id, company_name, role, description, date_start, date_end, status FROM jobexperiences").Scan(&experience).Error
	if err != nil {
		return experience, err
	}

	return experience, nil
}

func (db *experienceConnection) InsertExperience(experience entity.Jobexperience) (entity.Jobexperience, error) {
	err := db.connection.Save(&experience).Error
	if err != nil {
		return experience, err
	}
	return experience, nil
}

func (db *experienceConnection) Update(experience entity.Jobexperience) (entity.Jobexperience, error) {
	err := db.connection.Save(&experience).Error
	if err != nil {
		return experience, err
	}

	return experience, nil
}

func (db *experienceConnection) DeleteExperience(inputID int) (entity.Jobexperience, error) {
	var experience entity.Jobexperience
	err := db.connection.Where("id = ?", inputID).Delete(&experience).Error
	if err != nil {
		return experience, err
	}

	return experience, nil
}