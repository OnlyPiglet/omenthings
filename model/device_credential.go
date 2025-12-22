package model

import (
	"github.com/OnlyPiglet/omenthings/magic"
	"github.com/OnlyPiglet/omenthings/model/base"
)

type DeviceCredential struct {
	base.Base
	DeviceId        string `gorm:"column:device_id"`
	CredentialType  magic.DeviceType
	CredentialId    string `gorm:"column:credential_id"`
	CredentialValue string `gorm:"column:credential_value"`
}

func (dc *DeviceCredential) TableName() string {
	return "device_credential"
}
