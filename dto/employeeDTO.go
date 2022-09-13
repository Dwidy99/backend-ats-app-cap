package dto

type EmployeeDTO struct {
	ID     uint64 `json:"applicant_id"`
	UserID uint64 `json:"user_id"`
}

type EmployeeUpdateDTO struct {
	ID     uint64 `json:"applicant_id"`
	UserID uint64 `json:"user_id"`
}