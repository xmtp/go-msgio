//
// Adapted from gogo/protobuf to use multiformats/go-varint for
// efficient, interoperable length-prefixing.
//
// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2013, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
package protoio_test

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"testing"
	"time"

	"github.com/xmtp/go-msgio/test"

	"github.com/multiformats/go-varint"
	"github.com/xmtp/go-msgio/protoio"
)

func TestVarintNormal(t *testing.T) {
	buf := newBuffer()
	writer := protoio.NewDelimitedWriter(buf)
	reader := protoio.NewDelimitedReader(buf, 1024*1024)
	if err := iotest(writer, reader); err != nil {
		t.Error(err)
	}
	if !buf.closed {
		t.Fatalf("did not close buffer")
	}
}

func TestVarintNoClose(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	writer := protoio.NewDelimitedWriter(buf)
	reader := protoio.NewDelimitedReader(buf, 1024*1024)
	if err := iotest(writer, reader); err != nil {
		t.Error(err)
	}
}

// https://github.com/gogo/protobuf/issues/32
func TestVarintMaxSize(t *testing.T) {
	buf := newBuffer()
	writer := protoio.NewDelimitedWriter(buf)
	reader := protoio.NewDelimitedReader(buf, 20)
	if err := iotest(writer, reader); err != io.ErrShortBuffer {
		t.Error(err)
	} else {
		t.Logf("%s", err)
	}
}

func TestVarintError(t *testing.T) {
	buf := newBuffer()
	// beyond uvarint63 capacity.
	buf.Write([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff})
	reader := protoio.NewDelimitedReader(buf, 1024*1024)
	msg := &test.NinOptNative{}
	err := reader.ReadMsg(msg)
	if err != varint.ErrOverflow {
		t.Fatalf("expected varint.ErrOverflow error")
	}
}

type buffer struct {
	*bytes.Buffer
	closed bool
}

func (b *buffer) Close() error {
	b.closed = true
	return nil
}

func newBuffer() *buffer {
	return &buffer{bytes.NewBuffer(nil), false}
}

func iotest(writer protoio.WriteCloser, reader protoio.ReadCloser) error {
	size := 1000
	msgs := make([]*test.NinOptNative, size)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range msgs {
		msgs[i] = NewPopulatedNinOptNative(r, true)
		err := writer.WriteMsg(msgs[i])
		if err != nil {
			return err
		}
	}
	if err := writer.Close(); err != nil {
		return err
	}
	i := 0
	for {
		msg := &test.NinOptNative{}
		if err := reader.ReadMsg(msg); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if err := verboseEqual(msg, msgs[i]); err != nil {
			return err
		}
		i++
	}
	if i != size {
		panic("not enough messages read")
	}
	if err := reader.Close(); err != nil {
		return err
	}
	return nil
}

