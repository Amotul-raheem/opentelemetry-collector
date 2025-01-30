// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
)

func TestHistogramDataPointSlice(t *testing.T) {
	es := NewHistogramDataPointSlice()
	assert.Equal(t, 0, es.Len())
	state := internal.StateMutable
	es = newHistogramDataPointSlice(&[]*otlpmetrics.HistogramDataPoint{}, &state)
	assert.Equal(t, 0, es.Len())

	emptyVal := NewHistogramDataPoint()
	testVal := generateTestHistogramDataPoint()
	for i := 0; i < 7; i++ {
		el := es.AppendEmpty()
		assert.Equal(t, emptyVal, es.At(i))
		fillTestHistogramDataPoint(el)
		assert.Equal(t, testVal, es.At(i))
	}
	assert.Equal(t, 7, es.Len())
}

func TestHistogramDataPointSliceReadOnly(t *testing.T) {
	sharedState := internal.StateReadOnly
	es := newHistogramDataPointSlice(&[]*otlpmetrics.HistogramDataPoint{}, &sharedState)
	assert.Equal(t, 0, es.Len())
	assert.Panics(t, func() { es.AppendEmpty() })
	assert.Panics(t, func() { es.EnsureCapacity(2) })
	es2 := NewHistogramDataPointSlice()
	es.CopyTo(es2)
	assert.Panics(t, func() { es2.CopyTo(es) })
	assert.Panics(t, func() { es.MoveAndAppendTo(es2) })
	assert.Panics(t, func() { es2.MoveAndAppendTo(es) })
}

func TestHistogramDataPointSlice_CopyTo(t *testing.T) {
	dest := NewHistogramDataPointSlice()
	// Test CopyTo to empty
	NewHistogramDataPointSlice().CopyTo(dest)
	assert.Equal(t, NewHistogramDataPointSlice(), dest)

	// Test CopyTo larger slice
	generateTestHistogramDataPointSlice().CopyTo(dest)
	assert.Equal(t, generateTestHistogramDataPointSlice(), dest)

	// Test CopyTo same size slice
	generateTestHistogramDataPointSlice().CopyTo(dest)
	assert.Equal(t, generateTestHistogramDataPointSlice(), dest)
}

func TestHistogramDataPointSlice_EnsureCapacity(t *testing.T) {
	es := generateTestHistogramDataPointSlice()

	// Test ensure smaller capacity.
	const ensureSmallLen = 4
	es.EnsureCapacity(ensureSmallLen)
	assert.Less(t, ensureSmallLen, es.Len())
	assert.Equal(t, es.Len(), cap(*es.orig))
	assert.Equal(t, generateTestHistogramDataPointSlice(), es)

	// Test ensure larger capacity
	const ensureLargeLen = 9
	es.EnsureCapacity(ensureLargeLen)
	assert.Less(t, generateTestHistogramDataPointSlice().Len(), ensureLargeLen)
	assert.Equal(t, ensureLargeLen, cap(*es.orig))
	assert.Equal(t, generateTestHistogramDataPointSlice(), es)
}

func TestHistogramDataPointSlice_MoveAndAppendTo(t *testing.T) {
	// Test MoveAndAppendTo to empty
	expectedSlice := generateTestHistogramDataPointSlice()
	dest := NewHistogramDataPointSlice()
	src := generateTestHistogramDataPointSlice()
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestHistogramDataPointSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo empty slice
	src.MoveAndAppendTo(dest)
	assert.Equal(t, generateTestHistogramDataPointSlice(), dest)
	assert.Equal(t, 0, src.Len())
	assert.Equal(t, expectedSlice.Len(), dest.Len())

	// Test MoveAndAppendTo not empty slice
	generateTestHistogramDataPointSlice().MoveAndAppendTo(dest)
	assert.Equal(t, 2*expectedSlice.Len(), dest.Len())
	for i := 0; i < expectedSlice.Len(); i++ {
		assert.Equal(t, expectedSlice.At(i), dest.At(i))
		assert.Equal(t, expectedSlice.At(i), dest.At(i+expectedSlice.Len()))
	}
}

func TestHistogramDataPointSlice_RemoveIf(t *testing.T) {
	// Test RemoveIf on empty slice
	emptySlice := NewHistogramDataPointSlice()
	emptySlice.RemoveIf(func(el HistogramDataPoint) bool {
		t.Fail()
		return false
	})

	// Test RemoveIf
	filtered := generateTestHistogramDataPointSlice()
	pos := 0
	filtered.RemoveIf(func(el HistogramDataPoint) bool {
		pos++
		return pos%3 == 0
	})
	assert.Equal(t, 5, filtered.Len())
}

func TestHistogramDataPointSlice_Sort(t *testing.T) {
	es := generateTestHistogramDataPointSlice()
	es.Sort(func(a, b HistogramDataPoint) bool {
		return uintptr(unsafe.Pointer(a.orig)) < uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.Less(t, uintptr(unsafe.Pointer(es.At(i-1).orig)), uintptr(unsafe.Pointer(es.At(i).orig)))
	}
	es.Sort(func(a, b HistogramDataPoint) bool {
		return uintptr(unsafe.Pointer(a.orig)) > uintptr(unsafe.Pointer(b.orig))
	})
	for i := 1; i < es.Len(); i++ {
		assert.Greater(t, uintptr(unsafe.Pointer(es.At(i-1).orig)), uintptr(unsafe.Pointer(es.At(i).orig)))
	}
}

func generateTestHistogramDataPointSlice() HistogramDataPointSlice {
	es := NewHistogramDataPointSlice()
	fillTestHistogramDataPointSlice(es)
	return es
}

func fillTestHistogramDataPointSlice(es HistogramDataPointSlice) {
	*es.orig = make([]*otlpmetrics.HistogramDataPoint, 7)
	for i := 0; i < 7; i++ {
		(*es.orig)[i] = &otlpmetrics.HistogramDataPoint{}
		fillTestHistogramDataPoint(newHistogramDataPoint((*es.orig)[i], es.state))
	}
}
