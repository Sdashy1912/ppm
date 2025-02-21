package basedto

const (
	// StatusCodeOK indicates request has successfully handled
	StatusCodeOK int = 200
	// StatusCodeNotFound indicates the resource is not found
	StatusCodeNotFound int = 404
	// StatusCodeUnprocessableEntity indicates entity could not be process due to some reasons
	StatusCodeUnprocessableEntity int = 422
	// StatusCodeInternalServerError idicates an unexpected error happens
	StatusCodeInternalServerError int = 500
)

// ResInfo base informations for a response
type ResInfo struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
}

// SetStatusOK set status ok
func (info *ResInfo) SetStatusOK(message interface{}) {
	info.StatusCode = StatusCodeOK
	info.Message = message
}

// SetStatusNotFound set status not found
func (info *ResInfo) SetStatusNotFound() {
	info.StatusCode = StatusCodeNotFound
	info.Message = "Not found"
}

// SetStatusUnprocessableEntity set status unprocessable entity
func (info *ResInfo) SetStatusUnprocessableEntity(message interface{}) {
	info.StatusCode = StatusCodeUnprocessableEntity
	info.Message = message
}

// SetStatusInternalServerError set status internal server error
func (info *ResInfo) SetStatusInternalServerError() {
	info.StatusCode = StatusCodeInternalServerError
	info.Message = "Internal Server Error"
}
