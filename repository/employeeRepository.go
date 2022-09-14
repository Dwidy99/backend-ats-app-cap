package repository

import (
	"mini-project/entity"

	"github.com/jinzhu/gorm"
)

type EmployeeRepository interface {
	InsertEmployee(user entity.User, employee entity.Employee) entity.User
	SaveEmployee(employee entity.Employee) entity.Employee
	FindEmployeeByID(employeeUserID uint64) entity.Employee
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

func (db *employeeConnection) FindEmployeeByID(UserID uint64) entity.Employee {
	var employee entity.Employee

	err := db.connection.Where("user_id = ?", UserID).Find(&employee).Error
	if err != nil {
		return employee
	}

	return employee
}

func (db *employeeConnection) SaveEmployee(employee entity.Employee) entity.Employee {
	db.connection.Save(&employee)
	return employee
}

