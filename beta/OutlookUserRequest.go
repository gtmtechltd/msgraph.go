// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// OutlookUserRequestBuilder is request builder for OutlookUser
type OutlookUserRequestBuilder struct{ BaseRequestBuilder }

// Request returns OutlookUserRequest
func (b *OutlookUserRequestBuilder) Request() *OutlookUserRequest {
	return &OutlookUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OutlookUserRequest is request for OutlookUser
type OutlookUserRequest struct{ BaseRequest }

// Do performs HTTP request for OutlookUser
func (r *OutlookUserRequest) Do(method, path string, reqObj interface{}) (resObj *OutlookUser, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for OutlookUser
func (r *OutlookUserRequest) Get() (*OutlookUser, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for OutlookUser
func (r *OutlookUserRequest) Update(reqObj *OutlookUser) (*OutlookUser, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for OutlookUser
func (r *OutlookUserRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// MasterCategories returns request builder for OutlookCategory collection
func (b *OutlookUserRequestBuilder) MasterCategories() *OutlookUserMasterCategoriesCollectionRequestBuilder {
	bb := &OutlookUserMasterCategoriesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/masterCategories"
	return bb
}

// OutlookUserMasterCategoriesCollectionRequestBuilder is request builder for OutlookCategory collection
type OutlookUserMasterCategoriesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OutlookCategory collection
func (b *OutlookUserMasterCategoriesCollectionRequestBuilder) Request() *OutlookUserMasterCategoriesCollectionRequest {
	return &OutlookUserMasterCategoriesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OutlookCategory item
func (b *OutlookUserMasterCategoriesCollectionRequestBuilder) ID(id string) *OutlookCategoryRequestBuilder {
	bb := &OutlookCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OutlookUserMasterCategoriesCollectionRequest is request for OutlookCategory collection
type OutlookUserMasterCategoriesCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for OutlookCategory collection
func (r *OutlookUserMasterCategoriesCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *OutlookCategory, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for OutlookCategory collection
func (r *OutlookUserMasterCategoriesCollectionRequest) Paging(method, path string, obj interface{}) ([]OutlookCategory, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OutlookCategory
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []OutlookCategory
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for OutlookCategory collection
func (r *OutlookUserMasterCategoriesCollectionRequest) Get() ([]OutlookCategory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for OutlookCategory collection
func (r *OutlookUserMasterCategoriesCollectionRequest) Add(reqObj *OutlookCategory) (*OutlookCategory, error) {
	return r.Do("POST", "", reqObj)
}

// TaskFolders returns request builder for OutlookTaskFolder collection
func (b *OutlookUserRequestBuilder) TaskFolders() *OutlookUserTaskFoldersCollectionRequestBuilder {
	bb := &OutlookUserTaskFoldersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/taskFolders"
	return bb
}

// OutlookUserTaskFoldersCollectionRequestBuilder is request builder for OutlookTaskFolder collection
type OutlookUserTaskFoldersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OutlookTaskFolder collection
func (b *OutlookUserTaskFoldersCollectionRequestBuilder) Request() *OutlookUserTaskFoldersCollectionRequest {
	return &OutlookUserTaskFoldersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OutlookTaskFolder item
func (b *OutlookUserTaskFoldersCollectionRequestBuilder) ID(id string) *OutlookTaskFolderRequestBuilder {
	bb := &OutlookTaskFolderRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OutlookUserTaskFoldersCollectionRequest is request for OutlookTaskFolder collection
type OutlookUserTaskFoldersCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for OutlookTaskFolder collection
func (r *OutlookUserTaskFoldersCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *OutlookTaskFolder, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for OutlookTaskFolder collection
func (r *OutlookUserTaskFoldersCollectionRequest) Paging(method, path string, obj interface{}) ([]OutlookTaskFolder, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OutlookTaskFolder
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []OutlookTaskFolder
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for OutlookTaskFolder collection
func (r *OutlookUserTaskFoldersCollectionRequest) Get() ([]OutlookTaskFolder, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for OutlookTaskFolder collection
func (r *OutlookUserTaskFoldersCollectionRequest) Add(reqObj *OutlookTaskFolder) (*OutlookTaskFolder, error) {
	return r.Do("POST", "", reqObj)
}

// TaskGroups returns request builder for OutlookTaskGroup collection
func (b *OutlookUserRequestBuilder) TaskGroups() *OutlookUserTaskGroupsCollectionRequestBuilder {
	bb := &OutlookUserTaskGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/taskGroups"
	return bb
}

// OutlookUserTaskGroupsCollectionRequestBuilder is request builder for OutlookTaskGroup collection
type OutlookUserTaskGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OutlookTaskGroup collection
func (b *OutlookUserTaskGroupsCollectionRequestBuilder) Request() *OutlookUserTaskGroupsCollectionRequest {
	return &OutlookUserTaskGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OutlookTaskGroup item
func (b *OutlookUserTaskGroupsCollectionRequestBuilder) ID(id string) *OutlookTaskGroupRequestBuilder {
	bb := &OutlookTaskGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OutlookUserTaskGroupsCollectionRequest is request for OutlookTaskGroup collection
type OutlookUserTaskGroupsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for OutlookTaskGroup collection
func (r *OutlookUserTaskGroupsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *OutlookTaskGroup, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for OutlookTaskGroup collection
func (r *OutlookUserTaskGroupsCollectionRequest) Paging(method, path string, obj interface{}) ([]OutlookTaskGroup, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OutlookTaskGroup
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []OutlookTaskGroup
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for OutlookTaskGroup collection
func (r *OutlookUserTaskGroupsCollectionRequest) Get() ([]OutlookTaskGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for OutlookTaskGroup collection
func (r *OutlookUserTaskGroupsCollectionRequest) Add(reqObj *OutlookTaskGroup) (*OutlookTaskGroup, error) {
	return r.Do("POST", "", reqObj)
}

// Tasks returns request builder for OutlookTask collection
func (b *OutlookUserRequestBuilder) Tasks() *OutlookUserTasksCollectionRequestBuilder {
	bb := &OutlookUserTasksCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/tasks"
	return bb
}

// OutlookUserTasksCollectionRequestBuilder is request builder for OutlookTask collection
type OutlookUserTasksCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OutlookTask collection
func (b *OutlookUserTasksCollectionRequestBuilder) Request() *OutlookUserTasksCollectionRequest {
	return &OutlookUserTasksCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OutlookTask item
func (b *OutlookUserTasksCollectionRequestBuilder) ID(id string) *OutlookTaskRequestBuilder {
	bb := &OutlookTaskRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OutlookUserTasksCollectionRequest is request for OutlookTask collection
type OutlookUserTasksCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for OutlookTask collection
func (r *OutlookUserTasksCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *OutlookTask, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for OutlookTask collection
func (r *OutlookUserTasksCollectionRequest) Paging(method, path string, obj interface{}) ([]OutlookTask, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OutlookTask
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []OutlookTask
		)
		err := json.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for OutlookTask collection
func (r *OutlookUserTasksCollectionRequest) Get() ([]OutlookTask, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for OutlookTask collection
func (r *OutlookUserTasksCollectionRequest) Add(reqObj *OutlookTask) (*OutlookTask, error) {
	return r.Do("POST", "", reqObj)
}