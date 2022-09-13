package service

import (
	"fmt"
	"mini-project/dto"
	"mini-project/entity"
	"mini-project/repository"
)

type ApplicantService interface {
	IsAllowedToEdit(userID string, applicantUserID uint64) bool
	UpdateApplicant(applicant dto.ApplicantDTO, userId int, inputData dto.ApplicantUpdateDTO) entity.Applicant
}

type applicantService struct {
	applicantRepository repository.ApplicantRepository
}

func NewApplicantService(applicantRepo repository.ApplicantRepository) ApplicantService {
	return &applicantService{
		applicantRepository: applicantRepo,
	}
}

func (service *applicantService) IsAllowedToEdit(userID string, applicantUserID uint64) bool {
	applicant := service.applicantRepository.FindApplicantByID(applicantUserID)
	id := fmt.Sprintf("%v", applicant.UserID)

	return userID == id
}

func (service *applicantService) UpdateApplicant(inputID dto.ApplicantDTO, userId int, inputData dto.ApplicantUpdateDTO) entity.Applicant {
	applicant := service.applicantRepository.FindApplicantByID(inputID.UserID)

	if applicant.UserID == uint64(userId) {
		applicant.FirstName = inputData.FirstName
		applicant.LastName = inputData.LastName
		applicant.Phone = inputData.Phone
		applicant.LastEducation = inputData.LastEducation
		applicant.LinkedURL = inputData.LinkedinURL
		applicant.GithubURL = inputData.GithubURL

		updateApplicant := service.applicantRepository.SaveApplicant(applicant)

		return updateApplicant
	} else {
		return applicant
	}
}
