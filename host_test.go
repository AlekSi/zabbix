package zabbix_test

import (
	. "."
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func CreateHost(group *HostGroup, t *testing.T) *Host {
	name := fmt.Sprintf("%s-%d", getHost(), rand.Int())
	iface := HostInterface{DNS: name, Port: "42", Type: Agent, UseIP: 0, Main: 1}
	hosts := Hosts{{
		Host:       name,
		Name:       "Name for " + name,
		GroupIds:   HostGroupIds{{group.GroupId}},
		Interfaces: HostInterfaces{iface},
	}}

	err := getAPI(t).HostsCreate(hosts)
	if err != nil {
		t.Fatal(err)
	}
	return &hosts[0]
}

func DeleteHost(host *Host, t *testing.T) {
	err := getAPI(t).HostsDelete(Hosts{*host})
	if err != nil {
		t.Fatal(err)
	}
}

func TestHosts(t *testing.T) {
	api := getAPI(t)

	group := CreateHostGroup(t)
	defer DeleteHostGroup(group, t)

	hosts, err := api.HostsGetByHostGroups(HostGroups{*group})
	if err != nil {
		t.Fatal(err)
	}
	if len(hosts) != 0 {
		t.Errorf("Bad hosts: %#v", hosts)
	}

	host := CreateHost(group, t)
	if host.HostId == "" || host.Host == "" {
		t.Errorf("Something is empty: %#v", host)
	}
	host.GroupIds = nil
	host.Interfaces = nil

	host2, err := api.HostGetByHost(host.Host)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(host, host2) {
		t.Errorf("Hosts are not equal:\n%#v\n%#v", host, host2)
	}

	host2, err = api.HostGetById(host.HostId)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(host, host2) {
		t.Errorf("Hosts are not equal:\n%#v\n%#v", host, host2)
	}

	hosts, err = api.HostsGetByHostGroups(HostGroups{*group})
	if err != nil {
		t.Fatal(err)
	}
	if len(hosts) != 1 {
		t.Errorf("Bad hosts: %#v", hosts)
	}

	DeleteHost(host, t)

	hosts, err = api.HostsGetByHostGroups(HostGroups{*group})
	if err != nil {
		t.Fatal(err)
	}
	if len(hosts) != 0 {
		t.Errorf("Bad hosts: %#v", hosts)
	}
}
