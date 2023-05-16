package gen_uuid_test

import (
	"github.com/google/uuid"
	"testing"
)

func TestNewString(t *testing.T) {
	u := uuid.New()
	t.Logf("testing new string: %v", u)
}
