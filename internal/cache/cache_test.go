package cache

import "testing"

func TestCache_SetAndGet(t *testing.T) {
	// Create a new cache with a maximum size of 5.
	c := New(5)

	// Test setting a key-value pair.
	if err := c.Set("foo", "bar"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Test retrieving the value.
	value, exists := c.Get("foo")
	if !exists {
		t.Fatalf("expected key 'foo' to exist")
	}
	if value != "bar" {
		t.Fatalf("expected value 'bar', got %s", value)
	}
}

func TestCache_UpdateValue(t *testing.T) {
	c := New(5)

	// Insert an initial value.
	if err := c.Set("foo", "bar"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Update the value.
	if err := c.Set("foo", "baz"); err != nil {
		t.Fatalf("expected no error on update, got %v", err)
	}

	// Verify that the value has been updated.
	value, _ := c.Get("foo")
	if value != "baz" {
		t.Fatalf("expected updated value 'baz', got %s", value)
	}
}

func TestCache_Delete(t *testing.T) {
	c := New(5)

	// Insert a key and then delete it.
	_ = c.Set("foo", "bar")
	c.Delete("foo")

	// Verify that the key is no longer present.
	if _, exists := c.Get("foo"); exists {
		t.Fatalf("expected key 'foo' to be deleted")
	}
}

func TestCache_Full(t *testing.T) {
	// Create a cache that can hold only 2 items.
	c := New(2)

	if err := c.Set("key1", "value1"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if err := c.Set("key2", "value2"); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// This should return an error because the cache is full.
	err := c.Set("key3", "value3")
	if err == nil {
		t.Fatalf("expected an error when cache is full, got nil")
	}
	expectedError := "cache is full: maximum of 2 items reached"
	if err.Error() != expectedError {
		t.Fatalf("expected error %q, got %q", expectedError, err.Error())
	}
}
