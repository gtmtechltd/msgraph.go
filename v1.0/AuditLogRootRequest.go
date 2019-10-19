// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// AuditLogRootRequestBuilder is request builder for AuditLogRoot
type AuditLogRootRequestBuilder struct{ BaseRequestBuilder }

// Request returns AuditLogRootRequest
func (b *AuditLogRootRequestBuilder) Request() *AuditLogRootRequest {
	return &AuditLogRootRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AuditLogRootRequest is request for AuditLogRoot
type AuditLogRootRequest struct{ BaseRequest }

// Do performs HTTP request for AuditLogRoot
func (r *AuditLogRootRequest) Do(method, path string, reqObj interface{}) (resObj *AuditLogRoot, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for AuditLogRoot
func (r *AuditLogRootRequest) Get() (*AuditLogRoot, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for AuditLogRoot
func (r *AuditLogRootRequest) Update(reqObj *AuditLogRoot) (*AuditLogRoot, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for AuditLogRoot
func (r *AuditLogRootRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// DirectoryAudits returns request builder for DirectoryAudit collection
func (b *AuditLogRootRequestBuilder) DirectoryAudits() *AuditLogRootDirectoryAuditsCollectionRequestBuilder {
	bb := &AuditLogRootDirectoryAuditsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/directoryAudits"
	return bb
}

// AuditLogRootDirectoryAuditsCollectionRequestBuilder is request builder for DirectoryAudit collection
type AuditLogRootDirectoryAuditsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryAudit collection
func (b *AuditLogRootDirectoryAuditsCollectionRequestBuilder) Request() *AuditLogRootDirectoryAuditsCollectionRequest {
	return &AuditLogRootDirectoryAuditsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryAudit item
func (b *AuditLogRootDirectoryAuditsCollectionRequestBuilder) ID(id string) *DirectoryAuditRequestBuilder {
	bb := &DirectoryAuditRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuditLogRootDirectoryAuditsCollectionRequest is request for DirectoryAudit collection
type AuditLogRootDirectoryAuditsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for DirectoryAudit collection
func (r *AuditLogRootDirectoryAuditsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *DirectoryAudit, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for DirectoryAudit collection
func (r *AuditLogRootDirectoryAuditsCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryAudit, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryAudit
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DirectoryAudit
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

// Get performs GET request for DirectoryAudit collection
func (r *AuditLogRootDirectoryAuditsCollectionRequest) Get() ([]DirectoryAudit, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryAudit collection
func (r *AuditLogRootDirectoryAuditsCollectionRequest) Add(reqObj *DirectoryAudit) (*DirectoryAudit, error) {
	return r.Do("POST", "", reqObj)
}

// RestrictedSignIns returns request builder for RestrictedSignIn collection
func (b *AuditLogRootRequestBuilder) RestrictedSignIns() *AuditLogRootRestrictedSignInsCollectionRequestBuilder {
	bb := &AuditLogRootRestrictedSignInsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/restrictedSignIns"
	return bb
}

// AuditLogRootRestrictedSignInsCollectionRequestBuilder is request builder for RestrictedSignIn collection
type AuditLogRootRestrictedSignInsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for RestrictedSignIn collection
func (b *AuditLogRootRestrictedSignInsCollectionRequestBuilder) Request() *AuditLogRootRestrictedSignInsCollectionRequest {
	return &AuditLogRootRestrictedSignInsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for RestrictedSignIn item
func (b *AuditLogRootRestrictedSignInsCollectionRequestBuilder) ID(id string) *RestrictedSignInRequestBuilder {
	bb := &RestrictedSignInRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuditLogRootRestrictedSignInsCollectionRequest is request for RestrictedSignIn collection
type AuditLogRootRestrictedSignInsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for RestrictedSignIn collection
func (r *AuditLogRootRestrictedSignInsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *RestrictedSignIn, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for RestrictedSignIn collection
func (r *AuditLogRootRestrictedSignInsCollectionRequest) Paging(method, path string, obj interface{}) ([]RestrictedSignIn, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []RestrictedSignIn
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []RestrictedSignIn
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

// Get performs GET request for RestrictedSignIn collection
func (r *AuditLogRootRestrictedSignInsCollectionRequest) Get() ([]RestrictedSignIn, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for RestrictedSignIn collection
func (r *AuditLogRootRestrictedSignInsCollectionRequest) Add(reqObj *RestrictedSignIn) (*RestrictedSignIn, error) {
	return r.Do("POST", "", reqObj)
}

// SignIns returns request builder for SignIn collection
func (b *AuditLogRootRequestBuilder) SignIns() *AuditLogRootSignInsCollectionRequestBuilder {
	bb := &AuditLogRootSignInsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/signIns"
	return bb
}

// AuditLogRootSignInsCollectionRequestBuilder is request builder for SignIn collection
type AuditLogRootSignInsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SignIn collection
func (b *AuditLogRootSignInsCollectionRequestBuilder) Request() *AuditLogRootSignInsCollectionRequest {
	return &AuditLogRootSignInsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SignIn item
func (b *AuditLogRootSignInsCollectionRequestBuilder) ID(id string) *SignInRequestBuilder {
	bb := &SignInRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AuditLogRootSignInsCollectionRequest is request for SignIn collection
type AuditLogRootSignInsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for SignIn collection
func (r *AuditLogRootSignInsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *SignIn, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for SignIn collection
func (r *AuditLogRootSignInsCollectionRequest) Paging(method, path string, obj interface{}) ([]SignIn, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SignIn
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []SignIn
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

// Get performs GET request for SignIn collection
func (r *AuditLogRootSignInsCollectionRequest) Get() ([]SignIn, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SignIn collection
func (r *AuditLogRootSignInsCollectionRequest) Add(reqObj *SignIn) (*SignIn, error) {
	return r.Do("POST", "", reqObj)
}