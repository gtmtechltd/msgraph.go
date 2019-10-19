// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// BaseItemRequestBuilder is request builder for BaseItem
type BaseItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns BaseItemRequest
func (b *BaseItemRequestBuilder) Request() *BaseItemRequest {
	return &BaseItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BaseItemRequest is request for BaseItem
type BaseItemRequest struct{ BaseRequest }

// Do performs HTTP request for BaseItem
func (r *BaseItemRequest) Do(method, path string, reqObj interface{}) (resObj *BaseItem, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for BaseItem
func (r *BaseItemRequest) Get() (*BaseItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for BaseItem
func (r *BaseItemRequest) Update(reqObj *BaseItem) (*BaseItem, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for BaseItem
func (r *BaseItemRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// CreatedByUser is navigation property
func (b *BaseItemRequestBuilder) CreatedByUser() *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/createdByUser"
	return bb
}

// LastModifiedByUser is navigation property
func (b *BaseItemRequestBuilder) LastModifiedByUser() *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/lastModifiedByUser"
	return bb
}