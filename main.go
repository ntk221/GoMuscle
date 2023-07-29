package main

import (
	"bytes"
	"io"
	"testing"
)

func TestCopyBuffer(t *testing.T) {
	// Test Case 1: Normal operation with non-empty buffer
	srcData := []byte("Hello, world!")
	var dst bytes.Buffer
	buf := make([]byte, 512)
	written, err := io.CopyBuffer(&dst, bytes.NewReader(srcData), buf)
	if err != nil {
		t.Errorf("Error during CopyBuffer: %v", err)
	}
	if written != int64(len(srcData)) {
		t.Errorf("Expected %d bytes to be written, but %d bytes were written", len(srcData), written)
	}
	dstData := dst.Bytes()
	if !bytes.Equal(dstData, srcData) {
		t.Errorf("Destination data differs from source data. Expected: %q, Got: %q", srcData, dstData)
	}

	// Test Case 2: Error case with nil buffer
	_, err = io.CopyBuffer(&dst, bytes.NewReader(srcData), nil)
	if err == nil {
		t.Error("Expected error for nil buffer, but got nil")
	}

	// Test Case 3: Panic with empty buffer
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic for empty buffer, but no panic occurred")
		}
	}()
	io.CopyBuffer(&dst, bytes.NewReader(srcData), []byte{})
}
