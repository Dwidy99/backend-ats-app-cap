package dto

type Jobskill struct {
	Name string `json:"name" form:"name" binding:"required"`
}
