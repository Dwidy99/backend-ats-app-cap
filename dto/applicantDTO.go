package dto

import "os/user"

type ApplicantDTO struct {
	UserID uint64 `uri:"id" binding:"required"`
}

type ApplicantUpdateDTO struct {
	FirstName     string `json:"first_name" form:"first_name" binding:"required"`
	LastName      string `json:"last_name" form:"last_name" binding:"required"`
	Phone         string `json:"phone" form:"phone" binding:"required"`
	LastEducation string `json:"last_education" form:"last_education" binding:"required"`
	LinkedinURL   string `json:"linkedin_url" form:"linkedin_url" binding:"required"`
	GithubURL     string `json:"github_url" form:"github_url" binding:"required"`
	User          user.User
}
