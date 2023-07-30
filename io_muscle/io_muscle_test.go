package io_muscle_test

import (
	"github.com/stretchr/testify/assert"
	"go_muscle/io_muscle"
	"io"
	"testing"
)

type testReader struct {
	data []byte
}

func (r *testReader) Read(p []byte) (n int, err error) {
	if len(r.data) == 0 {
		return 0, io_muscle.ErrUnexpectedEOF
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

	data2 := []byte("Hello, world!")
	r2 := &testReader{data: data2}
	buf2 := make([]byte, 10)
	min2 := 5
	realN, err2 := io.ReadAtLeast(r2, buf2, min2)
	if err2 != nil {
		t.Errorf("ReadAtLeastの実行中にエラーが起きました: %v", err2)
	}
	if realN < min2 {
		t.Errorf("少なくとも%dバイトが読み込まれるはずですが，結果は%dバイトでした", min2, realN)
	}
	if string(buf[:n]) != string(buf2[:realN]) {
		t.Errorf("expected: %v, actual: %v", string(buf2[:realN]), string(buf[:n]))
	}

	// 異常系1: minバイト未満を読み取った後にEOFが発生した場合，ReadAtLeastはErrUnexpectedEOFを返します
	// このケースはつまり，少なくともminだけ読み取りたい時に，readerの保持するバッファ列が，それより少なかったという場合である
	data = []byte("")
	r = &testReader{data: data}
	buf = make([]byte, 10)
	min = 5
	n, err = io_muscle.ReadAtLeast(r, buf, min)

	data2 = []byte("")
	r2 = &testReader{data: data2}
	buf2 = make([]byte, 10)
	min2 = 5
	realN, err2 = io.ReadAtLeast(r2, buf2, min2)
	if string(buf[:n]) != string(buf2[:realN]) {
		t.Errorf("expected: %v, actual: %v", string(buf2[:realN]), string(buf[:n]))
	}

	assert.Equal(t, err, err2)

	// 異常系2: min が bufの長さよりも大きい場合には，ReadAtLeastはErrShortBufferを返します
	data = []byte("Hello, World")
	r = &testReader{data: data}
	buf = make([]byte, 2)
	min = 5
	n, err = io_muscle.ReadAtLeast(r, buf, min)

	data2 = []byte("Hello, World")
	r2 = &testReader{data: data2}
	buf2 = make([]byte, 2)
	min2 = 5
	realN, err2 = io.ReadAtLeast(r2, buf2, min2)

	assert.Equal(t, err, err2)

}
