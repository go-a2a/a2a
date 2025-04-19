// Copyright 2025 The Go A2A Authors
// SPDX-License-Identifier: Apache-2.0

package a2a

import (
	"encoding/json"

	"github.com/bytedance/sonic"
)

// A2A RPC method names.
const (
	// MethodSend is the method name for sending a task.
	MethodSend = "tasks/send"
	// MethodGet is the method name for getting a task.
	MethodGet = "tasks/get"
	// MethodCancel is the method name for canceling a task.
	MethodCancel = "tasks/cancel"
	// MethodPushNotificationSet is the method name for setting push notification configuration.
	MethodPushNotificationSet = "tasks/pushNotification/set"
	// MethodPushNotificationGet is the method name for getting push notification configuration.
	MethodPushNotificationGet = "tasks/pushNotification/get"
	// MethodSendSubscribe is the method name for sending a task and subscribing to updates.
	MethodSendSubscribe = "tasks/sendSubscribe"
	// MethodResubscribe is the method name for resubscribing to task updates.
	MethodResubscribe = "tasks/resubscribe"
)

// // ID represents the unique identifier for JSON-RPC messages.
// type ID any
//
// // NewID returns a new request ID.
// func NewID[T ~string | ~float64](v T) ID {
// 	return v
// }
//
// // MarshalID implements json.Marshaler.
// func MarshalID(id ID) ([]byte, error) {
// 	return sonic.ConfigDefault.Marshal(id)
// }
//
// // UnmarshalJSON implements json.Unmarshaler.
// func UnmarshalID(id ID, data []byte) error {
// 	return sonic.ConfigDefault.Unmarshal(data, &id)
// }

// ID represents the unique identifier for JSON-RPC messages.
type id struct {
	any
}

// NewID returns a new request ID.
func NewID[T string | int32](v T) id {
	return id{any: v}
}

// MarshalJSON implements json.Marshaler.
func (i id) MarshalJSON() ([]byte, error) {
	return sonic.ConfigDefault.Marshal(&i)
}

// UnmarshalJSON implements json.Unmarshaler.
func (i *id) UnmarshalJSON(data []byte) error {
	*i = id{}
	return sonic.ConfigDefault.Unmarshal(data, &i)
}

type ID = id

// // String returns a string representation of the ID.
// func (id ID) String() string {
// 	switch id := id.v.(type) {
// 	case string:
// 		return id
// 	case int32:
// 		return strconv.FormatInt(int64(id), 10)
// 	case float64:
// 		return strconv.FormatFloat(id, 'f', -1, 64)
// 	default:
// 		panic("unreachable")
// 	}
// }

// // Format writes the ID to the formatter.
// //
// // If the rune is q the representation is non ambiguous,
// // string forms are quoted, number forms are preceded by a #.
// func (id ID[T]) Format(f fmt.State, r rune) {
// 	numF, strF := `%d`, `%s`
// 	if r == 'q' {
// 		numF, strF = `#%d`, `%q`
// 	}
//
// 	switch {
// 	case id.name != "":
// 		fmt.Fprintf(f, strF, id.name)
// 	default:
// 		fmt.Fprintf(f, numF, id.number)
// 	}
// }

// JSONRPCMessage is the base structure for all JSON-RPC 2.0 messages.
type JSONRPCMessage struct {
	// JSONRPC version, always "2.0".
	JSONRPC string `json:"jsonrpc"`

	// ID is a unique identifier for the request/response correlation.
	ID `json:"id,omitzero"` // string, number, or null
}

// NewJSONRPCMessage creates a new [JSONRPCMessage] with the given id.
func NewJSONRPCMessage(id ID) JSONRPCMessage {
	return JSONRPCMessage{
		JSONRPC: "2.0",
		ID:      id,
	}
}

// JSONRPCRequest represents a JSON-RPC 2.0 request.
type JSONRPCRequest struct {
	JSONRPCMessage

	// Method identifies the operation to perform.
	Method string `json:"method"`

	// Params contains parameters for the method.
	Params json.RawMessage `json:"params,omitempty"`
}

// JSONRPCResponse represents a JSON-RPC 2.0 response.
type JSONRPCResponse struct {
	JSONRPCMessage

	// Result contains the successful result data (can be null).
	// Mutually exclusive with Error.
	Result any `json:"result,omitempty"`

	// Error contains an error object if the request failed.
	// Mutually exclusive with Result.
	Error *JSONRPCError `json:"error,omitempty"`
}