// Everything below this line is borrowed from the generated test code in gogoproto
// --------------------------------------------------------------------------------
type randyThetest interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func verboseEqual(this *test.NinOptNative, that1 *test.NinOptNative) error {
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *NidOptNative but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *NidOptNative but is not nil && this == nil")
	}
	if this.GetField1() != that1.GetField1() {
		return fmt.Errorf("Field1 this(%v) Not Equal that(%v)", this.Field1, that1.Field1)
	}
	if this.GetField2() != that1.GetField2() {
		return fmt.Errorf("Field2 this(%v) Not Equal that(%v)", this.Field2, that1.Field2)
	}
	if this.GetField3() != that1.GetField3() {
		return fmt.Errorf("GetField3() this(%v) Not Equal that(%v)", this.GetField3(), that1.GetField3())
	}
	if this.GetField4() != that1.GetField4() {
		return fmt.Errorf("GetField4() this(%v) Not Equal that(%v)", this.GetField4(), that1.GetField4())
	}
	if this.GetField5() != that1.GetField5() {
		return fmt.Errorf("GetField5() this(%v) Not Equal that(%v)", this.GetField5(), that1.GetField5())
	}
	if this.GetField6() != that1.GetField6() {
		return fmt.Errorf("GetField6() this(%v) Not Equal that(%v)", this.GetField6(), that1.GetField6())
	}
	if this.GetField7() != that1.GetField7() {
		return fmt.Errorf("GetField7() this(%v) Not Equal that(%v)", this.GetField7(), that1.GetField7())
	}
	if this.GetField8() != that1.GetField8() {
		return fmt.Errorf("GetField8() this(%v) Not Equal that(%v)", this.GetField8(), that1.GetField8())
	}
	if this.GetField9() != that1.GetField9() {
		return fmt.Errorf("GetField9() this(%v) Not Equal that(%v)", this.GetField9(), that1.GetField9())
	}
	if this.GetField10() != that1.GetField10() {
		return fmt.Errorf("GetField10() this(%v) Not Equal that(%v)", this.GetField10(), that1.GetField10())
	}
	if this.GetField11() != that1.GetField11() {
		return fmt.Errorf("GetField11() this(%v) Not Equal that(%v)", this.GetField11(), that1.GetField11())
	}
	if this.GetField12() != that1.GetField12() {
		return fmt.Errorf("GetField12() this(%v) Not Equal that(%v)", this.GetField12(), that1.GetField12())
	}
	if this.GetField13() != that1.GetField13() {
		return fmt.Errorf("GetField13() this(%v) Not Equal that(%v)", this.GetField13(), that1.GetField13())
	}
	if this.GetField14() != that1.GetField14() {
		return fmt.Errorf("GetField14() this(%v) Not Equal that(%v)", this.GetField14(), that1.GetField14())
	}
	if !bytes.Equal(this.Field15, that1.Field15) {
		return fmt.Errorf("Field15 this(%v) Not Equal that(%v)", this.Field15, that1.Field15)
	}
	return nil
}

func NewPopulatedNinOptNative(r randyThetest, easy bool) *test.NinOptNative {
	this := &test.NinOptNative{}
	if r.Intn(5) != 0 {
		v2 := float64(r.Float64())
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		this.Field1 = &v2
	}
	if r.Intn(5) != 0 {
		v3 := float32(r.Float32())
		if r.Intn(2) == 0 {
			v3 *= -1
		}
		this.Field2 = &v3
	}
	if r.Intn(5) != 0 {
		v4 := int32(r.Int31())
		if r.Intn(2) == 0 {
			v4 *= -1
		}
		this.Field3 = &v4
	}
	if r.Intn(5) != 0 {
		v5 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		this.Field4 = &v5
	}
	if r.Intn(5) != 0 {
		v6 := uint32(r.Uint32())
		this.Field5 = &v6
	}
	if r.Intn(5) != 0 {
		v7 := uint64(uint64(r.Uint32()))
		this.Field6 = &v7
	}
	if r.Intn(5) != 0 {
		v8 := int32(r.Int31())
		if r.Intn(2) == 0 {
			v8 *= -1
		}
		this.Field7 = &v8
	}
	if r.Intn(5) != 0 {
		v9 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v9 *= -1
		}
		this.Field8 = &v9
	}
	if r.Intn(5) != 0 {
		v10 := uint32(r.Uint32())
		this.Field9 = &v10
	}
	if r.Intn(5) != 0 {
		v11 := int32(r.Int31())
		if r.Intn(2) == 0 {
			v11 *= -1
		}
		this.Field10 = &v11
	}
	if r.Intn(5) != 0 {
		v12 := uint64(uint64(r.Uint32()))
		this.Field11 = &v12
	}
	if r.Intn(5) != 0 {
		v13 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v13 *= -1
		}
		this.Field12 = &v13
	}
	if r.Intn(5) != 0 {
		v14 := bool(bool(r.Intn(2) == 0))
		this.Field13 = &v14
	}
	if r.Intn(5) != 0 {
		v15 := string(randStringThetest(r))
		this.Field14 = &v15
	}
	if r.Intn(5) != 0 {
		v16 := r.Intn(100)
		this.Field15 = make([]byte, v16)
		for i := 0; i < v16; i++ {
			this.Field15[i] = byte(r.Intn(256))
		}
	}

	return this
}

func randStringThetest(r randyThetest) string {
	v258 := r.Intn(100)
	tmps := make([]rune, v258)
	for i := 0; i < v258; i++ {
		tmps[i] = randUTF8RuneThetest(r)
	}
	return string(tmps)
}

func randUTF8RuneThetest(r randyThetest) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
