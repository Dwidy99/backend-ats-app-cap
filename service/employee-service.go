package service

import (
	"fmt"
	"log"
	"mini-project/dto"
	"mini-project/entity"
	"mini-project/repository"

	"github.com/mashingan/smapping"
)

type EmployeeService interface {
	IsAllowedToEdit(UserID string, employeeUserID uint64) bool
	UpdateEmployee(employee dto.EmployeeUpdateDTO, id int) entity.Employee
}

type employeeService struct{
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeService(employeeRepo repository.EmployeeRepository) EmployeeService {
	return &employeeService{
		employeeRepository: employeeRepo,
	}
}

func (s *employeeService) IsAllowedToEdit(userID string, employeeUserID uint64) bool {
	applicant := s.employeeRepository.FindApplicantByID(employeeUserID)
	id := fmt.Sprintf("%v", applicant.UserID)

	return userID == id
}

func (s *employeeService) UpdateEmployee(e dto.EmployeeUpdateDTO, id int) entity.Employee {
	employee := s.employeeRepository.FindApplicantByID(uint64(id))

	employee.Name = e.Name
	employee.Contact = e.Contact

	err := smapping.FillStruct(&employee, smapping.MapFields(&e))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := s.employeeRepository.SaveEmployee(employee)
	return res
}