package service

import (
	"fmt"
	"testing"
)

func TestFormate(t *testing.T) {
	t.Log(fmt.Sprintf("%06d", 10))
}
