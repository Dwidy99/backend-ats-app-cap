package service

import (
	"log"
	"mini-project/dto"
	"mini-project/entity"
	"mini-project/repository"

	"github.com/mashingan/smapping"
)

type CompanyService interface {
	All() ([]entity.Company, error)
	FindByID(companyID int) (entity.Company, error)
	Insert(companyID int, company dto.CompanyDTO) (entity.Company, error)
	Update(companyID int, company dto.CompanyDTO) (entity.Company, error)
	Delete(companyID int) (entity.Company, error)
}

type companyService struct {
	companyRepository repository.CompanyRepository
}

func NewCompanyService(compRepository repository.CompanyRepository) CompanyService {
	return &companyService{
		companyRepository: compRepository,
	}
}

func (service *companyService) All() ([]entity.Company, error) {
	company, err := service.companyRepository.AllCompany()
	if err != nil {
		return company, err
	}
	return company, nil
}

func (service *companyService) FindByID(companyID int) (entity.Company, error) {
	company, err := service.companyRepository.FindCompanyByID(companyID)
	if err != nil {
		return company, err
	}
	return company, nil
}

func (service *companyService) Insert(companyID int, company dto.CompanyDTO) (entity.Company, error) {
	inputCompany := entity.Company{}
	inputCompany.ID = uint64(companyID)
	inputCompany.Name = company.Name
	inputCompany.Email = company.Email
	inputCompany.Address = company.Address
	inputCompany.Contact = company.Contact
	inputCompany.Website = company.Website

	err := smapping.FillStruct(&inputCompany, smapping.MapFields(&company))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res, err := service.companyRepository.InsertCompany(inputCompany)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (service *companyService) Update(companyID int, company dto.CompanyDTO) (entity.Company, error) {
	comp, err := service.companyRepository.FindCompanyByID(companyID)
	if err != nil {
		return comp, err
	}
	comp.Name = company.Name
	comp.Email = company.Email
	comp.Address = company.Address
	comp.Contact = company.Contact
	comp.Website = company.Website

	updateCompany, err := service.companyRepository.UpdateCompany(comp)
	if err != nil {
		return comp, nil
	}
	return updateCompany, nil
}

func (service *companyService) Delete(companyID int) (entity.Company, error) {
	companyDeleted, err := service.companyRepository.DeleteCompany(companyID)
	if err != nil {
		return companyDeleted, err
	}
	return companyDeleted, nil
}
