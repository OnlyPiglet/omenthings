package model

import (
	"github.com/OnlyPiglet/omenthings/magic"
	"github.com/OnlyPiglet/omenthings/model/base"
)

type Device struct {
	base.Base
	Type           magic.DeviceType
	ProductId      string
	RuleChainId    string
	ParentDeviceId string
}

func (p *Device) TableName() string {
	return "device"
}
