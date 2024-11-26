package valid

import (
	"fmt"
	"testing"
)

func TestValid(t *testing.T) {
	tests := []struct {
		name   string
		result interface{}
		err    error
		want   bool
	}{
		{
			name:   "Valid JSON with nil Error",
			result: map[string]interface{}{"key": "value"},
			err:    nil,
			want:   true,
		},
		{
			name:   "Invalid JSON with parse error",
			result: nil,
			err:    fmt.Errorf("unexpected end of input"),
			want:   false,
		},
		{
			name:   "Invalid JSON with syntax error",
			result: nil,
			err:    fmt.Errorf("unexpected end of input"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.result, tt.err)
			if got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}
