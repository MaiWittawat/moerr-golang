package moerr

import (
	"fmt"
	"testing"
)

// ------------------------------------------------------------
// Test New()
// ------------------------------------------------------------
func TestNew(t *testing.T) {
	err := New("hello world")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expected := "hello world"
	if err.Error() != expected {
		t.Fatalf("expected %q, got %q", expected, err.Error())
	}
}

func TestNewEmpty(t *testing.T) {
	err := New("")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "" {
		t.Fatalf("expected empty string, got %q", err.Error())
	}
}

// ------------------------------------------------------------
// Test NewErrorf()
// ------------------------------------------------------------
func TestNewErrorf(t *testing.T) {
	err := NewErrorf("id %d not found", 5)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	expected := "id 5 not found"
	if err.Error() != expected {
		t.Fatalf("expected %q, got %q", expected, err.Error())
	}
}

func TestNewErrorfMultipleArgs(t *testing.T) {
	err := NewErrorf("a:%d b:%s c:%v", 1, "x", true)
	expected := "a:1 b:x c:true"

	if err.Error() != expected {
		t.Fatalf("expected %q, got %q", expected, err.Error())
	}
}

func TestNewErrorfEmpty(t *testing.T) {
	err := NewErrorf("")
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err.Error() != "" {
		t.Fatalf("expected empty string, got %q", err.Error())
	}
}

func TestNewErrorfNoArgs(t *testing.T) {
	err := NewErrorf("static message")
	if err.Error() != "static message" {
		t.Fatalf("expected static message")
	}
}

// ------------------------------------------------------------
// Test errorString type behavior
// ------------------------------------------------------------
func TestErrorType(t *testing.T) {
	err := New("test")
	_, ok := err.(*errorString)
	if !ok {
		t.Fatalf("expected error to be of type *errorString, got %T", err)
	}
}

func TestErrorImplementsError(t *testing.T) {
	var _ error = (*errorString)(nil) // compile-time assertion
}

// ------------------------------------------------------------
// Test Error() method explicitly
// ------------------------------------------------------------
func TestErrorMethod(t *testing.T) {
	e := &errorString{s: "xyz"}
	if e.Error() != "xyz" {
		t.Fatalf("expected xyz, got %q", e.Error())
	}
}

// ------------------------------------------------------------
// Ensure formatting matches fmt.Sprintf behavior exactly
// ------------------------------------------------------------
func TestNewErrorfMatchesSprintf(t *testing.T) {
	format := "value=%d, key=%s, flag=%v"
	v1 := fmt.Sprintf(format, 3, "abc", false)
	v2 := NewErrorf(format, 3, "abc", false).Error()

	if v1 != v2 {
		t.Fatalf("expected %q, got %q", v1, v2)
	}
}
