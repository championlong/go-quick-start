package bridge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBridge(t *testing.T) {
	sender := NewEmailMsgSender([]string{"test@test.com"})
	n := NewSevereNotification(sender)
	err := n.Notify("test msg")

	assert.Nil(t, err)
}
