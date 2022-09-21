package dto

type Jobskill struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type GetSkillDetailDTO struct {
	ID int `uri:"id" binding:"required"`
}