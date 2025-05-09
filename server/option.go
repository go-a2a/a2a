// Copyright 2025 The Go A2A Authors
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"log/slog"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

// Option represents an option for configuring the [Server].
type Option func(*Server)

// WithEndpoint sets the custom endpoint for the [Server].
func WithEndpoint(endpoint string) Option {
	return func(s *Server) {
		s.endpoint = endpoint
	}
}

// WithHandlers sets the custom handlers for the [Server].
func WithHandlers(handlers ...func(http.Handler) http.Handler) Option {
	return func(s *Server) {
		s.handlers = handlers
	}
}

// WithLogger sets the [*slog.Logger] for the [Server].
func WithLogger(logger *slog.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

// WithTracer sets the [trace.Tracer] for the [Server].
func WithTracer(tracer trace.Tracer) Option {
	return func(s *Server) {
		s.tracer = tracer
	}
}
