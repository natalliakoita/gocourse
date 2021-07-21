package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "call main",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.NotPanics(t, main)
		})
	}
}
