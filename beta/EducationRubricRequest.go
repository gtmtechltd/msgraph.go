// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EducationRubricRequestBuilder is request builder for EducationRubric
type EducationRubricRequestBuilder struct{ BaseRequestBuilder }

// Request returns EducationRubricRequest
func (b *EducationRubricRequestBuilder) Request() *EducationRubricRequest {
	return &EducationRubricRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EducationRubricRequest is request for EducationRubric
type EducationRubricRequest struct{ BaseRequest }

// Do performs HTTP request for EducationRubric
func (r *EducationRubricRequest) Do(method, path string, reqObj interface{}) (resObj *EducationRubric, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for EducationRubric
func (r *EducationRubricRequest) Get() (*EducationRubric, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for EducationRubric
func (r *EducationRubricRequest) Update(reqObj *EducationRubric) (*EducationRubric, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for EducationRubric
func (r *EducationRubricRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}