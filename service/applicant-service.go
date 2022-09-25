package service

import (
	"errors"
	"fmt"
	"mini-project/dto"
	"mini-project/entity"
	"mini-project/repository"
	"os"

	"github.com/mashingan/smapping"
)

type ApplicantService interface {
	IsAllowedToEdit(userID string, applicantUserID uint64) (bool, error)
	UpdateApplicant(applicant dto.ApplicantUpdateDTO, id int) (entity.Applicant, error)
	GetApplicantByUserID(userId uint64) (entity.Applicant, error)
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

	applicant.FirstName =  a.FirstName
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

func (s *applicantService)  UploadAvatar(ID int, fileLocation string) (entity.Applicant, error) {
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