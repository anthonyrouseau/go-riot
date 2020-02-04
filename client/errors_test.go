package client

import "testing"

func TestErrors(t *testing.T) {
	t.Run("Format", func(t *testing.T) {
		expected := "Error 404: Data Not Found"
		newRiotError := &riotError{
			Status: &riotErrorStatus{
				Code:    404,
				Message: "Data Not Found",
			},
		}
		err := newRiotError.Format()
		if message := err.Error(); message != expected {
			t.Errorf("Expected error message to be '%s' but got '%s'", expected, message)
		}
	})
}
