// Copyright 2025 The Go A2A Authors
// SPDX-License-Identifier: Apache-2.0

// Package a2a provides Go bindings for the Google Agent-to-Agent (A2A) protocol.
package a2a

import (
	"time"
)

// Version of the A2A protocol
const Version = "1.0.0"

// Module structure:
// - schema.go: Contains the core data types for the A2A protocol
// - client/client.go: Client implementation for making A2A requests
// - server/server.go: Server implementation for handling A2A requests
// - examples/...: Example applications using the A2A library

// TaskState represents the possible states of an A2A task
type TaskState string

const (
	// TaskSubmitted indicates the task has been submitted to the agent
	TaskSubmitted TaskState = "submitted"
	// TaskWorking indicates the agent is working on the task
	TaskWorking TaskState = "working"
	// TaskInputRequired indicates the agent needs additional input to continue
	TaskInputRequired TaskState = "input-required"
	// TaskCompleted indicates the task has been completed successfully
	TaskCompleted TaskState = "completed"
	// TaskFailed indicates the task has failed
	TaskFailed TaskState = "failed"
	// TaskCanceled indicates the task was canceled
	TaskCanceled TaskState = "canceled"
)

// Role represents the role of a participant in an A2A conversation
type Role string

const (
	// RoleUser represents a user or client application
	RoleUser Role = "user"
	// RoleAgent represents an agent
	RoleAgent Role = "agent"
)

// PartType represents the type of content in a Message Part
type PartType string

const (
	// PartTypeText represents plain text content
	PartTypeText PartType = "text"
	// PartTypeFile represents file content (binary data or URI)
	PartTypeFile PartType = "file"
	// PartTypeData represents structured JSON data
	PartTypeData PartType = "data"
)

// Part represents the fundamental content unit within a Message or Artifact
type Part struct {
	Type      PartType  `json:"type"`
	Text      *string   `json:"text,omitempty"`
	Data      any       `json:"data,omitempty"`
	FileBytes *[]byte   `json:"fileBytes,omitempty"`
	FileURI   *string   `json:"fileUri,omitempty"`
	MimeType  *string   `json:"mimeType,omitempty"`
	Metadata  *Metadata `json:"metadata,omitempty"`
}

// Metadata contains additional information about a Part
type Metadata struct {
	FileName   *string `json:"fileName,omitempty"`
	Title      *string `json:"title,omitempty"`
	Additional any     `json:"additional,omitempty"`
}

// Message represents a communication turn between the client and the agent
type Message struct {
	Role  Role   `json:"role"`
	Parts []Part `json:"parts"`
}

// Artifact represents outputs generated by the agent during a task
type Artifact struct {
	Parts  []Part `json:"parts"`
	Index  int    `json:"index"`
	Title  string `json:"title,omitempty"`
	Labels any    `json:"labels,omitempty"`
}

// TaskStatus represents the current status of a task
type TaskStatus struct {
	State     TaskState  `json:"state"`
	Timestamp time.Time  `json:"timestamp"`
	Reason    *string    `json:"reason,omitempty"`
	Error     *TaskError `json:"error,omitempty"`
}

// TaskError represents an error that occurred during task execution
type TaskError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Task represents the central unit of work in the A2A protocol
type Task struct {
	ID                  string        `json:"id"`
	SessionID           string        `json:"sessionId,omitempty"`
	Status              TaskStatus    `json:"status"`
	Artifacts           []Artifact    `json:"artifacts,omitempty"`
	History             []Message     `json:"history,omitempty"`
	AcceptedOutputModes []string      `json:"acceptedOutputModes,omitempty"`
	PushNotifications   any           `json:"pushNotifications,omitempty"`
	CreatedAt           time.Time     `json:"createdAt"`
	UpdatedAt           time.Time     `json:"updatedAt"`
	Message             *Message      `json:"message,omitempty"`
	Metadata            any           `json:"metadata,omitempty"`
	ParentTask          *ParentTask   `json:"parentTask,omitempty"`
	PrevTasks           []*ParentTask `json:"prevTasks,omitempty"`
}

// ParentTask represents a reference to another task
type ParentTask struct {
	ID string `json:"id"`
}

// AgentCapabilities describes the capabilities of an agent
type AgentCapabilities struct {
	Streaming          bool   `json:"streaming,omitempty"`
	PushNotifications  bool   `json:"pushNotifications,omitempty"`
	MultiTurn          bool   `json:"multiTurn,omitempty"`
	MultiTask          bool   `json:"multiTask,omitempty"`
	AcceptedInputModes any    `json:"acceptedInputModes,omitempty"`
	DefaultOutputMode  string `json:"defaultOutputMode,omitempty"`
}

// AgentSkills describes the skills of an agent
type AgentSkills struct {
	AvailableSkills []Skill                 `json:"availableSkills,omitempty"`
	CustomSkills    map[string]SkillDetails `json:"customSkills,omitempty"`
}

