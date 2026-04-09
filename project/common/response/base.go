package response

type BaseResp[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
	// Message string `json:"message"`
	Error *ErrorInfo `json:"error,omitempty"`
	Meta  *Meta      `json:"meta,omitempty"`
}
