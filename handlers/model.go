package handlers

type registerRequest struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Mobile string `json:"mobile"`
	Avatar string `json:"avatar"`
	Password string `json:"password"`
}

type loginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}