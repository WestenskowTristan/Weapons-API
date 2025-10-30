package weapons

import (
	"time"
)

type WeaponAmmo string 

var (
	FiveFiveSix WeaponAmmo = "556"
	SevenSixTwo WeaponAmmo = "762"
	NineMillimeter WeaponAmmo = "9mm"
)


type Weapon struct {
	Id					string				`json:"id"`
	CreatedAt  	time.Time   	`json:"createdAt"`
	UpdatedAt  	time.Time  		`json:"updatedAt"`
	UserId			string				`json:"userId"`
	Image				[]byte				`json:"image"`
	Name				string				`json:"name"`
	Ammo 				WeaponAmmo		`json:"ammo"`
}