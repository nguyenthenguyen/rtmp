package control_test

import (
	"github.com/nguyenthenguyen/rtmp/chunk"
	"github.com/nguyenthenguyen/rtmp/control"
	"github.com/stretchr/testify/mock"
)

type MockParser struct {
	mock.Mock
}

func (p *MockParser) Parse(chunk *chunk.Chunk) (control.Control, error) {
	args := p.Called(chunk)
	return args.Get(0).(control.Control), args.Error(1)
}
