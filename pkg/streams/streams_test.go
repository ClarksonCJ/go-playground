package streams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	Processor = &Streams{}
)

func TestIntStreamGenerator(t *testing.T) {
	assert := assert.New(t)

	arrange := []int{1, 2, 3, 4, 5, 6, 7}

	stream := Processor.CreateIntStream(arrange)

	assert.Equal(7, len(stream), "Expected stream length to be 7")
}

func TestStringStreamGenerator(t *testing.T) {
	assert := assert.New(t)

	arrange := []string{"1", "2", "3", "4", "5"}

	stream := Processor.CreateStringStream(arrange)
	assert.Equal(5, len(stream), "Expected stream length to be 5")
}
func TestStreamGenerator(t *testing.T) {
	assert := assert.New(t)
	arrange := []string{"1", "2", "3", "4", "5"}
	arranged := make([]interface{}, len(arrange))
	for k, i := range arrange {
		arranged[k] = i
	}

	stream := Processor.CreateInterfaceStream(arranged)

	assert.Equal(5, len(stream), "Expected stream length to be 5")
}
