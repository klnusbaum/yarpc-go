// Code generated by thriftrw v1.5.0. DO NOT EDIT.
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gauntlet

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type SecondService_BlahBlah_Args struct{}

func (v *SecondService_BlahBlah_Args) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SecondService_BlahBlah_Args) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *SecondService_BlahBlah_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [0]string
	i := 0
	return fmt.Sprintf("SecondService_BlahBlah_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *SecondService_BlahBlah_Args) Equals(rhs *SecondService_BlahBlah_Args) bool {
	return true
}

func (v *SecondService_BlahBlah_Args) MethodName() string {
	return "blahBlah"
}

func (v *SecondService_BlahBlah_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var SecondService_BlahBlah_Helper = struct {
	Args           func() *SecondService_BlahBlah_Args
	IsException    func(error) bool
	WrapResponse   func(error) (*SecondService_BlahBlah_Result, error)
	UnwrapResponse func(*SecondService_BlahBlah_Result) error
}{}

func init() {
	SecondService_BlahBlah_Helper.Args = func() *SecondService_BlahBlah_Args {
		return &SecondService_BlahBlah_Args{}
	}
	SecondService_BlahBlah_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	SecondService_BlahBlah_Helper.WrapResponse = func(err error) (*SecondService_BlahBlah_Result, error) {
		if err == nil {
			return &SecondService_BlahBlah_Result{}, nil
		}
		return nil, err
	}
	SecondService_BlahBlah_Helper.UnwrapResponse = func(result *SecondService_BlahBlah_Result) (err error) {
		return
	}
}

type SecondService_BlahBlah_Result struct{}

func (v *SecondService_BlahBlah_Result) ToWire() (wire.Value, error) {
	var (
		fields [0]wire.Field
		i      int = 0
	)
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SecondService_BlahBlah_Result) FromWire(w wire.Value) error {
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		}
	}
	return nil
}

func (v *SecondService_BlahBlah_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [0]string
	i := 0
	return fmt.Sprintf("SecondService_BlahBlah_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *SecondService_BlahBlah_Result) Equals(rhs *SecondService_BlahBlah_Result) bool {
	return true
}

func (v *SecondService_BlahBlah_Result) MethodName() string {
	return "blahBlah"
}

func (v *SecondService_BlahBlah_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
