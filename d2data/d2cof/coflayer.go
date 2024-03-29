package d2cof

import "github.com/OpenDiablo2/D2Shared/d2common/d2enum"

type CofLayer struct {
	Type        d2enum.CompositeType
	Shadow      byte
	Transparent bool
	DrawEffect  d2enum.DrawEffect
	WeaponClass d2enum.WeaponClass
}
