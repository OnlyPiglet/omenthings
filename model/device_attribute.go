package model

import "github.com/OnlyPiglet/omenthings/model/base"

type DeviceAttribute struct {
	base.Base
	DeviceId   string      `gorm:"column:device_id"`
	Key        string      `json:"key" gorm:"column:key"`
	StrValue   string      `json:"str_v" gorm:"column:str_v"`
	LongValue  int64       `json:"long_v" gorm:"column:long_v"`
	floatValue float64     `json:"float_v" gorm:"column:float_v"`
	BoolValue  bool        `json:"bool_v" gorm:"column:bool_v"`
	JsonValue  string      `json:"json_v" gorm:"column:json_v"`
	Value      interface{} `json:"value" gorm:"-"`
}

func (ad *DeviceAttribute) TableName() string {
	return "device_attribute"
}
