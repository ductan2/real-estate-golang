package vo

type InvestorCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	Phone       string `json:"phone" binding:"required"`
	Website     string `json:"website"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Background  string `json:"background"`
	Type        string `json:"type" binding:"required"`
}

type InvestorUpdateRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Email       string `json:"email" binding:"omitempty,email"`
	Phone       string `json:"phone"`
	Website     string `json:"website"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Background  string `json:"background"`
	Type        string `json:"type"`
}
