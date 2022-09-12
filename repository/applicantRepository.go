package repository

import (
	"fmt"
	"mini-project/entity"

	"github.com/jinzhu/gorm"
)

type ApplicantRepository interface {
	InsertApplicant(user entity.User, applicant entity.Applicant) entity.User
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
	fmt.Println("PASSWORD", user.Password)
	return user
}