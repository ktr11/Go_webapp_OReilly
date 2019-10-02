package trace

import (
	"fmt"
	"io"
)

// Tracer はコードの中での出来事を記録できるオブジェクトを表すインターフェースです。
type Tracer interface {
	Trace(...interface{})
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...interface{}) {
	t.out.Write([]byte(fmt.Sprint(a...)))
	t.out.Write([]byte("\n"))
}
