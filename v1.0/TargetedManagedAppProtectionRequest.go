// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TargetedManagedAppProtectionRequestBuilder is request builder for TargetedManagedAppProtection
type TargetedManagedAppProtectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns TargetedManagedAppProtectionRequest
func (b *TargetedManagedAppProtectionRequestBuilder) Request() *TargetedManagedAppProtectionRequest {
	return &TargetedManagedAppProtectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TargetedManagedAppProtectionRequest is request for TargetedManagedAppProtection
type TargetedManagedAppProtectionRequest struct{ BaseRequest }

// Do performs HTTP request for TargetedManagedAppProtection
func (r *TargetedManagedAppProtectionRequest) Do(method, path string, reqObj interface{}) (resObj *TargetedManagedAppProtection, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for TargetedManagedAppProtection
func (r *TargetedManagedAppProtectionRequest) Get() (*TargetedManagedAppProtection, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for TargetedManagedAppProtection
func (r *TargetedManagedAppProtectionRequest) Update(reqObj *TargetedManagedAppProtection) (*TargetedManagedAppProtection, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for TargetedManagedAppProtection
func (r *TargetedManagedAppProtectionRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Assignments returns request builder for TargetedManagedAppPolicyAssignment collection
func (b *TargetedManagedAppProtectionRequestBuilder) Assignments() *TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder {
	bb := &TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder is request builder for TargetedManagedAppPolicyAssignment collection
type TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TargetedManagedAppPolicyAssignment collection
func (b *TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder) Request() *TargetedManagedAppProtectionAssignmentsCollectionRequest {
	return &TargetedManagedAppProtectionAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TargetedManagedAppPolicyAssignment item
func (b *TargetedManagedAppProtectionAssignmentsCollectionRequestBuilder) ID(id string) *TargetedManagedAppPolicyAssignmentRequestBuilder {
	bb := &TargetedManagedAppPolicyAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TargetedManagedAppProtectionAssignmentsCollectionRequest is request for TargetedManagedAppPolicyAssignment collection
type TargetedManagedAppProtectionAssignmentsCollectionRequest struct{ BaseRequest }

// Do performs HTTP request for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppProtectionAssignmentsCollectionRequest) Do(method, path string, reqObj interface{}) (resObj *TargetedManagedAppPolicyAssignment, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Paging perfoms paging operation for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppProtectionAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]TargetedManagedAppPolicyAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TargetedManagedAppPolicyAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []TargetedManagedAppPolicyAssignment
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

// Get performs GET request for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppProtectionAssignmentsCollectionRequest) Get() ([]TargetedManagedAppPolicyAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for TargetedManagedAppPolicyAssignment collection
func (r *TargetedManagedAppProtectionAssignmentsCollectionRequest) Add(reqObj *TargetedManagedAppPolicyAssignment) (*TargetedManagedAppPolicyAssignment, error) {
	return r.Do("POST", "", reqObj)
}