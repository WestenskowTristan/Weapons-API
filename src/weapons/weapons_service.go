package weapons

import "time"


type IWeaponsService interface {
	Get(weaponId string) (*Weapon, error)
}

type WeaponsService struct {}

func InitWeaponsService() IWeaponsService {
	return &WeaponsService{}
}

func (service *WeaponsService) Get(weaponId string) (*Weapon, error) {
	weapon := Weapon{
		Id: weaponId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserId: "5678",
		Image: []byte{},
		Name: "AR 15",
		Ammo: FiveFiveSix,
	}
	return &weapon, nil
}