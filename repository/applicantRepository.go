package repository

import (
	"github.com/PutraFajarF/backend-ats-app-cap/entity"

	"github.com/jinzhu/gorm"
)

type ApplicantRepository interface {
	InsertApplicant(user entity.User, applicant entity.Applicant) entity.User
	FindApplicantByUserID(applicantUserID uint64) (entity.Applicant, error)
	SaveApplicant(applicant entity.Applicant) (entity.Applicant, error)
	GetExperienceByApplicantID(applicantID uint64) ([]entity.Jobexperience, error)
	GetJobSkillApplicantByApplicantID(applicantID uint64) ([]entity.Jobskillapplicant, error)
	GetJobSkillByJobSkillApplicantID(skillID uint64) ([]entity.Jobskill, error)
}

type applicantConnection struct {
	connection *gorm.DB
}

func NewApplicantRepository(db *gorm.DB) ApplicantRepository {
	return &applicantConnection{
		connection: db,
	}
}

func (db *applicantConnection) InsertApplicant(user entity.User, applicant entity.Applicant) entity.User {
	user.Password = HashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	applicant.UserID = user.ID
	db.connection.Save(&applicant)
	return user
}

func (db *applicantConnection) FindApplicantByUserID(UserID uint64) (entity.Applicant, error) {
	var applicant entity.Applicant

	err := db.connection.Where("user_id = ?", UserID).Find(&applicant).Error
	if err != nil {
		return applicant, err
	}

	return applicant, nil
}

func (db *applicantConnection) GetExperienceByApplicantID(applicantID uint64) ([]entity.Jobexperience, error) {
	var experience []entity.Jobexperience

	err := db.connection.Where("applicant_id = ?", applicantID).Find(&experience).Error
	if err != nil {
		return experience, err
	}

	return experience, nil
}

func (db *applicantConnection) GetJobSkillApplicantByApplicantID(applicantID uint64) ([]entity.Jobskillapplicant, error) {
	var jobskillapplicant []entity.Jobskillapplicant

	err := db.connection.Where("applicant_id = ?", applicantID).Find(&jobskillapplicant).Error
	if err != nil {
		return jobskillapplicant, err
	}

	return jobskillapplicant, nil
}

func (db *applicantConnection) GetJobSkillByJobSkillApplicantID(skillID uint64) ([]entity.Jobskill, error) {
	var jobskillapplicant []entity.Jobskill

	err := db.connection.Where("id = ?", skillID).Find(&jobskillapplicant).Error
	if err != nil {
		return jobskillapplicant, err
	}

	return jobskillapplicant, nil
}

func (db *applicantConnection) SaveApplicant(applicant entity.Applicant) (entity.Applicant, error) {
	err := db.connection.Save(&applicant).Error
	if err != nil {
		return applicant, err
	}
	return applicant, nil
}
