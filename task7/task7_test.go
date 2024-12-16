package task7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask7(t *testing.T) {
	Task7()
}

func TestGenerateSender(t *testing.T) {
	ch := generateSender()
	var count int
	for range ch {
		count++
	}
	assert.Equal(t, count, ChannelLen)
}
