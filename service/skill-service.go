package service

import (
	"errors"
	"log"
	"mini-project/dto"
	"mini-project/entity"
	"mini-project/repository"

	"github.com/mashingan/smapping"
)

type SkillService interface {
	GetUserByID(userID int) (entity.User, error)
	CreateSkill(skillInput dto.Jobskill, applicantID int) (entity.Jobskill, error)
	GetApplicantByID(userID int) (entity.Applicant, error)
	GetEmployeeByID(userID int) (entity.Employee, error)
	GetSkillByID(inputID int) (entity.Jobskill, error)
	UpdateSkill(inputID int, inputData dto.Jobskill, applicantID int, userID int) (entity.Jobskill, error)
	GetSkillDetailByID(inputID int, applicantID int) (entity.Jobskill, error)
	DeleteSkill(inputID int, applicantID int) (entity.Jobskill, error)
	GetSkills(applicantID int) ([]entity.JobSkillsDetailFormatter, error)
}

type skillService struct {
	skillRepository repository.SkillRepository
}

func NewSkillService(skillRepository repository.SkillRepository) SkillService {
	return &skillService{
		skillRepository: skillRepository,
	}
}

func (s *skillService) GetSkills(applicantID int) ([]entity.JobSkillsDetailFormatter, error) {
	skills := []entity.JobSkillsDetailFormatter{}

	getSkills, err := s.skillRepository.GetSkills(applicantID)
	if err != nil {
		return skills, err
	}
	
	for _, getSkill := range getSkills {
		skillApplicant, _ := s.GetSkillByID(int(getSkill.SkillID))
		skills = append(skills, entity.JobSkillsDetailFormatter{
			ApplicantID: getSkill.ApplicantID,
			SkillID: getSkill.SkillID,
			Name: skillApplicant.Name,
		})
	}

	return skills, nil
}

func (s *skillService) GetSkillDetailByID(inputID int, applicantID int) (entity.Jobskill, error) {
	var jobSkills entity.Jobskill

	jobSkillApplicant, err := s.skillRepository.GetJobSkillApplicantBySkillID(inputID)
	if err != nil {
		return jobSkills, err
	}

	
	jobSkill, err := s.skillRepository.GetSkillByID(inputID)
	if err != nil {
		return jobSkill, err
	}
	
	for _, jobSkillApp := range jobSkillApplicant {
		if int(jobSkillApp.ApplicantID) != applicantID {
			return jobSkill, errors.New("not an owner skill job")
		}
	}

	return jobSkill, nil
}

func (s *skillService) CreateSkill(skillInput dto.Jobskill, applicantID int) (entity.Jobskill, error) {
	createSkill := entity.Jobskill{}
	createJobSkillApplicant := entity.Jobskillapplicant{}

	createSkill.Name = skillInput.Name
	
	err := smapping.FillStruct(&createSkill, smapping.MapFields(&skillInput))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res, err := s.skillRepository.CreateSkill(createSkill, createJobSkillApplicant, applicantID)
	if err != nil {
		return res, err
	}
	
	return res, nil
}

func (s *skillService) GetUserByID(userID int) (entity.User, error) {
	user, err := s.skillRepository.FindUserByID(userID)
	if err != nil {
		return user, err
	}
	
	if user.ID == 0 {
		return user, errors.New("no user logged in")
	}
	
	return user, nil
}

func (s *skillService) GetApplicantByID(userID int) (entity.Applicant, error) {
	applicant, err := s.skillRepository.FindApplicantByID(userID)
	if err != nil {
		return applicant, err
	}
	
	if applicant.UserID == 0 {
		return applicant, errors.New("no user logged in")
	}

	return applicant, nil
}

func (s *skillService) GetEmployeeByID(userID int) (entity.Employee, error) {
	employee, err := s.skillRepository.FindEmployeeByID(userID)
	if err != nil {
		return employee, err
	}
	
	if employee.UserID == 0 {
		return employee, errors.New("no user logged in")
	}

	return employee, nil
}

func (s *skillService) GetSkillByID(inputID int) (entity.Jobskill, error) {
	applicant, err := s.skillRepository.GetSkillByID(inputID)
	if err != nil {
		return applicant, err
	}

	return applicant, nil
}

func (s *skillService) UpdateSkill(inputID int, inputData dto.Jobskill, applicantID int, userID int) (entity.Jobskill, error) {
	var jobSkills entity.Jobskill

	jobSkillApplicant, err := s.skillRepository.GetJobSkillApplicantBySkillID(inputID)
	if err != nil {
		return jobSkills, err
	}

	jobSkill, err := s.skillRepository.GetSkillByID(inputID)
	if err != nil {
		return jobSkill, err
	}

	for _, jobSkillApp := range jobSkillApplicant {
		if int(jobSkillApp.ApplicantID) != applicantID {
			return jobSkill, errors.New("not an owner skill job")
		}
	}

	jobSkill.Name = inputData.Name

	skill, err := s.skillRepository.Update(jobSkill)
	if err != nil {
		return skill, nil
	}

	return skill, nil
}

func (s *skillService) DeleteSkill(inputID int, applicantID int) (entity.Jobskill, error) {
	var jobSkills entity.Jobskill

	jobSkillApplicant, err := s.skillRepository.GetJobSkillApplicantBySkillID(inputID)
	if err != nil {
		return jobSkills, err
	}

	jobSkill, err := s.skillRepository.GetSkillByID(inputID)
	if err != nil {
		return jobSkill, err
	}

	for _, jobSkillApp := range jobSkillApplicant {
		if int(jobSkillApp.ApplicantID) != applicantID {
			return jobSkill, errors.New("not an owner skill job")
		}
	}

	skill, err := s.skillRepository.Delete(inputID)
	if err != nil {
		return skill, err
	}
	return skill, nil
}