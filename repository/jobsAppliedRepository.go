package repository

import (
	"github.com/PutraFajarF/backend-ats-app-cap/entity"

	"github.com/jinzhu/gorm"
)

type JobsAppliedRepository interface {
	JobApplicationsByApplicantID(applicantID int) ([]entity.Jobapplication, error)
	JobApplyByApplicantID(applicantID int) (entity.Jobapplication, error)
	FindApplicantByID(userID int) (entity.Applicant, error)
	GetJobByID(jobID int) (entity.Jobs, error)
	GetJobRequirementByID(jobID int) (entity.Jobskillrequirement, error)
	GetSkillByID(skillID int) (entity.Jobskill, error)
}

type jobsAppliedConnection struct {
	connection *gorm.DB
}

func NewJobAppliedConnection(db *gorm.DB) JobsAppliedRepository {
	return &jobsAppliedConnection{
		connection: db,
	}
}

func (db *jobsAppliedConnection) GetJobByID(inputID int) (entity.Jobs, error) {
	var job entity.Jobs

	err := db.connection.Where("id = ?", inputID).Find(&job).Error
	if err != nil {
		return job, err
	}

	return job, nil
}

func (db *jobsAppliedConnection) GetJobRequirementByID(jobID int) (entity.Jobskillrequirement, error) {
	var jobRequirement entity.Jobskillrequirement

	err := db.connection.Where("id_job = ?", jobID).Find(&jobRequirement).Error
	if err != nil {
		return jobRequirement, err
	}

	return jobRequirement, nil
}

func (db *jobsAppliedConnection) GetSkillByID(skillID int) (entity.Jobskill, error) {
	var jobSkill entity.Jobskill

	err := db.connection.Where("id = ?", skillID).Find(&jobSkill).Error
	if err != nil {
		return jobSkill, err
	}

	return jobSkill, nil
}

func (db *jobsAppliedConnection) JobApplicationsByApplicantID(applicantID int) ([]entity.Jobapplication, error) {
	var jobapplication []entity.Jobapplication

	err := db.connection.Where("applicant_id = ?", applicantID).Find(&jobapplication).Error
	if err != nil {
		return jobapplication, err
	}

	return jobapplication, nil
}

func (db *jobsAppliedConnection) JobApplyByApplicantID(applicantID int) (entity.Jobapplication, error) {
	var jobapplication entity.Jobapplication

	err := db.connection.Where("applicant_id = ?", applicantID).Find(&jobapplication).Error
	if err != nil {
		return jobapplication, err
	}

	return jobapplication, nil
}

func (db *jobsAppliedConnection) FindApplicantByID(userID int) (entity.Applicant, error) {
	var applicant entity.Applicant

	err := db.connection.Where("user_id = ?", userID).Find(&applicant).Error
	if err != nil {
		return applicant, err
	}

	return applicant, nil
}
