package serializers

type LeadSerializer struct {
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"required"`
	JobTitle       string `json:"job_title" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	PhoneNumber    string `json:"phone_number" binding:"required"`
	Status         string `json:"status" binding:"required"`
	MiddleName     string `json:"middle_name"`
	City           string `json:"city"`
	CurrentCompany string `json:"current_company"`
	CompanyWebsite string `json:"company_website"`
	LinkedIn       string `json:"linked_in"`
}