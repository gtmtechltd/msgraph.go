// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// AppCatalogsRequestBuilder is request builder for AppCatalogs
type AppCatalogsRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppCatalogsRequest
func (b *AppCatalogsRequestBuilder) Request() *AppCatalogsRequest {
	return &AppCatalogsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AppCatalogsRequest is request for AppCatalogs
type AppCatalogsRequest struct{ BaseRequest }

// Get performs GET request for AppCatalogs
func (r *AppCatalogsRequest) Get(ctx context.Context) (resObj *AppCatalogs, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AppCatalogs
func (r *AppCatalogsRequest) Update(ctx context.Context, reqObj *AppCatalogs) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AppCatalogs
func (r *AppCatalogsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
