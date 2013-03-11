package zabbix_test

import (
	. "."
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func CreateHostGroup(t *testing.T) *HostGroup {
	hostGroups := HostGroups{{Name: fmt.Sprintf("zabbix-testing-%d", rand.Int())}}
	err := getAPI(t).HostGroupsCreate(hostGroups)
	if err != nil {
		t.Fatal(err)
	}
	return &hostGroups[0]
}

func DeleteHostGroup(hostGroup *HostGroup, t *testing.T) {
	err := getAPI(t).HostGroupsDelete(HostGroups{*hostGroup})
	if err != nil {
		t.Fatal(err)
	}
}

func TestHostGroups(t *testing.T) {
	api := getAPI(t)

	groups, err := api.HostGroupsGet(Params{})
	if err != nil {
		t.Fatal(err)
	}

	hostGroup := CreateHostGroup(t)
	if hostGroup.GroupId == "" || hostGroup.Name == "" {
		t.Errorf("Something is empty: %#v", hostGroup)
	}

	hostGroup2, err := api.HostGroupGetById(hostGroup.GroupId)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(hostGroup, hostGroup2) {
		t.Errorf("Error getting group.\nOld group: %#v\nNew group: %#v", hostGroup, hostGroup2)
	}

	groups2, err := api.HostGroupsGet(Params{})
	if err != nil {
		t.Fatal(err)
	}
	if len(groups2) != len(groups)+1 {
		t.Errorf("Error creating group.\nOld groups: %#v\nNew groups: %#v", groups, groups2)
	}

	DeleteHostGroup(hostGroup, t)

	groups2, err = api.HostGroupsGet(Params{})
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(groups, groups2) {
		t.Errorf("Error deleting group.\nOld groups: %#v\nNew groups: %#v", groups, groups2)
	}
}
