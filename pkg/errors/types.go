// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

package errors

import "errors"

var (
	// ErrMalformedEntity indicates a malformed entity specification.
	ErrMalformedEntity = New("malformed entity specification")

	// ErrUnsupportedContentType indicates invalid content type.
	ErrUnsupportedContentType = errors.New("invalid content type")

	// ErrUnidentified indicates unidentified error.
	ErrUnidentified = errors.New("unidentified error")

	// ErrEmptyPath indicates empty file path.
	ErrEmptyPath = errors.New("empty file path")
)
