package v1

type AddAdminUserRequest struct {
	Username string   `json:"username" binding:"required" example:"admin"`
	Email    string   `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Password string   `json:"password" binding:"required" example:"123456"`
	RoleTags []string `json:"roleTags"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"123456"`
}
type LoginResponseData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type LoginResponse struct {
	Response
	Data LoginResponseData
}

type UpdateProfileRequest struct {
	Email string `json:"email" binding:"required,email" example:"1234@gmail.com"`
}
type GetProfileResponseData struct {
	UserId   string   `json:"userId"`
	Username string   `json:"username" example:"alan"`
	Roles    []string `json:"roles"`
}
type GetProfileResponse struct {
	Response
	Data GetProfileResponseData
}

type SetUserRoleRequest struct {
	RoleTags []string `json:"roleTags"`
	UserId   string   `json:"userId"`
}
