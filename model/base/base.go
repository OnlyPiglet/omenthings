package base

import "time"

type Base struct {
	CreateTime time.Time
	UpdateTime time.Time
	DeleteTime time.Time
	Uuid       string
	TenantId   string
	OwnerId    string
}
