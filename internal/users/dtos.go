package users

type RegisterUserRequest struct {
	Auth0ID   string `json:"auth0Id" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"firstName" validate:"required,min=1,max=25"`
	LastName  string `json:"lastName" validate:"required,min=1,max=25"`
}
