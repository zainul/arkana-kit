package contract

import (
	"github.com/zainul/arkana-kit/internal/pkg/error/deliveryerror"
)

// Response is generic response for HTTP response
type Response struct {
	Data  interface{}          `json:"data"`
	Error *deliveryerror.Error `json:"error"`
}

// Success ...
type Success struct {
	Success bool `json:"success"`
}
