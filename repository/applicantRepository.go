package repository

import (
	"mini-project/entity"

	"github.com/jinzhu/gorm"
)

type ApplicantRepository interface {
	InsertApplicant(user entity.User, applicant entity.Applicant) entity.User
	FindApplicantByID(applicantUserID uint64) entity.Applicant
	SaveApplicant(applicant entity.Applicant) entity.Applicant
}

type applicantConnection struct {
	connection *gorm.DB
}

func NewApplicantRepository(db *gorm.DB) ApplicantRepository {
	return &applicantConnection {
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

func (db *applicantConnection) FindApplicantByID(UserID uint64) entity.Applicant {
	var applicant entity.Applicant

	db.connection.Raw("SELECT user_id FROM applicants WHERE user_id = ?", UserID).Scan(&applicant)
	return applicant
}

func (db *applicantConnection) SaveApplicant(applicant entity.Applicant) entity.Applicant {
	db.connection.Save(&applicant)
	db.connection.Preload("User").Find(&applicant)
	return applicant
}