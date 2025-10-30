package weapons

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

type WeaponInput struct {
	UserId			string				`json:"userId"`
	Image				[]byte				`json:"image"`
	Name				string				`json:"name"`
	Ammo 				WeaponAmmo		`json:"ammo"`
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func WeaponInputFrom(raw io.Reader) (*WeaponInput, error) {
	value := WeaponInput{}
	err := json.NewDecoder(raw).Decode(&value)
	if err != nil {
		return nil, err
	}

	err = validate.Struct(&value)
	if err != nil {
		return nil, err
	}

	return &value, nil
}