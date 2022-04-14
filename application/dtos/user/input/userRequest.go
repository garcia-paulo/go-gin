package input_user

import "gopkg.in/validator.v2"

type UserRequest struct {
	Username string `json:"username" validate:"min=4,max=16"`
	Password string `json:"password" validate:"min=6,max=18"`
}

func (u *UserRequest) Validate() error {
	if err := validator.Validate(u); err != nil {
		return err
	}
	return nil
}
