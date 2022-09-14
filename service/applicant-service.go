package service

import (
	"fmt"
	"log"
	"mini-project/dto"
	"mini-project/entity"
	"mini-project/repository"

	"github.com/mashingan/smapping"
)

type ApplicantService interface {
	IsAllowedToEdit(userID string, applicantUserID uint64) bool
	UpdateApplicant(applicant dto.ApplicantUpdateDTO, id int) entity.Applicant
	GetApplicantByID(userId uint64) entity.Applicant
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

func (service *applicantService) UpdateApplicant(a dto.ApplicantUpdateDTO, id int) entity.Applicant {
	applicant := service.applicantRepository.FindApplicantByID(uint64(id))

	applicant.FirstName =  a.FirstName
	applicant.LastName = a.LastName
	applicant.Phone = a.Phone
	applicant.LastEducation = a.LastEducation
	applicant.LinkedURL = a.LinkedinURL
	applicant.GithubURL = a.GithubURL

	err := smapping.FillStruct(&applicant, smapping.MapFields(&a))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.applicantRepository.SaveApplicant(applicant)
	return res
}

func (s *applicantService) GetApplicantByID(userId uint64) entity.Applicant {
	res := s.applicantRepository.FindApplicantByID(userId)
	err := smapping.FillStruct(&userId, smapping.MapFields(&userId))
	if err != nil {
		log.Fatalf("Failet map %v: ", err)
	}

	return res
}