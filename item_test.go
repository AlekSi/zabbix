package zabbix_test

import (
	. "."
	"testing"
)

func CreateItem(app *Application, t *testing.T) *Item {
	items := Items{{
		HostId:         app.HostId,
		Key:            "key.lala.laa",
		Name:           "name for key",
		Type:           ZabbixTrapper,
		ApplicationIds: []string{app.ApplicationId},
	}}
	err := getAPI(t).ItemsCreate(items)
	if err != nil {
		t.Fatal(err)
	}
	return &items[0]
}

func DeleteItem(item *Item, t *testing.T) {
	err := getAPI(t).ItemsDelete(Items{*item})
	if err != nil {
		t.Fatal(err)
	}
}

func TestItems(t *testing.T) {
	api := getAPI(t)

	group := CreateHostGroup(t)
	defer DeleteHostGroup(group, t)

	host := CreateHost(group, t)
	defer DeleteHost(host, t)

	app := CreateApplication(host, t)
	defer DeleteApplication(app, t)

	items, err := api.ItemsGetByApplicationId(app.ApplicationId)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 0 {
		t.Fatal("Found items")
	}

	item := CreateItem(app, t)
	DeleteItem(item, t)
}
