package model

import "github.com/OnlyPiglet/omenthings/model/base"

type Product struct {
	base.Base
	Name        string `gorm:"comment:'名称';column:name"`
	ProtocolId  string `gorm:"comment:'协议id';column:protocol_id"`
	RuleChainId string `gorm:"comment:'规则链id';column:rule_chain_id"`
}

func (p *Product) TableName() string {
	return "product"
}
