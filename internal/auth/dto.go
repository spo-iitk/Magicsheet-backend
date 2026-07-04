package auth


type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	Role        string `json:"role"`
}

type MeResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
}