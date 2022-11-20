package composite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewOrganization(t *testing.T) {
	got := NewOrganization().Count()
	assert.Equal(t, 2, got)
}