// Skill represents a skill that an agent can perform
type Skill struct {
	Type     string `json:"type"`
	Required bool   `json:"required,omitempty"`
}

// SkillDetails provides detailed information about a custom skill
type SkillDetails struct {
	Description string `json:"description"`
	Parameters  any    `json:"parameters,omitempty"`
}

// AgentAuthentication describes the authentication methods supported by an agent
type AgentAuthentication struct {
	Schemes     []string `json:"schemes"`
	Credentials *string  `json:"credentials,omitempty"`
}

// AgentCard contains metadata about an agent
type AgentCard struct {
	AgentType        string               `json:"agentType"`
	Name             string               `json:"name"`
	Description      string               `json:"description"`
	Version          string               `json:"version"`
	Capabilities     AgentCapabilities    `json:"capabilities"`
	Skills           AgentSkills          `json:"skills,omitempty"`
	Authentication   *AgentAuthentication `json:"authentication,omitempty"`
	Provider         *AgentProvider       `json:"provider,omitempty"`
	AdditionalFields map[string]any       `json:"additionalFields,omitempty"`
}

// AgentProvider contains information about the provider of an agent
type AgentProvider struct {
	Name    string `json:"name"`
	Website string `json:"website,omitempty"`
}

// JsonRpcRequest represents a JSON-RPC request
type JsonRpcRequest struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	ID      any    `json:"id"`
	Params  any    `json:"params,omitempty"`
}

// JsonRpcResponse represents a JSON-RPC response
type JsonRpcResponse struct {
	JsonRpc string        `json:"jsonrpc"`
	ID      any           `json:"id"`
	Result  any           `json:"result,omitempty"`
	Error   *JsonRpcError `json:"error,omitempty"`
}

// JsonRpcError represents a JSON-RPC error
type JsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// TasksGetRequest represents a request to get task information
type TasksGetRequest struct {
	ID string `json:"id"`
}

// TasksSendRequest represents a request to send a message to start or continue a task
type TasksSendRequest struct {
	ID                  string        `json:"id"`
	SessionID           string        `json:"sessionId,omitempty"`
	AcceptedOutputModes []string      `json:"acceptedOutputModes,omitempty"`
	Message             Message       `json:"message"`
	ParentTask          *ParentTask   `json:"parentTask,omitempty"`
	PrevTasks           []*ParentTask `json:"prevTasks,omitempty"`
	Metadata            any           `json:"metadata,omitempty"`
}

// TasksCancelRequest represents a request to cancel a task
type TasksCancelRequest struct {
	ID     string `json:"id"`
	Reason string `json:"reason,omitempty"`
}

// TasksSubscribeRequest represents a request to subscribe to task updates
type TasksSubscribeRequest struct {
	ID string `json:"id"`
}

// TasksSendSubscribeRequest represents a request to send a message and subscribe to updates
type TasksSendSubscribeRequest struct {
	ID                  string        `json:"id"`
	SessionID           string        `json:"sessionId,omitempty"`
	AcceptedOutputModes []string      `json:"acceptedOutputModes,omitempty"`
	Message             Message       `json:"message"`
	ParentTask          *ParentTask   `json:"parentTask,omitempty"`
	PrevTasks           []*ParentTask `json:"prevTasks,omitempty"`
	Metadata            any           `json:"metadata,omitempty"`
}

// AgentCardRequest represents a request to get an agent's card
type AgentCardRequest struct{}

// NewTextPart creates a new text part with the provided content
func NewTextPart(text string) Part {
	return Part{
		Type: PartTypeText,
		Text: &text,
	}
}

// NewDataPart creates a new data part with the provided data
func NewDataPart(data any) Part {
	return Part{
		Type: PartTypeData,
		Data: data,
	}
}

// NewFilePart creates a new file part with the provided content
func NewFilePart(fileBytes []byte, mimeType string, fileName string) Part {
	return Part{
		Type:      PartTypeFile,
		FileBytes: &fileBytes,
		MimeType:  &mimeType,
		Metadata: &Metadata{
			FileName: &fileName,
		},
	}
}

// NewFileUriPart creates a new file part with a URI reference
func NewFileUriPart(fileUri string, mimeType string, fileName string) Part {
	return Part{
		Type:     PartTypeFile,
		FileURI:  &fileUri,
		MimeType: &mimeType,
		Metadata: &Metadata{
			FileName: &fileName,
		},
	}
}

// NewUserMessage creates a new message with the user role
func NewUserMessage(parts ...Part) Message {
	return Message{
		Role:  RoleUser,
		Parts: parts,
	}
}

// NewAgentMessage creates a new message with the agent role
func NewAgentMessage(parts ...Part) Message {
	return Message{
		Role:  RoleAgent,
		Parts: parts,
	}
}

// NewArtifact creates a new artifact with the provided parts
func NewArtifact(index int, title string, parts ...Part) Artifact {
	return Artifact{
		Parts: parts,
		Index: index,
		Title: title,
	}
}
