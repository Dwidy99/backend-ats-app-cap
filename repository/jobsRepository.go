package repository

import (
	"mini-project/entity"

	"github.com/jinzhu/gorm"
)

type JobsRepository interface {
	FindUserByID(userID int) (entity.User, error)
	FindEmployeeByID(userID int) (entity.Employee, error)
	FindJobsByID(inputID int) (entity.Jobs, error)
	GetAllJob() ([]entity.Jobs, error)
	InsertJobs(j entity.Jobs) (entity.Jobs, error)
	Update(j entity.Jobs) (entity.Jobs, error)
	DeleteJob(inputID int) (entity.Jobs, error)
	CheckID(inputID int) (bool, error)
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

func (db *jobsConnection) FindEmployeeByID(userID int) (entity.Employee, error) {
	var employee entity.Employee

	err := db.connection.Where("user_id = ?", userID).Find(&employee).Error
	if err != nil {
		return employee, err
	}
	return employee, nil
}

func (db *jobsConnection) FindJobsByID(inputID int) (entity.Jobs, error) {
	var jobs entity.Jobs

	err := db.connection.Where("id = ?", inputID).Find(&jobs).Error
	if err != nil {
		return jobs, err
	}

	return jobs, nil
}

func (db *jobsConnection) InsertJobs(j entity.Jobs) (entity.Jobs, error) {
	err := db.connection.Save(&j).Error
	if err != nil {
		return j, err
	}
	return j, nil
}

func (db *jobsConnection) GetAllJob() ([]entity.Jobs, error) {
	var jobs []entity.Jobs
	err := db.connection.Raw("SELECT id, company_id, title, description, location, salary, type, level_of_experience, date_start, date_end, created_at, posted_by FROM jobs").Scan(&jobs).Error
	if err != nil {
		return jobs, err
	}
	return jobs, nil
}

func (db *jobsConnection) Update(j entity.Jobs) (entity.Jobs, error) {
	err := db.connection.Save(&j).Error
	if err != nil {
		return j, err
	}

	return j, nil
}

func (db *jobsConnection) CheckID(inputID int) (bool, error) {
	var exists bool
	err := db.connection.
		Select("count(*) > 0").
		Where("id = ?", inputID).
		Find(&exists).
		Error
	if err != nil {
		return exists, err
	}

	return exists, nil
}

func (db *jobsConnection) DeleteJob(inputID int) (entity.Jobs, error) {
	var jobs entity.Jobs
	err := db.connection.Where("id = ?", inputID).Delete(&jobs).Error
	if err != nil {
		return jobs, err
	}

	return jobs, nil
}
