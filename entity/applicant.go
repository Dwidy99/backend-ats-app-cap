package entity

type Applicant struct {
	ID            int    `json:"id"`
	UserID        int    `json:"user_id"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Avatar        string `json:"avatar"`
	AccountStatus int    `json:"account_status"`
	LastEducation string `json:"last_education"`
	LinkedURL     string `json:"linkedin_url"`
	GithubURL     string `json:"github_url"`
}