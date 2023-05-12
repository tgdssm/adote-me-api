package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID          uint64    `json:"id,omitempty"` // omitempty não deixa passar o valor zero do uint64 para o json
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	Passwd      string    `json:"passwd,omitempty"`
	PicturePath string    `json:"picture,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}

func (u *User) Prepare() error {
	if err := u.validator(); err != nil {
		return err
	}

	u.format()
	return nil
}

func (u User) validator() error {
	if u.Name == "" {
		return errors.New("The name is mandatory and cannot be blank")
	}
	if u.Email == "" {
		return errors.New("The email is mandatory and cannot be blank")
	}
	if u.Passwd == "" {
		return errors.New("The password is mandatory and cannot be blank")
	}
	if u.PicturePath == "" {
		return errors.New("The picture path is mandatory and cannot be blank")
	}

	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
}
