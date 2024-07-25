package v1

type AddAdminUserRequest struct {
	Username string   `json:"username" binding:"required" example:"admin"`
	Email    string   `json:"email" binding:"required,email" example:"1234@gmail.com"`
	Password string   `json:"password" binding:"required" example:"123456"`
	RoleIds  []string `json:"role_ids"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"123456"`
}
type LoginResponseData struct {
	AccessToken string `json:"accessToken"`
}
type LoginResponse struct {
	Response
	Data LoginResponseData
}

type UpdateProfileRequest struct {
	Email string `json:"email" binding:"required,email" example:"1234@gmail.com"`
}
type GetProfileResponseData struct {
	UserId   string `json:"userId"`
	Username string `json:"username" example:"alan"`
}
type GetProfileResponse struct {
	Response
	Data GetProfileResponseData
}

type SetUserRoleRequest struct {
	RoleIds []string `json:"role_ids"`
	UserId  string   `json:"user_id"`
}
