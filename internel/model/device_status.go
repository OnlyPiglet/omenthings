package model

type DeviceStatus struct {
	TenantId        string
	DeviceId        string
	ClusterId       string
	InstanceId      string
	LastConnectTime int64
	LastActiveTime  int64
}
