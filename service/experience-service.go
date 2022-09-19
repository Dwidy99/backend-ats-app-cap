package service

import (
	"errors"
	"log"
	"mini-project/dto"
	"mini-project/entity"
	"mini-project/repository"

	"github.com/mashingan/smapping"
)

type ExperienceService interface {
	GetApplicantByID(userID int) (entity.Applicant, error)
	GetUserByID(userID int) (entity.User, error)
	GetExperienceByIdApplicant(idApplicant int) (entity.Jobexperience, error)
	GetExperienceByID(inputID int) (entity.Jobexperience, error)
	GetAllExperiences(userID int) ([]entity.Jobexperience, error)
	CreateExperience(experience dto.CreateExperienceDTO, applicantID int) (entity.Jobexperience, error)
	UpdateExperience(inputID int, inputData dto.CreateExperienceDTO) (entity.Jobexperience, error)
	DeleteExperience(inputID int) (entity.Jobexperience, error)
}

type experienceService struct {
	experienceRepository repository.ExperienceRepository
}

func NewExperienceService(expRepository repository.ExperienceRepository) ExperienceService {
	return &experienceService{
		experienceRepository: expRepository,
	}
}

func (s *experienceService) GetApplicantByID(userID int) (entity.Applicant, error) {
	applicant, err := s.experienceRepository.FindApplicantByID(userID)
	if err != nil {
		return applicant, err
	}
	
	if applicant.UserID == 0 {
		return applicant, errors.New("no user logged in")
	}

	return applicant, nil
}

func (s *experienceService) GetUserByID(userID int) (entity.User, error) {
	user, err := s.experienceRepository.FindUserByID(userID)
	if err != nil {
		return user, err
	}
	
	if user.ID == 0 {
		return user, errors.New("no user logged in")
	}
	
	return user, nil
}

func (s *experienceService) GetAllExperiences(userID int) ([]entity.Jobexperience, error) {
	if userID != 0 {
		applicant, _ := s.GetApplicantByID(userID)
		experience, err := s.experienceRepository.GetAllExperienceByID(int(applicant.ID))
		if err != nil {
			return experience, err
		}

		return experience, nil
	}

	experience, err := s.experienceRepository.GetAllExperience()
	if err != nil {
		return experience, err
	}

	return experience, nil
}

// func (s *experienceService) GetExperienceByID(applicantID int) (entity.Jobexperience, error) {

// }

func (s *experienceService) GetExperienceByIdApplicant(idApplicant int) (entity.Jobexperience, error) {
	experience, err := s.experienceRepository.FindExperienceByIdApplicant(idApplicant)
	if err != nil {
		return experience, err
	}

	if experience.ID == 0 {
		return experience, errors.New("no user logged in")
	}

	return experience, nil
}

func (s *experienceService) GetExperienceByID(inputID int) (entity.Jobexperience, error) {
	experience, err := s.experienceRepository.FindExperienceByID(inputID)
	if err != nil {
		return experience, err
	}

	if experience.ID == 0 {
		return experience, errors.New("no user logged in")
	}

	return experience, nil
}

func (service *experienceService) CreateExperience(experienceInput dto.CreateExperienceDTO, applicantID int) (entity.Jobexperience, error) {
	createJobExperience := entity.Jobexperience{}
	createJobExperience.ApplicantID = uint64(applicantID)
	createJobExperience.CompanyName = experienceInput.CompanyName
	createJobExperience.DateStart = experienceInput.DateStart
	createJobExperience.DateEnd = experienceInput.DateEnd
	createJobExperience.Description = experienceInput.Description
	createJobExperience.Role = experienceInput.Role
	createJobExperience.Status = experienceInput.Status

	err := smapping.FillStruct(&createJobExperience, smapping.MapFields(&experienceInput))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res, err := service.experienceRepository.InsertExperience(createJobExperience)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *experienceService) UpdateExperience(inputID int, inputData dto.CreateExperienceDTO) (entity.Jobexperience, error) {
	experience, err := s.experienceRepository.FindExperienceByID(inputID)
	if err != nil {
		return experience, err
	}
	
	if experience.ID != uint64(inputID) {
		return experience, errors.New("not a user applicant owner experience")
	}

	experience.CompanyName = inputData.CompanyName
	experience.DateStart = inputData.DateStart
	experience.DateEnd = inputData.DateEnd
	experience.Description = inputData.Description
	experience.Role = inputData.Role
	experience.Status = inputData.Status

	updateExperience, err := s.experienceRepository.Update(experience)
	if err != nil {
		return experience, nil
	}

	return updateExperience, nil
}

func (s *experienceService) DeleteExperience(inputID int) (entity.Jobexperience, error) {
	campaignDeleted, err := s.experienceRepository.DeleteExperience(inputID)
	if err != nil {
		return campaignDeleted, err
	}
	return campaignDeleted, nil
}