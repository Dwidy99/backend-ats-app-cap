package service

import (
	"errors"
	"fmt"
	"github.com/PutraFajarF/backend-ats-app-cap/dto"
	"github.com/PutraFajarF/backend-ats-app-cap/entity"
	"github.com/PutraFajarF/backend-ats-app-cap/repository"

	"github.com/mashingan/smapping"
)

type EmployeeService interface {
	IsAllowedToEdit(UserID string, employeeUserID uint64) (bool, error)
	UpdateEmployee(employee dto.EmployeeUpdateDTO, id int) (entity.Employee, error)
	GetEmployeeById(userId uint64) (entity.Employee, error)
}

type employeeService struct {
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeService(employeeRepo repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		employeeRepository: employeeRepo,
	}
}

func (s *employeeService) IsAllowedToEdit(userID string, employeeUserID uint64) (bool, error) {
	employee, err := s.employeeRepository.FindEmployeeByID(employeeUserID)
	if err != nil {
		return false, err
	}
	id := fmt.Sprintf("%v", employee.UserID)

	if userID != id {
		return false, errors.New("not allowed to edit")
	}
	return true, nil
}

func (s *employeeService) UpdateEmployee(e dto.EmployeeUpdateDTO, id int) (entity.Employee, error) {
	employee, err := s.employeeRepository.FindEmployeeByID(uint64(id))

	employee.Name = e.Name
	employee.Contact = e.Contact

	err = smapping.FillStruct(&employee, smapping.MapFields(&e))
	if err != nil {
		return employee, err
	}

	res, err := s.employeeRepository.SaveEmployee(employee)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *employeeService) GetEmployeeById(userId uint64) (entity.Employee, error) {
	res, err := s.employeeRepository.FindEmployeeByID(userId)
	err = smapping.FillStruct(userId, smapping.MapFields(&userId))
	if err != nil {
		return res, err
	}

	return res, nil
}
