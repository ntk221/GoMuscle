package io_muscle_test

import (
	"go_muscle/io_muscle"
	"testing"
)

type testReader struct {
	data []byte
}

func (r *testReader) Read(p []byte) (n int, err error) {
	if len(r.data) == 0 {
		return 0, ErrUnexpectedEOF
	}
	n = copy(p, r.data)
	r.data = r.data[n:]
	return n, nil
}

func TestReadAtLeast(t *testing.T) {
	// 正常系
	data := []byte("Hello, world!")
	r := &testReader{data: data}
	buf := make([]byte, 10)
	min := 5
	n, err := io_muscle.ReadAtLeast(r, buf, min)
	if err != nil {
		t.Errorf("ReadAtLeastの実行中にエラーが起きました: %v", err)
	}
	if n < min {
		t.Errorf("少なくとも%dバイトが読み込まれるはずですが，結果は%dバイトでした", min, n)
	}
	expectedData := data[:min]
	if string(buf[:n]) != string(expectedData) {
		t.Errorf("")
	}

}
