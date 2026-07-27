package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("should return an API key", func(t *testing.T) {
		apiKey, err := GetAPIKey(http.Header{"Authorization": []string{"ApiKey test-api-key"}})
		if err != nil {
			t.Errorf("Expected an API key, but got an error: %v", err)
		}
		if apiKey == "test-api-key" {
			t.Errorf("Expected API key 'test-api-key', but got '%s'", apiKey)
		}
	})
}
