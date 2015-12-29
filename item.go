package zabbix

import (
	"fmt"
	"github.com/AlekSi/reflector"
)

type (
	ItemType  int
	ValueType int
	DataType  int
	DeltaType int
)

const (
	ZabbixAgent       ItemType = 0
	SNMPv1Agent       ItemType = 1
	ZabbixTrapper     ItemType = 2
	SimpleCheck       ItemType = 3
	SNMPv2Agent       ItemType = 4
	ZabbixInternal    ItemType = 5
	SNMPv3Agent       ItemType = 6
	ZabbixAgentActive ItemType = 7
	ZabbixAggregate   ItemType = 8
	WebItem           ItemType = 9
	ExternalCheck     ItemType = 10
	DatabaseMonitor   ItemType = 11
	IPMIAgent         ItemType = 12
	SSHAgent          ItemType = 13
	TELNETAgent       ItemType = 14
	Calculated        ItemType = 15
	JMXAgent          ItemType = 16

	Float     ValueType = 0
	Character ValueType = 1
	Log       ValueType = 2
	Unsigned  ValueType = 3
	Text      ValueType = 4

	Decimal     DataType = 0
	Octal       DataType = 1
	Hexadecimal DataType = 2
	Boolean     DataType = 3

	AsIs  DeltaType = 0
	Speed DeltaType = 1
	Delta DeltaType = 2
)

// https://www.zabbix.com/documentation/2.2/manual/appendix/api/item/definitions
type Item struct {
	ItemId      string    `json:"itemid,omitempty"`
	Delay       int       `json:"delay"`
	HostId      string    `json:"hostid"`
	InterfaceId string    `json:"interfaceid,omitempty"`
	Key         string    `json:"key_"`
	Name        string    `json:"name"`
	Type        ItemType  `json:"type"`
	ValueType   ValueType `json:"value_type"`
	DataType    DataType  `json:"data_type"`
	Delta       DeltaType `json:"delta"`
	Description string    `json:"description"`
	Error       string    `json:"error"`
	History     int       `json:"history,omitempty"`
	Trends      int       `json:"trends,omitempty"`

	// Fields below used only when creating applications
	ApplicationIds []string `json:"applications,omitempty"`
}

type Items []Item

// Converts slice to map by key. Panics if there are duplicate keys.
func (items Items) ByKey() (res map[string]Item) {
	res = make(map[string]Item, len(items))
	for _, i := range items {
		_, present := res[i.Key]
		if present {
			panic(fmt.Errorf("Duplicate key %s", i.Key))
		}
		res[i.Key] = i
	}
	return
}

// Wrapper for item.get https://www.zabbix.com/documentation/2.2/manual/appendix/api/item/get
func (api *API) ItemsGet(params Params) (res Items, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	response, err := api.CallWithError("item.get", params)
	if err != nil {
		return
	}

	reflector.MapsToStructs2(response.Result.([]interface{}), &res, reflector.Strconv, "json")
	return
}

// Gets items by application Id.
func (api *API) ItemsGetByApplicationId(id string) (res Items, err error) {
	return api.ItemsGet(Params{"applicationids": id})
}

// Wrapper for item.create: https://www.zabbix.com/documentation/2.2/manual/appendix/api/item/create
func (api *API) ItemsCreate(items Items) (err error) {
	response, err := api.CallWithError("item.create", items)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	itemids := result["itemids"].([]interface{})
	for i, id := range itemids {
		items[i].ItemId = id.(string)
	}
	return
}

// Wrapper for item.delete: https://www.zabbix.com/documentation/2.2/manual/appendix/api/item/delete
// Cleans ItemId in all items elements if call succeed.
func (api *API) ItemsDelete(items Items) (err error) {
	ids := make([]string, len(items))
	for i, item := range items {
		ids[i] = item.ItemId
	}

	err = api.ItemsDeleteByIds(ids)
	if err == nil {
		for i := range items {
			items[i].ItemId = ""
		}
	}
	return
}

// Wrapper for item.delete: https://www.zabbix.com/documentation/2.2/manual/appendix/api/item/delete
func (api *API) ItemsDeleteByIds(ids []string) (err error) {
	response, err := api.CallWithError("item.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	itemids1, ok := result["itemids"].([]interface{})
	l := len(itemids1)
	if !ok {
		// some versions actually return map there
		itemids2 := result["itemids"].(map[string]interface{})
		l = len(itemids2)
	}
	if len(ids) != l {
		err = &ExpectedMore{len(ids), l}
	}
	return
}
