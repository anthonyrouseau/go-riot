package client

import "testing"

func TestNewClient(t *testing.T) {
	key := "testkey"
	c, err := NewClient(key)
	if err != nil {
		t.Error(err)
	}
	t.Run("Variant", func(t *testing.T) {
		if v := c.Variant(); v != unspecifiedClient {
			t.Errorf("Expected client variant %d, got %d", unspecifiedClient, v)
		}
	})
	t.Run("API Key", func(t *testing.T) {
		if v := c.APIKey(); v != key {
			t.Errorf("Expected client APIKey %s, got %s", key, v)
		}
	})
}
