package repository

import (
	"fmt"
	"log"

	"github.com/PutraFajarF/backend-ats-app-cap/entity"

	"github.com/jinzhu/gorm"
)

type SkillRepository interface {
	CreateSkill(skill entity.Jobskill, jobSkillApplicant entity.Jobskillapplicant, applicantID int) (entity.Jobskill, error)
	FindUserByID(userID int) (entity.User, error)
	FindApplicantByID(userID int) (entity.Applicant, error)
	FindEmployeeByID(userID int) (entity.Employee, error)
	GetSkillByID(inputID int) (entity.Jobskill, error)
	GetJobSkillApplicantBySkillID(inputID int) ([]entity.Jobskillapplicant, error)
	Update(jobSkill entity.Jobskill) (entity.Jobskill, error)
	Delete(inputID int) (entity.Jobskill, error)
	GetSkills(applicantID int) ([]entity.Jobskillapplicant, error)
	GetJobSkillAppByApplicantID(applicantID int) ([]entity.Jobskillapplicant, error)
}

type skillConnection struct {
	connection *gorm.DB
}

func NewSkillRepository(db *gorm.DB) SkillRepository {
	return &skillConnection{
		connection: db,
	}
}

func (db *skillConnection) GetSkills(applicantID int) ([]entity.Jobskillapplicant, error) {
	var jobSkillApplicant []entity.Jobskillapplicant

	err := db.connection.Where("applicant_id = ?", applicantID).Find(&jobSkillApplicant).Error
	if err != nil {
		return jobSkillApplicant, err
	}
	fmt.Println("entity", jobSkillApplicant)

	return jobSkillApplicant, nil
}

func (db *skillConnection) GetJobSkillAppByApplicantID(applicantID int) ([]entity.Jobskillapplicant, error) {
	var jobSkillApplicant []entity.Jobskillapplicant

	err := db.connection.Where("applicant_id = ?", applicantID).Find(&jobSkillApplicant).Error
	if err != nil {
		return jobSkillApplicant, err
	}

	return jobSkillApplicant, nil
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

func (db *skillConnection) FindEmployeeByID(userID int) (entity.Employee, error) {
	var employee entity.Employee

	err := db.connection.Where("user_id = ?", userID).Find(&employee).Error
	if err != nil {
		return employee, err
	}

	return employee, nil
}

func (db *skillConnection) GetSkillByID(inputID int) (entity.Jobskill, error) {
	var skill entity.Jobskill

	err := db.connection.Where("id = ?", inputID).Find(&skill).Error
	if err != nil {
		return skill, err
	}

	return skill, nil
}

func (db *skillConnection) GetJobSkillApplicantBySkillID(inputID int) ([]entity.Jobskillapplicant, error) {
	var jobSkillApplicant []entity.Jobskillapplicant

	err := db.connection.Where("skill_id = ?", inputID).Find(&jobSkillApplicant).Error
	if err != nil {
		return jobSkillApplicant, err
	}

	return jobSkillApplicant, nil
}

func (db *skillConnection) Update(jobSkill entity.Jobskill) (entity.Jobskill, error) {
	err := db.connection.Save(&jobSkill).Error
	if err != nil {
		return jobSkill, err
	}

	return jobSkill, nil
}

func (db *skillConnection) Delete(inputID int) (entity.Jobskill, error) {
	var jobSkill entity.Jobskill
	var jobSkillApplicant entity.Jobskillapplicant

	err := db.connection.Where("id = ?", inputID).Delete(&jobSkill).Error
	err = db.connection.Where("skill_id = ?", inputID).Delete(&jobSkillApplicant).Error
	if err != nil {
		return jobSkill, err
	}

	return jobSkill, nil
}
