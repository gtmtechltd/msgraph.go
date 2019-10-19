// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PlaceRequestBuilder is request builder for Place
type PlaceRequestBuilder struct{ BaseRequestBuilder }

// Request returns PlaceRequest
func (b *PlaceRequestBuilder) Request() *PlaceRequest {
	return &PlaceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PlaceRequest is request for Place
type PlaceRequest struct{ BaseRequest }

// Do performs HTTP request for Place
func (r *PlaceRequest) Do(method, path string, reqObj interface{}) (resObj *Place, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for Place
func (r *PlaceRequest) Get() (*Place, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for Place
func (r *PlaceRequest) Update(reqObj *Place) (*Place, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for Place
func (r *PlaceRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}