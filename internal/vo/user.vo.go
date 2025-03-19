package vo


type UserLoginRequest struct {
	Email   string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required,min=1"`
	Email   string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Purpose  *string `json:"purpose"` // TEST_USER, TRADER, ADMIN, etc.
}

type UserVerifyOtpRequest struct {
	Email string `json:"email" binding:"required,email"`
	Otp   int    `json:"otp" binding:"required"`
}

type UserUpdateInfoRequest struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Phone     *string `json:"phone"`
	Address   *string `json:"address"`
	Country   *string `json:"country"`
	City      *string `json:"city"`
	BirthDate *string `json:"birth_date"`
	Gender    *string `json:"gender"`
	Bio       *string `json:"bio"`
}
