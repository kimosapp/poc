package organization

type CreateOrganizationRequest struct {
	FirstName        string `json:"firstName" binding:"required"`
	LastName         string `json:"lastName" binding:"required"`
	Password         string `json:"password" binding:"required"`
	ConfirmPassword  string `json:"confirmPassword" binding:"required"`
	OrganizationName string `json:"organizationName" binding:"required"`
	BillingEmail     string `json:"billingEmail" binding:"required"`
	Captcha          string `json:"captcha" binding:"required"`
}
