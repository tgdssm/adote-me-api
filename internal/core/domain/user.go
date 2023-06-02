package domain

import (
	"api/helpers"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

type User struct {
	ID        uint64 `json:"id,omitempty"` // omitempty não deixa passar o valor zero do uint64 para o json
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Cellphone string `json:"cellphone,omitempty"`
	Passwd    string `json:"passwd,omitempty"`
	//ProfileImage ProfileImage `json:"profile_image,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Token     string    `json:"token,omitempty"`
}

func (u *User) Prepare(pwCanBeEmpty bool) error {
	if err := u.validator(pwCanBeEmpty); err != nil {
		return err
	}

	if err := u.format(pwCanBeEmpty); err != nil {
		return err
	}
	return nil
}

func (u *User) validator(pwCanBeEmpty bool) error {
	if u.Name == "" {
		return errors.New("the name is mandatory and cannot be blank")
	}
	if u.Email == "" {
		return errors.New("the email is mandatory and cannot be blank")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("the email entered is invalid")
	}
	if !pwCanBeEmpty && u.Passwd == "" {
		return errors.New("the password is mandatory and cannot be blank")
	}
	if u.Cellphone == "" {
		return errors.New("the cellphone number is mandatory and cannot be blank")
	}

	if len(u.Cellphone) > 15 || len(u.Cellphone) < 9 {
		return errors.New("number of invalid digits for a mobile number")
	}

	return nil
}

func (u *User) format(pwCanBeEmpty bool) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)

	if !pwCanBeEmpty {
		passwdWithHash, err := helpers.Hash(u.Passwd)
		if err != nil {
			return err
		}
		u.Passwd = string(passwdWithHash)
	}
	return nil
}
