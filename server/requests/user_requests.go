package requests

type LoginRequest struct {
	Name     string `json:"name" validate:"required" example:"John Doe"`
	Password string `json:"password" validate:"required" example:"11111111"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required" example:"John Doe"`
	Password string `json:"password" validate:"required" example:"11111111"`
}

type RefreshRequest struct {
	Token string `json:"token" validate:"required" example:"access_token"`
}
