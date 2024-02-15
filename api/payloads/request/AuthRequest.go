package request

import "errors"

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (e *LoginRequest) Validate() error {
	if e.Email == "" {
		return errors.New("email is required")
	} else if e.Password == "" {
		return errors.New("password is required")
	} else if len(e.Password) < 8 || len(e.Password) > 30 {
		return errors.New("password must be between 8 and 30 characters")
	}

	return nil
}

type RegisterRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	NoTelp    string `json:"no_telp" binding:"required"`
	Photo     string `json:"photo" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func (e *RegisterRequest) Validate() error {
	if e.FirstName == "" {
		return errors.New("first name is required")
	} else if e.LastName == "" {
		return errors.New("last name is required")
	} else if e.Username == "" {
		return errors.New("username is required")
	} else if e.Email == "" {
		return errors.New("email is required")
	} else if e.NoTelp == "" {
		return errors.New("no telp is required")
	} else if e.Photo == "" {
		return errors.New("photo is required")
	} else if e.Password == "" {
		return errors.New("password is required")
	} else if len(e.Password) < 8 || len(e.Password) > 30 {
		return errors.New("password must be between 8 and 30 characters")
	}

	return nil
}

type ForgotPasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func (e *ForgotPasswordRequest) Validate() error {
	if e.OldPassword == "" {
		return errors.New("old password is required")
	} else if e.NewPassword == "" {
		return errors.New("new password is required")
	} else if len(e.NewPassword) < 8 || len(e.NewPassword) > 30 {
		return errors.New("password must be between 8 and 30 characters")
	}

	return nil
}

type ResetPasswordRequest struct {
	Email    string `json:"email" binding:"required"`
}

func (e *ResetPasswordRequest) Validate() error {
	if e.Email == "" {
		return errors.New("email is required")
	}

	return nil
}
