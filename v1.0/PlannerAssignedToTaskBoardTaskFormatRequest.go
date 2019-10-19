// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PlannerAssignedToTaskBoardTaskFormatRequestBuilder is request builder for PlannerAssignedToTaskBoardTaskFormat
type PlannerAssignedToTaskBoardTaskFormatRequestBuilder struct{ BaseRequestBuilder }

// Request returns PlannerAssignedToTaskBoardTaskFormatRequest
func (b *PlannerAssignedToTaskBoardTaskFormatRequestBuilder) Request() *PlannerAssignedToTaskBoardTaskFormatRequest {
	return &PlannerAssignedToTaskBoardTaskFormatRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PlannerAssignedToTaskBoardTaskFormatRequest is request for PlannerAssignedToTaskBoardTaskFormat
type PlannerAssignedToTaskBoardTaskFormatRequest struct{ BaseRequest }

// Do performs HTTP request for PlannerAssignedToTaskBoardTaskFormat
func (r *PlannerAssignedToTaskBoardTaskFormatRequest) Do(method, path string, reqObj interface{}) (resObj *PlannerAssignedToTaskBoardTaskFormat, err error) {
	err = r.JSONRequest(method, path, reqObj, &resObj)
	return
}

// Get performs GET request for PlannerAssignedToTaskBoardTaskFormat
func (r *PlannerAssignedToTaskBoardTaskFormatRequest) Get() (*PlannerAssignedToTaskBoardTaskFormat, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Do("GET", query, nil)
}

// Update performs PATCH request for PlannerAssignedToTaskBoardTaskFormat
func (r *PlannerAssignedToTaskBoardTaskFormatRequest) Update(reqObj *PlannerAssignedToTaskBoardTaskFormat) (*PlannerAssignedToTaskBoardTaskFormat, error) {
	return r.Do("PATCH", "", reqObj)
}

// Delete performs DELETE request for PlannerAssignedToTaskBoardTaskFormat
func (r *PlannerAssignedToTaskBoardTaskFormatRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}