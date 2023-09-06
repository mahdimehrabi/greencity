package main

import (
	"fmt"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	testCases := []struct {
		length int
	}{
		{0},
		{1},
		{10},
		{20},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Length%d", tc.length), func(t *testing.T) {
			randomBytes, err := GenerateRandomBytes(tc.length)

			if err != nil {
				t.Fatalf("Error generating random string: %v", err)
			}

			if len(randomBytes) != tc.length {
				t.Errorf("Generated string length does not match%d", tc.length)
			}
		})
	}
}
