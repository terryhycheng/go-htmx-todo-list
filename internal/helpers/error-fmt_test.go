package helpers_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/terryhycheng/go-todo-list/internal/helpers"
)

func TestError(t *testing.T) {
	tests := []struct{
		name string
		errorType string
		want error
	}{
		{
			name: "ErrTitleRequired",
			errorType: "ERR_TITLE_REQUIRED",
			want: errors.New("ERR_TITLE_REQUIRED"),
		},
		{
			name: "ErrUnknownError",
			errorType: "ERR_UNKNOWN_ERROR",
			want: errors.New("ERR_UNKNOWN_ERROR"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.True(t, assert.Equal(t, helpers.Error(tt.errorType), tt.want))
		})
	}
}