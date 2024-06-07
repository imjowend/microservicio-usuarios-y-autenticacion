package handler

type UserProfileDTO struct {
	UserID   int64 `json:"userId"`
	Username string
	Email    string
	Password string
}
