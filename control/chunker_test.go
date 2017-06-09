package control_test

import (
	"github.com/nguyenthenguyen/rtmp/chunk"
	"github.com/nguyenthenguyen/rtmp/control"
	"github.com/stretchr/testify/mock"
)

type MockChunker struct {
	mock.Mock
}

func (c *MockChunker) Chunk(control control.Control) (*chunk.Chunk, error) {
	args := c.Called(control)
	return args.Get(0).(*chunk.Chunk), args.Error(1)
}
