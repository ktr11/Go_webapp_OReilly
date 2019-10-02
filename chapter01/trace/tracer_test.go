package trace

import (
	"bytes"
	"io"
	"testing"
)

// TestNew 名前がTestで始まり、*tesing.T型の引数を受け取る関数はすべてユニットテストとみなされ、
// テストを実行すると、この条件を満たす関数がすべて呼び出されます。
func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("Newからの戻り値がnilです。")
	} else {
		tracer.Trace("こんにちは、traceパッケージ")
		if buf.String() != "こんにちは、traceパッケージ\n" {
			t.Errorf("'%s'という誤った文字列が出力されました。", buf.String())
		}
	}
}

// New Tracer型のオブジェクトを返す。
func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
