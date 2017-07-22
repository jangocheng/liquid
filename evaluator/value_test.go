package evaluator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValue_Interface(t *testing.T) {
	nv := ValueOf(nil)
	iv := ValueOf(123)
	require.Equal(t, nil, nv.Interface())
	require.Equal(t, true, ValueOf(true).Interface())
	require.Equal(t, false, ValueOf(false).Interface())
	require.Equal(t, 123, iv.Interface())
}

func TestValue_Equal(t *testing.T) {
	iv := ValueOf(123)
	require.True(t, iv.Equal(ValueOf(123)))
	require.True(t, iv.Equal(ValueOf(123.0)))
}

func TestValue_Less(t *testing.T) {
	iv := ValueOf(123)
	require.False(t, iv.Less(ValueOf(100)))
	require.True(t, iv.Less(ValueOf(200)))
	require.False(t, iv.Less(ValueOf(100.5)))
	require.True(t, iv.Less(ValueOf(200.5)))

	sv := ValueOf("b")
	require.False(t, sv.Less(ValueOf("a")))
	require.True(t, sv.Less(ValueOf("c")))
}

func TestValue_Int(t *testing.T) {
	nv := ValueOf(nil)
	iv := ValueOf(123)
	require.Equal(t, 123, iv.Int())
	require.Panics(t, func() { nv.Int() })
}

func TestValue_IndexValue(t *testing.T) {
	require.Nil(t, ValueOf(nil).PropertyValue(ValueOf("first")).Interface())
	require.Nil(t, ValueOf(false).PropertyValue(ValueOf("first")).Interface())
	require.Nil(t, ValueOf(12).PropertyValue(ValueOf("first")).Interface())

	empty := ValueOf([]string{})
	require.Equal(t, nil, empty.IndexValue(ValueOf(0)).Interface())
	require.Equal(t, nil, empty.IndexValue(ValueOf(-1)).Interface())

	av := ValueOf([]string{"first", "second", "third"})
	require.Equal(t, "first", av.IndexValue(ValueOf(0)).Interface())
	require.Equal(t, "third", av.IndexValue(ValueOf(-1)).Interface())
	require.Equal(t, "second", av.IndexValue(ValueOf(1.0)).Interface())
	require.Equal(t, "second", av.IndexValue(ValueOf(1.1)).Interface())

	hv := ValueOf(map[string]interface{}{"key": "value"})
	require.Equal(t, "value", hv.IndexValue(ValueOf("key")).Interface())
	require.Equal(t, nil, hv.IndexValue(ValueOf("missing_key")).Interface())

	hashPtr := ValueOf(&map[string]interface{}{"key": "value"})
	require.Equal(t, "value", hashPtr.IndexValue(ValueOf("key")).Interface())
	require.Equal(t, nil, hashPtr.IndexValue(ValueOf("missing_key")).Interface())
}

func TestValue_PropertyValue(t *testing.T) {
	empty := ValueOf([]string{})
	require.Equal(t, nil, empty.PropertyValue(ValueOf("first")).Interface())
	require.Equal(t, nil, empty.PropertyValue(ValueOf("last")).Interface())

	av := ValueOf([]string{"first", "second", "third"})
	require.Equal(t, "first", av.PropertyValue(ValueOf("first")).Interface())
	require.Equal(t, "third", av.PropertyValue(ValueOf("last")).Interface())
}

func TestValue_Contains(t *testing.T) {
	// array
	require.True(t, ValueOf([]int{1, 2}).Contains(ValueOf(2)))
	require.False(t, ValueOf([]int{1, 2}).Contains(ValueOf(3)))

	av := ValueOf([]string{"first", "second", "third"})
	require.True(t, av.Contains(ValueOf("first")))
	require.False(t, av.Contains(ValueOf("missing")))

	// string
	sv := ValueOf("seafood")
	require.True(t, sv.Contains(ValueOf("foo")))
	require.False(t, sv.Contains(ValueOf("bar")))

	// string contains stringifies its argument
	require.True(t, ValueOf("seaf00d").Contains(ValueOf(0)))

	// hash
	hv := ValueOf(map[string]interface{}{"key": "value"})
	require.True(t, hv.Contains(ValueOf("key")))
	require.False(t, hv.Contains(ValueOf("missing_key")))
}

func TestValue_PropertyValue_size(t *testing.T) {
	require.Nil(t, ValueOf(nil).PropertyValue(ValueOf("size")).Interface())
	require.Nil(t, ValueOf(false).PropertyValue(ValueOf("size")).Interface())
	require.Nil(t, ValueOf(12).PropertyValue(ValueOf("size")).Interface())

	require.Equal(t, 7, ValueOf("seafood").PropertyValue(ValueOf("size")).Interface())

	empty := ValueOf([]string{})
	require.Equal(t, 0, empty.PropertyValue(ValueOf("size")).Interface())

	av := ValueOf([]string{"first", "second", "third"})
	require.Equal(t, 3, av.PropertyValue(ValueOf("size")).Interface())

	hv := ValueOf(map[string]interface{}{"key": "value"})
	require.Equal(t, 1, hv.PropertyValue(ValueOf("size")).Interface())

	withSizeKey := ValueOf(map[string]interface{}{"size": "value"})
	require.Equal(t, "value", withSizeKey.IndexValue(ValueOf("size")).Interface())

	hashPtr := ValueOf(&map[string]interface{}{"key": "value"})
	require.Equal(t, 1, hashPtr.PropertyValue(ValueOf("size")).Interface())
}
