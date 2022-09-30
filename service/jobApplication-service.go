package service

import (
	"errors"
	"github.com/PutraFajarF/backend-ats-app-cap/dto"
	"github.com/PutraFajarF/backend-ats-app-cap/entity"
	"github.com/PutraFajarF/backend-ats-app-cap/repository"

	"github.com/mashingan/smapping"
)

type JobApplicationService interface {
	CreateJobApplicant(inputData dto.CreateJobApplication, applicantID int) (entity.Jobapplication, error)
	GetUserByID(userID int) (entity.User, error)
	GetApplicantByID(userID int) (entity.Applicant, error)
}

type jobApplicationService struct {
	jobApplicationRepository repository.JobApplicationRepository
}

func NewJobApplicationService(jobApplicationRepository repository.JobApplicationRepository) JobApplicationService {
	return &jobApplicationService{
		jobApplicationRepository: jobApplicationRepository,
	}
}

func (s *jobApplicationService) CreateJobApplicant(inputData dto.CreateJobApplication, applicantID int) (entity.Jobapplication, error) {
	createJobApplication := entity.Jobapplication{}
	createJobApplication.ApplicantID = uint64(applicantID)
	createJobApplication.JobID = inputData.JobID
	createJobApplication.Status = inputData.Status

	err := smapping.FillStruct(&createJobApplication, smapping.MapFields(&inputData))
	if err != nil {
		return createJobApplication, err
	}

	apply, err := s.jobApplicationRepository.CreateApply(createJobApplication)
	if err != nil {
		return apply, err
	}

	return apply, nil
}

func (s *jobApplicationService) GetUserByID(userID int) (entity.User, error) {
	user, err := s.jobApplicationRepository.FindUserByID(userID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user logged in")
	}

	return user, nil
}

func (s *jobApplicationService) GetApplicantByID(userID int) (entity.Applicant, error) {
	applicant, err := s.jobApplicationRepository.FindApplicantByID(userID)
	if err != nil {
		return applicant, err
	}

	if applicant.UserID == 0 {
		return applicant, errors.New("no user logged in")
	}

	return applicant, nil
}
