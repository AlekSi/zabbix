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

// https://www.zabbix.com/documentation/2.2/manual/appendix/api/hostinterface/definitions
type HostInterface struct {
	DNS   string        `json:"dns"`
	IP    string        `json:"ip"`
	Main  int           `json:"main"`
	Port  string        `json:"port"`
	Type  InterfaceType `json:"type"`
	UseIP int           `json:"useip"`
}

type HostInterfaces []HostInterface
