// original source: https://github.com/go-task/task/blob/master/internal/execext/devnull.go <3
package common

import (
	"io"
)

var _ io.ReadWriteCloser = devNull{}

type devNull struct{}

func (devNull) Read(p []byte) (int, error)  { return 0, io.EOF }
func (devNull) Write(p []byte) (int, error) { return len(p), nil }
func (devNull) Close() error                { return nil }
