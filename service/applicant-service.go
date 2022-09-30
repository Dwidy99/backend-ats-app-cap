package service

import (
	"errors"
	"fmt"
	"github.com/PutraFajarF/backend-ats-app-cap/dto"
	"github.com/PutraFajarF/backend-ats-app-cap/entity"
	"github.com/PutraFajarF/backend-ats-app-cap/repository"
	"os"

	"github.com/mashingan/smapping"
)

type ApplicantService interface {
	IsAllowedToEdit(userID string, applicantUserID uint64) (bool, error)
	UpdateApplicant(applicant dto.ApplicantUpdateDTO, id int) (entity.Applicant, error)
	GetApplicantByUserID(userId uint64) (entity.Applicant, error)
	GetDetailApplicant(userId uint64) (entity.DetailApplicant, error)
	UploadAvatar(ID int, fileLocation string) (entity.Applicant, error)
}

type applicantService struct {
	applicantRepository repository.ApplicantRepository
}

func NewApplicantService(applicantRepo repository.ApplicantRepository) ApplicantService {
	return &applicantService{
		applicantRepository: applicantRepo,
	}
}

func (service *applicantService) IsAllowedToEdit(userID string, applicantUserID uint64) (bool, error) {
	applicant, err := service.applicantRepository.FindApplicantByUserID(applicantUserID)
	if err != nil {
		return false, err
	}
	id := fmt.Sprintf("%v", applicant.UserID)

	if userID != id {
		return false, errors.New("not allowed to edit")
	}
	return true, nil
}

func (service *applicantService) UpdateApplicant(a dto.ApplicantUpdateDTO, id int) (entity.Applicant, error) {
	applicant, err := service.applicantRepository.FindApplicantByUserID(uint64(id))
	if err != nil {
		return applicant, err
	}

	applicant.FirstName = a.FirstName
	applicant.LastName = a.LastName
	applicant.Phone = a.Phone
	applicant.LastEducation = a.LastEducation
	applicant.LinkedURL = a.LinkedinURL
	applicant.GithubURL = a.GithubURL

	err = smapping.FillStruct(&applicant, smapping.MapFields(&a))
	if err != nil {
		return applicant, err
	}

	res, err := service.applicantRepository.SaveApplicant(applicant)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *applicantService) GetApplicantByUserID(userId uint64) (entity.Applicant, error) {
	res, err := s.applicantRepository.FindApplicantByUserID(userId)
	if err != nil {
		return res, err
	}
	err = smapping.FillStruct(&userId, smapping.MapFields(&userId))
	if err != nil {
		return res, err
	}

	return res, nil
}

func (s *applicantService) GetDetailApplicant(userId uint64) (entity.DetailApplicant, error) {
	createDetailApplicant := entity.DetailApplicant{}

	applicant, err := s.applicantRepository.FindApplicantByUserID(userId)
	if err != nil {
		return createDetailApplicant, err
	}

	// return array
	jobSkillApplicant, err := s.applicantRepository.GetJobSkillApplicantByApplicantID(applicant.ID)
	if err != nil {
		return createDetailApplicant, err
	}

	experience, err := s.applicantRepository.GetExperienceByApplicantID(applicant.ID)
	if err != nil {
		return createDetailApplicant, err
	}

	createDetailApplicant.ID = applicant.ID
	createDetailApplicant.UserID = applicant.UserID
	createDetailApplicant.FirstName = applicant.FirstName
	createDetailApplicant.LastName = applicant.LastName
	createDetailApplicant.Avatar = applicant.Avatar
	createDetailApplicant.Phone = applicant.Phone
	createDetailApplicant.LastEducation = applicant.LastEducation
	createDetailApplicant.LinkedURL = applicant.LinkedURL
	createDetailApplicant.GithubURL = applicant.GithubURL

	for _, ex := range experience {
		createDetailApplicant.JobExperience = append(createDetailApplicant.JobExperience, entity.Jobexperience{
			ID:          ex.ID,
			ApplicantID: ex.ApplicantID,
			CompanyName: ex.CompanyName,
			Description: ex.Description,
			Status:      ex.Status,
			Role:        ex.Role,
			DateStart:   ex.DateStart,
			DateEnd:     ex.DateEnd,
		})
	}

	for _, jobSkillApp := range jobSkillApplicant {
		skill, err := s.applicantRepository.GetJobSkillByJobSkillApplicantID(jobSkillApp.SkillID)
		if err != nil {
			return createDetailApplicant, err
		}
		for _, s := range skill {
			createDetailApplicant.JobSkill = append(createDetailApplicant.JobSkill, entity.Jobskill{
				Name: s.Name,
				ID:   s.ID,
			})
		}
	}

	return createDetailApplicant, nil
}

func (s *applicantService) UploadAvatar(ID int, fileLocation string) (entity.Applicant, error) {
	applicant, err := s.applicantRepository.FindApplicantByUserID(uint64(ID))
	if err != nil {
		return applicant, err
	}
	fmt.Println(applicant.Avatar)

	if applicant.Avatar != "" {
		e := os.Remove(applicant.Avatar)
		if e != nil {
			return applicant, e
		}
	}

	applicant.Avatar = fileLocation

	updatedApplicant, err := s.applicantRepository.SaveApplicant(applicant)
	if err != nil {
		return updatedApplicant, err
	}
	fmt.Println(updatedApplicant.Avatar)

	return updatedApplicant, nil
}
