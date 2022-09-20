package repository

import (
	"mini-project/entity"

	"github.com/jinzhu/gorm"
)

type CompanyRepository interface {
	GetAllCompany() ([]entity.Company, error)
	GetCompanyByID(inputID int) (entity.Company, error)
	InsertCompany(company entity.Company) (entity.Company, error)
	UpdateCompany(company entity.Company) (entity.Company, error)
	DeleteCompany(inputID int) (entity.Company, error)
}

type companyConnection struct {
	connection *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) CompanyRepository {
	return &companyConnection{
		connection: db,
	}
}

func (db *companyConnection) GetAllCompany() ([]entity.Company, error) {
	var companies []entity.Company

	err := db.connection.Raw("SELECT id, name, email, address, contact, website, created_at FROM company").Scan(&companies).Error
	if err != nil {
		return companies, err
	}

	return companies, nil
}

func (db *companyConnection) GetCompanyByID(inputID int) (entity.Company, error) {
	var company entity.Company

	err := db.connection.Where("id = ?", inputID).Find(&company).Error
	if err != nil {
		return company, err
	}
	return company, nil
}

func (db *companyConnection) InsertCompany(company entity.Company) (entity.Company, error) {
	err := db.connection.Save(&company).Error
	if err != nil {
		return company, err
	}
	return company, nil
}

func (db *companyConnection) UpdateCompany(company entity.Company) (entity.Company, error) {
	err := db.connection.Save(&company).Error
	if err != nil {
		return company, err
	}
	return company, nil
}

func (db *companyConnection) DeleteCompany(inputID int) (entity.Company, error) {
	var company entity.Company
	err := db.connection.Where("id = ?", inputID).Delete(&company).Error
	if err != nil {
		return company, err
	}
	return company, nil
}
