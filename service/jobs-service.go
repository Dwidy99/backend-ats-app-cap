package service

import (
	"errors"
	"log"
	"mini-project/dto"
	"mini-project/entity"
	"mini-project/repository"

	"github.com/mashingan/smapping"
)

type JobsService interface {
	GetUserByID(userID int) (entity.User, error)
	CreateJobs(jobs dto.CreateJobsDTO, userID int) (entity.Jobs, error)
}

type jobsService struct {
	jobsRepository repository.JobsRepository
}

func NewJobsService(jobRepository repository.JobsRepository) JobsService {
	return &jobsService{
		jobsRepository: jobRepository,
	}
}

func (s *jobsService) GetUserByID(userID int) (entity.User, error) {
	user, err := s.jobsRepository.FindUserByID(userID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user logged in")
	}

	return user, nil
}

func (s *jobsService) CreateJobs(jobs dto.CreateJobsDTO, userID int) (entity.Jobs, error) {
	createJob := entity.Jobs{}
	createJob.CompanyID = jobs.CompanyID
	createJob.Title = jobs.Title
	createJob.Description = jobs.Description
	createJob.Location = jobs.Location
	createJob.Salary = jobs.Salary
	createJob.Type = jobs.Type
	createJob.LevelOfExperience = jobs.LevelOfExperience
	createJob.DateStart = jobs.DateStart
	createJob.DateEnd = jobs.DateEnd
	createJob.PostedBy = uint64(userID)

	err := smapping.FillStruct(&createJob, smapping.MapFields(&jobs))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res, err := s.jobsRepository.InsertJobs(createJob)
	if err != nil {
		return res, err
	}
	return res, nil
}
