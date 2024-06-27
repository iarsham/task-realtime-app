package helpers

type BadRequest struct {
	Error string `json:"error" example:"bad request"`
}

type UserNotFound struct {
	Error string `json:"error" example:"user not found"`
}

type InvalidPassword struct {
	Error string `json:"error" example:"invalid password"`
}

type InternalServerError struct {
	Error string `json:"error" example:"internal server error"`
}

type AccessToken struct {
	AccessToken string `json:"access-token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`
}

type UserCreated struct {
	Response string `json:"response" example:"user created successfully"`
}

type EmailAlreadyExists struct {
	Error string `json:"error" example:"email already exists"`
}
