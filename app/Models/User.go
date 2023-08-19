package Models

type UserModel struct {
	Name        string `json:"name" validate:"required,max=1"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
}
