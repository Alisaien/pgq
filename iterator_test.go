package pgq

import (
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func TestIterator(t *testing.T) {
	var T *struct{}
	t.Log(jsoniter.Unmarshal([]byte("null"), &T))
	t.Log(T)
}
