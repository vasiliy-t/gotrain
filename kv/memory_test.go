package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet_NonExistingKey_ReturnsNil(t *testing.T) {
	srv := NewServer()
	assert.Nil(t, srv.Get("non_existing_key"))
}

func TestSet_OneValue_StoreContainsOneElement(t *testing.T) {
	srv := NewServer()
	val := "afadasdasd"
	srv.Set("qwerty", &val)
	assert.Equal(t, 1, len(srv.store))
}

func TestDel_RemovesExistingValue(t *testing.T) {
	srv := NewServer()
	data := []struct {
		key string
		val string
	}{
		{
			"qwerty",
			"qwerty",
		},
	}
	srv.Set(data[0].key, &data[0].val)
	srv.Del(data[0].key)
	assert.Equal(t, len(srv.store), 0)
}
