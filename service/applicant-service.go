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
	UpdateApplicant(applicant dto.ApplicantUpdateDTO) entity.Applicant
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

func (service *applicantService) UpdateApplicant(a dto.ApplicantUpdateDTO) entity.Applicant {
	applicant := entity.Applicant{}
	err := smapping.FillStruct(&applicant, smapping.MapFields(&a))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.applicantRepository.SaveApplicant(applicant)
	return res
}
