package service

import (
	"errors"
	"mini-project/entity"
	"mini-project/repository"
)

type JobAppliedService interface {
	JobAppliedByApplicantID(applicantID int) ([]entity.Jobapplication, error)
	GetApplicantByID(userID int) (entity.Applicant, error)
}

type jobAppliedService struct {
	jobsAppliedRepository repository.JobsAppliedRepository
}

func NewJobAppliedService(jobsAppliedRepository repository.JobsAppliedRepository) JobAppliedService {
	return &jobAppliedService{
		jobsAppliedRepository: jobsAppliedRepository,
	}
}

func (s *jobAppliedService) JobAppliedByApplicantID(applicantID int) ([]entity.Jobapplication, error) {
	applies, err := s.jobsAppliedRepository.JobApplicationsByApplicantID(applicantID)
	if err != nil {
		return applies, err
	}

	return applies, nil
}

func (s *jobAppliedService) GetApplicantByID(userID int) (entity.Applicant, error) {
	applicant, err := s.jobsAppliedRepository.FindApplicantByID(userID)
	if err != nil {
		return applicant, err
	}
	
	if applicant.UserID == 0 {
		return applicant, errors.New("no user logged in")
	}

	return applicant, nil
}