// Error types For handeling errors with fiber
package apierror

// Format for returning errors from API calls
// If an Handler returns an error it will be wrapped into this struct.
// The Message field contains the returned error message.
type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
