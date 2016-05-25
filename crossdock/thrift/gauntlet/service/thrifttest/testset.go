// Code generated by thriftrw

// Copyright (c) 2016 Uber Technologies, Inc.
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

package thrifttest

import (
	"errors"
	"fmt"
	"github.com/thriftrw/thriftrw-go/wire"
	"strings"
)

type TestSetArgs struct {
	Thing map[int32]struct{} `json:"thing"`
}

type _Set_I32_ValueList map[int32]struct{}

func (v _Set_I32_ValueList) ForEach(f func(wire.Value) error) error {
	for x := range v {
		w, err := wire.NewValueI32(x), error(nil)
		if err != nil {
			return err
		}
		err = f(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (v _Set_I32_ValueList) Close() {
}

func (v *TestSetArgs) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Thing != nil {
		w, err = wire.NewValueSet(wire.Set{ValueType: wire.TI32, Size: len(v.Thing), Items: _Set_I32_ValueList(v.Thing)}), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Set_I32_Read(s wire.Set) (map[int32]struct{}, error) {
	if s.ValueType != wire.TI32 {
		return nil, nil
	}
	o := make(map[int32]struct{}, s.Size)
	err := s.Items.ForEach(func(x wire.Value) error {
		i, err := x.GetI32(), error(nil)
		if err != nil {
			return err
		}
		o[i] = struct{}{}
		return nil
	})
	s.Items.Close()
	return o, err
}

func (v *TestSetArgs) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TSet {
				v.Thing, err = _Set_I32_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *TestSetArgs) String() string {
	var fields [1]string
	i := 0
	if v.Thing != nil {
		fields[i] = fmt.Sprintf("Thing: %v", v.Thing)
		i++
	}
	return fmt.Sprintf("TestSetArgs{%v}", strings.Join(fields[:i], ", "))
}

func (v *TestSetArgs) MethodName() string {
	return "testSet"
}

func (v *TestSetArgs) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

type TestSetResult struct {
	Success map[int32]struct{} `json:"success"`
}

func (v *TestSetResult) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueSet(wire.Set{ValueType: wire.TI32, Size: len(v.Success), Items: _Set_I32_ValueList(v.Success)}), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("TestSetResult should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *TestSetResult) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TSet {
				v.Success, err = _Set_I32_Read(field.Value.GetSet())
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("TestSetResult should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *TestSetResult) String() string {
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("TestSetResult{%v}", strings.Join(fields[:i], ", "))
}

func (v *TestSetResult) MethodName() string {
	return "testSet"
}

func (v *TestSetResult) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}

var TestSetHelper = struct {
	IsException    func(error) bool
	Args           func(thing map[int32]struct{}) *TestSetArgs
	WrapResponse   func(map[int32]struct{}, error) (*TestSetResult, error)
	UnwrapResponse func(*TestSetResult) (map[int32]struct{}, error)
}{}

func init() {
	TestSetHelper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	TestSetHelper.Args = func(thing map[int32]struct{}) *TestSetArgs {
		return &TestSetArgs{Thing: thing}
	}
	TestSetHelper.WrapResponse = func(success map[int32]struct{}, err error) (*TestSetResult, error) {
		if err == nil {
			return &TestSetResult{Success: success}, nil
		}
		return nil, err
	}
	TestSetHelper.UnwrapResponse = func(result *TestSetResult) (success map[int32]struct{}, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}
