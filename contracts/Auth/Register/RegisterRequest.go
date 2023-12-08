package Register

type RegisterRequest struct {
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
