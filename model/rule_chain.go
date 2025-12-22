package model

import "github.com/OnlyPiglet/omenthings/model/base"

type RuleChain struct {
	base.Base
	Name string `gorm:"comment:'名称';column:name"`
}

func (rc *RuleChain) TableName() string {
	return "rule_chain"
}
