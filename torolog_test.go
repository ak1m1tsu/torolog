package torolog_test

import (
	"bytes"
	"errors"
	"testing"
	"time"

	"github.com/romankravchuk/torolog"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

type obj struct {
	Value string
}

func (o obj) MarshalZerologObject(e *zerolog.Event) {
	e.Str("value", o.Value)
}

func TestInfo(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := torolog.New(out)
		log.Info("", torolog.Fields{})
		assert.Equal(t, `{"level":"info","data":[]}`+"\n", out.String())
	})

	t.Run("one-field", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := torolog.New(out)
		log.Info("", torolog.Fields{{"foo", "bar"}})
		assert.Equal(t, `{"level":"info","data":[{"foo":"bar"}]}`+"\n", out.String())
	})

	t.Run("two-field", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := torolog.New(out)
		log.Info("", torolog.Fields{{"foo", "bar"}, {"n", 999}})
		assert.Equal(t, `{"level":"info","data":[{"foo":"bar"},{"n":999}]}`+"\n", out.String())
	})
}

func TestFields(t *testing.T) {
	out := &bytes.Buffer{}
	log := torolog.New(out)
	log.Info("", torolog.Fields{
		{"nil", nil},
		{"string", "foo"},
		{"bytes", []byte("bar")},
		{"error", errors.New("some error")},
		{"bool", true},
		{"int", int(1)},
		{"int8", int8(2)},
		{"int16", int16(3)},
		{"int32", int32(4)},
		{"int64", int64(5)},
		{"uint", uint(6)},
		{"uint8", uint8(7)},
		{"uint16", uint16(8)},
		{"uint32", uint32(9)},
		{"uint64", uint64(10)},
		{"float32", float32(11)},
		{"float64", float64(12)},
		{"dur", 1 * time.Second},
		{"time", time.Time{}},
		{"obj", obj{"a"}},
	})
	assert.Equal(t,
		`{"level":"info","data":[{"nil":null},{"string":"foo"},{"bytes":"YmFy"},{"error":{}},{"bool":true},{"int":1},{"int8":2},{"int16":3},{"int32":4},{"int64":5},{"uint":6},{"uint8":7},{"uint16":8},{"uint32":9},{"uint64":10},{"float32":11},{"float64":12},{"dur":1000000000},{"time":"0001-01-01T00:00:00Z"},{"obj":{"value":"a"}}]}`+"\n",
		out.String(),
	)
}

func TestLevel(t *testing.T) {
	t.Run("trace", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := torolog.NewWithLevel(out, torolog.TraceLevel)
		log.Trace("", errors.New(""), torolog.Fields{})
		assert.Equal(t, `{"level":"trace","data":[],"error":""}`+"\n", out.String())
	})
	t.Run("debug", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := torolog.NewWithLevel(out, torolog.DebugLevel)
		log.Debug("", torolog.Fields{})
		assert.Equal(t, `{"level":"debug","data":[]}`+"\n", out.String())
	})
	t.Run("info", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := torolog.NewWithLevel(out, torolog.InfoLevel)
		log.Info("", torolog.Fields{})
		assert.Equal(t, `{"level":"info","data":[]}`+"\n", out.String())
	})
	t.Run("warn", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := torolog.NewWithLevel(out, torolog.WarnLevel)
		log.Warn("", torolog.Fields{})
		assert.Equal(t, `{"level":"warn","data":[]}`+"\n", out.String())
	})
	t.Run("error", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := torolog.NewWithLevel(out, torolog.ErrorLevel)
		log.Error("", errors.New(""), torolog.Fields{})
		assert.Equal(t, `{"level":"error","data":[],"error":""}`+"\n", out.String())
	})
	t.Run("fatal", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := torolog.NewWithLevel(out, torolog.FatalLevel)
		log.Fatal("", errors.New(""), torolog.Fields{})
		assert.Equal(t, `{"level":"fatal","data":[],"error":""}`+"\n", out.String())
	})
	t.Run("panic", func(t *testing.T) {
		out := &bytes.Buffer{}
		log := torolog.NewWithLevel(out, torolog.PanicLevel)
		log.Panic("", errors.New(""), torolog.Fields{})
		assert.Equal(t, `{"level":"panic","data":[],"error":""}`+"\n", out.String())
	})
}
