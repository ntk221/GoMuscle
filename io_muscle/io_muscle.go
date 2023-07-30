package io_muscle

import (
	"errors"
)

// Reader
// Readerは基本的なReadメソッドをラップするインターフェースである
// Readは最大でlen(p)バイトを読み込みます。返り値は，読み込んだバイト数(0<=n<=len(p))と，発生したエラーを返す
// Readがn < len(p)を返しても、呼び出し中にpをスクラッチスペースとして使用することがあります。 <- ???
// もし，一部のデータが利用可能だがlen(p)バイトでない場合，Readは待つ代わりに利用可能なデータを返します。
// Readはn　>　0のバイトを正常に読み込んだ後にエラーまたはファイルの終端に遭遇すると，読み込んだバイト数を返します。
// その場での呼び出しから(非nilの)エラーを返すことも，後続の呼び出しからエラー(およびn==0)を返すこともあります。
// 一般的なケースとして，入力ストリームの終わりで非ゼロのバイト数を返すReaderは，err==EOFまたはerr==nilのいずれかを返すかもしれない
// 次のReadは０，EOFを返すべきである
// 呼び出し元はエラーerrを考慮する前に常に返されたn>0のバイトを処理する必要があります。
// これにより，一部のバイトを読み取った後に発生するI/Oエラーや，許可されているEOFの動作の両方を正しく処理できます。
// len(p)==0の場合，Readは常にn==0を返すべきである。EOFなど特定のエラー条件が既知の場合は，非nilのエラーを返すことができる
// Readの実装は，len(p)==0の場合を除いて，0バイトのカウントとnilエラーを返すことを避けるべきです。
// 呼び出しもとは，0とnilを返す場合には何も怒らなかったとして処理し，特にEOFを示すものではないと扱うべきです
// 実装はpを保持してはならない。
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
// バイトが一つも読み取られなかった場合にのみEOFをエラーとして返します
// minバイト未満を読み取った後にEOFが発生した場合，ReadAtLeastはErrUnexpectedEOFを返します
// min が bufの長さよりも大きい場合には，ReadAtLeastはErrShirtBufferを返します
// minがbufの長さよりも大きい場合，ReadAtLeastはErrShortBufferを返します
// 戻り値として n >= minの場合はerr==nilです
// r が少なくともminバイトを読み取った後にエラーを返した場合，そのエラーは無視されます
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error) {
	// TODO: implement this !
}
