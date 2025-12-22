package magic

const (
	SuperTenantName string = "超管租户"
	SuperTenantId   string = "2b184c5f-2025-0418-1040-711db7cbb485"
)

type DeviceType string

const (
	DirectDevice DeviceType = "direct"
	ParentDevice DeviceType = "parent_device"
	ShadowDevice DeviceType = "shadow_device"
	AvatarDevice DeviceType = "avatar_device"
)

type CredentialType string

const (
	MqttBasic CredentialType = "MQTT_BASIC"
)
