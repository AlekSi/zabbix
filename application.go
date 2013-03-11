package zabbix

import (
	"github.com/AlekSi/reflector"
)

// https://www.zabbix.com/documentation/2.0/manual/appendix/api/application/definitions
type Application struct {
	ApplicationId string `json:"applicationid,omitempty"`
	HostId        string `json:"hostid"`
	Name          string `json:"name"`
	TemplateId    string `json:"templateid,omitempty"`
}

type Applications []Application

// Wrapper for application.get: https://www.zabbix.com/documentation/2.0/manual/appendix/api/application/get
func (api *API) ApplicationsGet(params Params) (res Applications, err error) {
	if _, present := params["output"]; !present {
		params["output"] = "extend"
	}
	response, err := api.CallWithError("application.get", params)
	if err != nil {
		return
	}

	reflector.MapsToStructs2(response.Result.([]interface{}), &res, reflector.Strconv, "json")
	return
}

// Gets application by Id.
func (api *API) ApplicationGetById(id string) (res *Application, err error) {
	apps, err := api.ApplicationsGet(Params{"applicationids": id})
	if err != nil {
		return
	}

	if len(apps) == 1 {
		res = &apps[0]
	} else {
		e := ExpectedOneResult(len(apps))
		err = &e
	}
	return
}

// Wrapper for application.create: https://www.zabbix.com/documentation/2.0/manual/appendix/api/application/create
func (api *API) ApplicationsCreate(apps Applications) (err error) {
	response, err := api.CallWithError("application.create", apps)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	applicationids := result["applicationids"].([]interface{})
	for i, id := range applicationids {
		apps[i].ApplicationId = id.(string)
	}
	return
}

// Wrapper for application.delete: https://www.zabbix.com/documentation/2.0/manual/appendix/api/application/delete
func (api *API) ApplicationsDelete(apps Applications) (err error) {
	ids := make([]string, len(apps))
	for i, app := range apps {
		ids[i] = app.ApplicationId
	}
	return api.ApplicationsDeleteIds(ids)
}

// Wrapper for application.delete: https://www.zabbix.com/documentation/2.0/manual/appendix/api/application/delete
func (api *API) ApplicationsDeleteIds(ids []string) (err error) {
	response, err := api.CallWithError("application.delete", ids)
	if err != nil {
		return
	}

	result := response.Result.(map[string]interface{})
	applicationids := result["applicationids"].([]interface{})
	if len(ids) != len(applicationids) {
		err = &ExpectedMore{len(ids), len(applicationids)}
	}
	return
}
