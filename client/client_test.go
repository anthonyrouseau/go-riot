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

func TestNewDevClient(t *testing.T) {
	key := "devKey"
	c, err := NewDevClient(key)
	if err != nil {
		t.Error(err)
	}
	t.Run("Variant", func(t *testing.T) {
		if v := c.Variant(); v != devClient {
			t.Errorf("Expected client variant %d, got %d", devClient, v)
		}
	})
	t.Run("API Key", func(t *testing.T) {
		if v := c.APIKey(); v != key {
			t.Errorf("Expected client APIKey %s, got %s", key, v)
		}
	})
}

func TestNewPersonalClient(t *testing.T) {
	key := "personalKey"
	c, err := NewPersonalClient(key)
	if err != nil {
		t.Error(err)
	}
	t.Run("Variant", func(t *testing.T) {
		if v := c.Variant(); v != personalClient {
			t.Errorf("Expected client variant %d, got %d", personalClient, v)
		}
	})
	t.Run("API Key", func(t *testing.T) {
		if v := c.APIKey(); v != key {
			t.Errorf("Expected client APIKey %s, got %s", key, v)
		}
	})
}

func TestNewProductionClient(t *testing.T) {
	key := "productionKey"
	c, err := NewProductionClient(key)
	if err != nil {
		t.Error(err)
	}
	t.Run("Variant", func(t *testing.T) {
		if v := c.Variant(); v != productionClient {
			t.Errorf("Expected client variant %d, got %d", productionClient, v)
		}
	})
	t.Run("API Key", func(t *testing.T) {
		if v := c.APIKey(); v != key {
			t.Errorf("Expected client APIKey %s, got %s", key, v)
		}
	})
}
