package service

import (
	"errors"
	"fmt"

	"github.com/PutraFajarF/backend-ats-app-cap/dto"
	"github.com/PutraFajarF/backend-ats-app-cap/entity"
	"github.com/PutraFajarF/backend-ats-app-cap/repository"
)

type JobApplicationEmployeeService interface {
	UpdateProgress(inputData dto.UpdateJobApplicationEmployeeDTO, ID int) (entity.Jobapplication, error)
	GetUserByID(userID int) (entity.User, error)
	GetApplicantByID(userID int) (entity.Applicant, error)
}

type jobApplicationEmployeeService struct {
	jobApplicationEmployeeRepository repository.JobApplicationEmployeeRepository
}

func NewJobApplicationEmployeeService(jobApplicationEmployeeRepository repository.JobApplicationEmployeeRepository) JobApplicationEmployeeService {
	return &jobApplicationEmployeeService{
		jobApplicationEmployeeRepository: jobApplicationEmployeeRepository,
	}
}

func (s *jobApplicationEmployeeService) UpdateProgress(inputData dto.UpdateJobApplicationEmployeeDTO, ID int) (entity.Jobapplication, error) {
	fmt.Println(ID)
	jobApplication, err := s.jobApplicationEmployeeRepository.FindJobApplicantByID(ID)
	if err != nil {
		return jobApplication, err
	}

	jobApplication.Status = inputData.Status

	apply, err := s.jobApplicationEmployeeRepository.UpdateProgress(jobApplication)
	if err != nil {
		return apply, err
	}

	return apply, nil
}


func (s *jobApplicationEmployeeService) GetUserByID(userID int) (entity.User, error) {
	user, err := s.jobApplicationEmployeeRepository.FindUserByID(userID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user logged in")
	}

	return user, nil
}

func (s *jobApplicationEmployeeService) GetApplicantByID(userID int) (entity.Applicant, error) {
	applicant, err := s.jobApplicationEmployeeRepository.FindApplicantByID(userID)
	if err != nil {
		return applicant, err
	}

	if applicant.UserID == 0 {
		return applicant, errors.New("no user logged in")
	}

	return applicant, nil
}
