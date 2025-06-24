package handler

type AuthService interface {
	Login(username, password string) (string, error)
	Refresh(token string) (string, error)
	Verify(token string) (string, error)
}

type Handler struct {
	authService AuthService
}

func NewHandler(authService AuthService) *Handler {
	return &Handler{authService: authService}
}
