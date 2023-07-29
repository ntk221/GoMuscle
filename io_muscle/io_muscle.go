package io_muscle

import (
	"errors"
	"testing"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

var ErrShortBuffer = errors.New("short buffer")
var ErrUnexpectedEOF = errors.New("unexpected EOF")

// ReadAtLeast
// ReadAtLeastは、rからbufに少なくともminバイト読み取るようにします。
// コピーされたバイト数と，読み取られたバイトが少なかった場合はエラーを返します
// バイトが一つも読み取られなかった場合に飲みEOFをエラーとして返します
// minバイト未満を読み取った後にEOFが発生した場合，ReadAtLeastはErrUnexpectedEOFを返します
// min が bufの長さよりも大きい場合には，ReadAtLeastはErrShirtBufferを返します
// minがbufの長さよりも大きい場合，ReadAtLeastはErrShortBufferを返します
// 戻り値として n >= minの場合はerr==nilです
// r が少なくともminバイトを読み取った後にエラーを返した場合，そのエラーは無視されます
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error) {
	// TODO: implement this!
}

func TestReadAtLeast(t *testing.T) {

}
