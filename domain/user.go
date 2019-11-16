package domain

type User struct {
	ID       int64  `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique_index"`
	Password string `json:"password"`
	PIN      string `json:"pin"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" gorm:"unique_index" binding:"required"`
	Password string `json:"password" binding:"required"`
	PIN      string `json:"pin" binding:"required"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

type ChangePINRequest struct {
	OldPIN string `json:"old_pin" binding:"required"`
	NewPIN string `json:"new_pin" binding:"required"`
}

type UserResponse struct {
	ID    int64  `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique_index"`
}
