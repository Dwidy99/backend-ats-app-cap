package service

import (
	"github.com/PutraFajarF/backend-ats-app-cap/dto"
	"github.com/PutraFajarF/backend-ats-app-cap/entity"
	"github.com/PutraFajarF/backend-ats-app-cap/repository"
	"log"

	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) bool
	CreateApplicant(user dto.RegisterApplicantDTO) entity.User
	CreateEmployee(user dto.RegisterEmployeeDTO) (entity.User, error)
	GetUserByID(userID int) (entity.User, error)
}

type authService struct {
	userRepository      repository.UserRepository
	applicantRepository repository.ApplicantRepository
	employeeRepository  repository.EmployeeRepository
}

func NewAuthService(userRep repository.UserRepository, applicantRep repository.ApplicantRepository, employeeRep repository.EmployeeRepository) AuthService {
	return &authService{
		userRepository:      userRep,
		applicantRepository: applicantRep,
		employeeRepository:  employeeRep,
	}
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func (service *authService) VerifyCredential(email string, password string) interface{} {
	res := service.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (service *authService) CreateApplicant(user dto.RegisterApplicantDTO) entity.User {
	userToCreate := entity.User{}
	applicantToCreate := entity.Applicant{}
	userToCreate.Role = "user"

	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.applicantRepository.InsertApplicant(userToCreate, applicantToCreate)
	return res
}

func (service *authService) CreateEmployee(user dto.RegisterEmployeeDTO) (entity.User, error) {
	userToCreate := entity.User{}
	employeeToCreate := entity.Employee{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res, err := service.employeeRepository.InsertEmployee(userToCreate, employeeToCreate)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *authService) GetUserByID(userID int) (entity.User, error) {
	user, err := s.employeeRepository.FindUserByID(userID)
	if err != nil {
		return user, err
	}
	return user, nil
}
