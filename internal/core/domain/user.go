package domain

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID           uint64       `json:"id,omitempty"` // omitempty não deixa passar o valor zero do uint64 para o json
	Name         string       `json:"name,omitempty"`
	Email        string       `json:"email,omitempty"`
	Cellphone    string       `json:"cellphone,omitempty"`
	Passwd       string       `json:"passwd,omitempty"`
	ProfileImage ProfileImage `json:"profile_image,omitempty"`
	CreatedAt    time.Time    `json:"create_at,omitempty"`
}

func (u *User) Prepare() error {
	if err := u.validator(); err != nil {
		return err
	}

	u.format()
	return nil
}

func (u *User) validator() error {
	if u.Name == "" {
		return errors.New("the name is mandatory and cannot be blank")
	}
	if u.Email == "" {
		return errors.New("the email is mandatory and cannot be blank")
	}
	if u.Passwd == "" {
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

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
}
