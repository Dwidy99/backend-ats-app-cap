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
	GetEmployeeByID(userID int) (entity.Employee, error)
	AllJobs() ([]entity.Jobs, error)
	GetJobByID(inputID int) (entity.Jobs, error)
	CreateJobs(jobs dto.CreateJobsDTO, userID int) (entity.Jobs, error)
	UpdateJob(inputData dto.CreateJobsDTO, inputID dto.JobDetailDTO, userID int) (entity.Jobs, error)
	DeletedJob(inputID int) (entity.Jobs, error)
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

func (s *jobsService) GetEmployeeByID(userID int) (entity.Employee, error) {
	employee, err := s.jobsRepository.FindEmployeeByID(userID)
	if err != nil {
		return employee, err
	}

	if employee.UserID == 0 {
		return employee, errors.New("no user logged in")
	}

	return employee, nil
}

func (s *jobsService) GetJobByID(inputID int) (entity.Jobs, error) {
	jobs, err := s.jobsRepository.FindJobsByID(inputID)
	if err != nil {
		return jobs, err
	}

	if jobs.ID == 0 {
		return jobs, errors.New("No user logged in")
	}

	return jobs, nil
}

func (s *jobsService) AllJobs() ([]entity.Jobs, error) {
	jobs, err := s.jobsRepository.GetAllJob()
	if err != nil {
		return jobs, err
	}
	return jobs, nil
}

func (s *jobsService) CreateJobs(jobs dto.CreateJobsDTO, userID int) (entity.Jobs, error) {
	createJob := entity.Jobs{}
	createJob.CompanyName = jobs.CompanyName
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

func (s *jobsService) UpdateJob(inputData dto.CreateJobsDTO, inputID dto.JobDetailDTO, userID int) (entity.Jobs, error) {
	job, err := s.jobsRepository.FindJobsByID(int(inputID.ID))
	if err != nil {
		return job, nil
	}

	job.CompanyName = inputData.CompanyName
	job.Title = inputData.Title
	job.Description = inputData.Description
	job.Type = inputData.Type
	job.Location = inputData.Location
	job.LevelOfExperience = inputData.LevelOfExperience
	job.Salary = inputData.Salary
	job.DateStart = inputData.DateStart
	job.DateEnd = inputData.DateEnd
	job.PostedBy = uint64(userID)

	updateJob, err := s.jobsRepository.Update(job)
	if err != nil {
		return job, nil
	}

	return updateJob, nil
}

func (s *jobsService) DeletedJob(inputID int) (entity.Jobs, error) {
	jobs, err := s.jobsRepository.DeleteJob(inputID)
	if err != nil {
		return jobs, err
	}

	return jobs, nil
}
