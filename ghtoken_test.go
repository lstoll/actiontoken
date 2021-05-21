package main

import (
	"context"
	"testing"
)

func TestInvalidTokenValidate(t *testing.T) {
	ctx := context.Background()

	ok, err := validateToken(ctx, "lstoll", "actiontoken", "invalidtoken")
	if err != nil {
		t.Errorf("validating: %v", err)
	}
	if ok {
		t.Error("should be invalid, but was valid")
	}
}
