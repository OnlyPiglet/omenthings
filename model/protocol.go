package model

type Protocol struct {
	Name       string                 `gorm:"column:name;type:varchar(255);not null;comment:'协议名称'" json:"name"`
	ConfigData string                 `gorm:"type:text;column:config_data;comment:'协议配置信息'"`
	Config     map[string]interface{} `gorm:"-" json:"config"` //transform configData to config
}

func (p *Protocol) TableName() string {
	return "protocol"
}
