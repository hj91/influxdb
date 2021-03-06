// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: arrays.gen.go.tmpl

package gen

import (
	"github.com/influxdata/influxdb/tsdb/cursors"
	"github.com/influxdata/influxdb/tsdb/tsm1"
)

type FloatValues interface {
	Copy(*cursors.FloatArray)
}

type floatArray struct {
	cursors.FloatArray
}

func newFloatArrayLen(sz int) *floatArray {
	return &floatArray{
		FloatArray: cursors.FloatArray{
			Timestamps: make([]int64, sz),
			Values:     make([]float64, sz),
		},
	}
}

func (a *floatArray) Encode(b []byte) ([]byte, error) {
	return tsm1.EncodeFloatArrayBlock(&a.FloatArray, b)
}

func (a *floatArray) Copy(dst *cursors.FloatArray) {
	dst.Timestamps = append(dst.Timestamps[:0], a.Timestamps...)
	dst.Values = append(dst.Values[:0], a.Values...)
}

type IntegerValues interface {
	Copy(*cursors.IntegerArray)
}

type integerArray struct {
	cursors.IntegerArray
}

func newIntegerArrayLen(sz int) *integerArray {
	return &integerArray{
		IntegerArray: cursors.IntegerArray{
			Timestamps: make([]int64, sz),
			Values:     make([]int64, sz),
		},
	}
}

func (a *integerArray) Encode(b []byte) ([]byte, error) {
	return tsm1.EncodeIntegerArrayBlock(&a.IntegerArray, b)
}

func (a *integerArray) Copy(dst *cursors.IntegerArray) {
	dst.Timestamps = append(dst.Timestamps[:0], a.Timestamps...)
	dst.Values = append(dst.Values[:0], a.Values...)
}

type UnsignedValues interface {
	Copy(*cursors.UnsignedArray)
}

type unsignedArray struct {
	cursors.UnsignedArray
}

func newUnsignedArrayLen(sz int) *unsignedArray {
	return &unsignedArray{
		UnsignedArray: cursors.UnsignedArray{
			Timestamps: make([]int64, sz),
			Values:     make([]uint64, sz),
		},
	}
}

func (a *unsignedArray) Encode(b []byte) ([]byte, error) {
	return tsm1.EncodeUnsignedArrayBlock(&a.UnsignedArray, b)
}

func (a *unsignedArray) Copy(dst *cursors.UnsignedArray) {
	dst.Timestamps = append(dst.Timestamps[:0], a.Timestamps...)
	dst.Values = append(dst.Values[:0], a.Values...)
}

type StringValues interface {
	Copy(*cursors.StringArray)
}

type stringArray struct {
	cursors.StringArray
}

func newStringArrayLen(sz int) *stringArray {
	return &stringArray{
		StringArray: cursors.StringArray{
			Timestamps: make([]int64, sz),
			Values:     make([]string, sz),
		},
	}
}

func (a *stringArray) Encode(b []byte) ([]byte, error) {
	return tsm1.EncodeStringArrayBlock(&a.StringArray, b)
}

func (a *stringArray) Copy(dst *cursors.StringArray) {
	dst.Timestamps = append(dst.Timestamps[:0], a.Timestamps...)
	dst.Values = append(dst.Values[:0], a.Values...)
}

type BooleanValues interface {
	Copy(*cursors.BooleanArray)
}

type booleanArray struct {
	cursors.BooleanArray
}

func newBooleanArrayLen(sz int) *booleanArray {
	return &booleanArray{
		BooleanArray: cursors.BooleanArray{
			Timestamps: make([]int64, sz),
			Values:     make([]bool, sz),
		},
	}
}

func (a *booleanArray) Encode(b []byte) ([]byte, error) {
	return tsm1.EncodeBooleanArrayBlock(&a.BooleanArray, b)
}

func (a *booleanArray) Copy(dst *cursors.BooleanArray) {
	dst.Timestamps = append(dst.Timestamps[:0], a.Timestamps...)
	dst.Values = append(dst.Values[:0], a.Values...)
}
