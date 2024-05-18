package request

type UserCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Address  string `json:"address" `
	Phone    string `json:"phone" `
	Role     string `json:"role"`
}

type UserUpdateRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UserResponse struct {
	ID   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

func (UserResponse) TableName() string {
	return "users"
}
