package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headerValue string
		expectKey   string
		expectErr   bool
	}{
		{"no header", "", "", true},
		{"wrong scheme", "Bearer abc", "", true},
		{"missing key", "ApiKey", "", true},
		{"valid", "ApiKey xyz123", "xyz123", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := http.Header{}
			if tt.headerValue != "" {
				h.Set("Authorization", tt.headerValue)
			}

			key, err := GetAPIKey(h)

			if tt.expectErr && err == nil {
				t.Fatalf("expected error, got nil")
			}
			if !tt.expectErr && err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if key != tt.expectKey {
				t.Fatalf("expected key '%s', got '%s'", tt.expectKey, key)
			}
		})
	}
}