// Standard JSON-RPC 2.0 error codes.
const (
	// JSONParseErrorCode indicates invalid JSON payload.
	JSONParseErrorCode = -32700
	// InvalidRequestErrorCode indicates request payload validation error.
	InvalidRequestErrorCode = -32600
	// MethodNotFoundErrorCode indicates the method does not exist.
	MethodNotFoundErrorCode = -32601
	// InvalidParamsErrorCode indicates invalid method parameters.
	InvalidParamsErrorCode = -32602
	// InternalErrorCode indicates an internal server error.
	InternalErrorCode = -32603
)

// A2A specific error codes.
const (
	// TaskNotFoundErrorCode indicates the specified task ID was not found.
	TaskNotFoundErrorCode = -32001
	// TaskNotCancelableErrorCode indicates the task is in a final state and cannot be canceled.
	TaskNotCancelableErrorCode = -32002
	// PushNotificationNotSupportedErrorCode indicates the agent does not support push notifications.
	PushNotificationNotSupportedErrorCode = -32003
	// UnsupportedOperationErrorCode indicates the requested operation is not supported.
	UnsupportedOperationErrorCode = -32004
	// ContentTypeNotSupportedErrorCode indicates a mismatch in supported content types.
	ContentTypeNotSupportedErrorCode = -32005
)

// JSONRPCError represents a JSON-RPC 2.0 error.
type JSONRPCError struct {
	// Code is the error code.
	Code int `json:"code"`
	// Message is a short description of the error.
	Message string `json:"message"`
	// Data contains optional additional error details.
	Data any `json:"data,omitempty"`
}

// NewJSONParseError creates a new JSONParseError.
func NewJSONParseError() *JSONRPCError {
	return &JSONRPCError{
		Code:    JSONParseErrorCode,
		Message: "Invalid JSON payload",
	}
}

// NewInvalidRequestError creates a new InvalidRequestError.
func NewInvalidRequestError() *JSONRPCError {
	return &JSONRPCError{
		Code:    InvalidRequestErrorCode,
		Message: "Request payload validation error",
	}
}

// NewMethodNotFoundError creates a new MethodNotFoundError.
func NewMethodNotFoundError() *JSONRPCError {
	return &JSONRPCError{
		Code:    MethodNotFoundErrorCode,
		Message: "Method not found",
	}
}

// NewInvalidParamsError creates a new InvalidParamsError.
func NewInvalidParamsError() *JSONRPCError {
	return &JSONRPCError{
		Code:    InvalidParamsErrorCode,
		Message: "Invalid parameters",
	}
}

// NewInternalError creates a new InternalError.
func NewInternalError() *JSONRPCError {
	return &JSONRPCError{
		Code:    InternalErrorCode,
		Message: "Internal error",
	}
}

// NewTaskNotFoundError creates a new TaskNotFoundError.
func NewTaskNotFoundError() *JSONRPCError {
	return &JSONRPCError{
		Code:    TaskNotFoundErrorCode,
		Message: "Task not found",
	}
}

// NewTaskNotCancelableError creates a new TaskNotCancelableError.
func NewTaskNotCancelableError() *JSONRPCError {
	return &JSONRPCError{
		Code:    TaskNotCancelableErrorCode,
		Message: "Task cannot be canceled",
	}
}

// NewPushNotificationNotSupportedError creates a new PushNotificationNotSupportedError.
func NewPushNotificationNotSupportedError() *JSONRPCError {
	return &JSONRPCError{
		Code:    PushNotificationNotSupportedErrorCode,
		Message: "Push Notification is not supported",
	}
}

// NewUnsupportedOperationError creates a new UnsupportedOperationError.
func NewUnsupportedOperationError() *JSONRPCError {
	return &JSONRPCError{
		Code:    UnsupportedOperationErrorCode,
		Message: "This operation is not supported",
	}
}

// NewContentTypeNotSupportedError creates a new ContentTypeNotSupportedError.
func NewContentTypeNotSupportedError() *JSONRPCError {
	return &JSONRPCError{
		Code:    ContentTypeNotSupportedErrorCode,
		Message: "Content type not supported",
	}
}
