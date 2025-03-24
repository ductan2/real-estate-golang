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

type UserResendOtpRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type UserUpdateInfoRequest struct {
	Username  *string `json:"username"`
	Phone     *string `json:"phone"`
	Address   *string `json:"address"`
	Country   *string `json:"country"`
	City      *string `json:"city"`
	BirthDate *string `json:"birth_date"`
	Gender    *string `json:"gender"`
	Bio       *string `json:"bio"`
	Avatar    *string `json:"avatar"`
	Province  *string `json:"province"`
	District  *string `json:"district"`
	Ward      *string `json:"ward"`
}

type UserUpdatePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type UserUpdateEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}