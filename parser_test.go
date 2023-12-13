package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseJSON(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     []byte
		wantErr  error
	}{
		{
			name:     "valid JSON file",
			filename: "./tests/step1/valid.json",
			want:     []byte("{}"),
			wantErr:  nil,
		},
		{
			name:     "invalid JSON file",
			filename: "./tests/step1/invalid.json",
			want:     nil,
			wantErr:  errInvalidJSON,
		},
		{
			name:     "non-existent JSON file",
			filename: "./non-existent.json",
			want:     nil,
			wantErr:  errors.New("open ./non-existent.json: The system cannot find the file specified."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseJSON(tt.filename)

			assert.Equal(t, string(tt.want), string(got))
			if tt.wantErr != nil {
				assert.Equal(t, tt.wantErr.Error(), err.Error())
			}
		})
	}
}
