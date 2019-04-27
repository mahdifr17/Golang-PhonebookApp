package exception

import (
	"fmt"

	"github.com/mahdifr17/phonebook-v2/structs"
)

// NotFound is a type to handle error not found exception
type NotFound int

func (e NotFound) Error() string {
	return fmt.Sprintf("Record with id %v not found", int(e))
}

// SerializeError is a function to serialize every error type to BaseResponse
func SerializeError(e error) structs.BaseResponse {
	switch e.(type) {
	case NotFound:
		return structs.BaseResponse{Code: 404, Msg: e.Error()}
	default:
		return structs.BaseResponse{Code: 500, Msg: e.Error()}
	}
}
