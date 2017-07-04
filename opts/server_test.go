package opts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewServer_WithMaxClients_SetsServerMaxClients(t *testing.T) {
	srv := NewServer(WithMaxClients(10))
	assert.Equal(t, srv.MaxClients, 10)
}
