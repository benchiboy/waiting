package util

import (
	"testing"
)

func TestHttp_get(t *testing.T) {
	Http_get(map[string]string{"a": "1", "b": "2", "c": "3"})
}
