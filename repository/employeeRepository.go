package repository

import (
	"mini-project/entity"

	"github.com/jinzhu/gorm"
)

type EmployeeRepository interface {
	InsertEmployee(user entity.User, employee entity.Employee) entity.User
}

type employeeConnection struct {
	connection *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeConnection{
		connection: db,
	}
}

func (db *employeeConnection) InsertEmployee(user entity.User, employee entity.Employee) entity.User {
	user.Password = HashAndSalt([]byte(user.Password))
	db.connection.Save(&user)
	employee.UserID = int(user.ID)
	db.connection.Save(&employee)
	return user
}