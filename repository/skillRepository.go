package repository

import (
	"log"
	"mini-project/entity"

	"github.com/jinzhu/gorm"
)

type SkillRepository interface {
	CreateSkill(skill entity.Jobskill, jobSkillApplicant entity.Jobskillapplicant, applicantID int) (entity.Jobskill, error)
	FindUserByID(userID int) (entity.User, error)
	FindApplicantByID(userID int) (entity.Applicant, error)
}

type skillConnection struct {
	connection *gorm.DB
}

func NewSkillRepository(db *gorm.DB) SkillRepository {
	return &skillConnection{
		connection: db,
	}
}

func (db *skillConnection) CreateSkill(skill entity.Jobskill, jobsSkillApplicant entity.Jobskillapplicant, applicantID int) (entity.Jobskill, error) {
	err := db.connection.Save(&skill).Error
	if err != nil {
		return skill, err
	}
	
	jobsSkillApplicant.SkillID = skill.ID
	jobsSkillApplicant.ApplicantID = uint64(applicantID)
	err = db.connection.Save(&jobsSkillApplicant).Error
	if err != nil {
		log.Println(err)
		panic("Failed to save jobskillapplicant table")
	}

	return skill, nil
}

func (db *skillConnection) FindUserByID(userID int) (entity.User, error) {
	var user entity.User

	err := db.connection.Where("id = ?", userID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (db *skillConnection) FindApplicantByID(userID int) (entity.Applicant, error) {
	var applicant entity.Applicant

	err := db.connection.Where("user_id = ?", userID).Find(&applicant).Error
	if err != nil {
		return applicant, err
	}

	return applicant, nil
}