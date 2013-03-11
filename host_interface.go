package zabbix

type (
	InterfaceType int
)

const (
	Agent InterfaceType = 1
	SNMP  InterfaceType = 2
	IPMI  InterfaceType = 3
	JMX   InterfaceType = 4
)

// https://www.zabbix.com/documentation/2.0/manual/appendix/api/hostinterface/definitions
type HostInterface struct {
	DNS   string        `json:"dns"`
	IP    string        `json:"ip"`
	Main  bool          `json:"main"`
	Port  string        `json:"port"`
	Type  InterfaceType `json:"type"`
	UseIP bool          `json:"useip"`
}

type HostInterfaces []HostInterface
