package auth

type LoginResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Token    string `json:"token"`
}
