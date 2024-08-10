package pokectx

import (
	"strconv"
)

type Te4nickPokeCTX struct {
	root *te4nickNode
}

func New(path ...string) *Te4nickPokeCTX {
	ctx := &Te4nickPokeCTX{
		root: &te4nickNode{
			children: make(map[string]*te4nickNode),
		},
	}

	ctx.Set(path...)
	return ctx
}

// Set last provided string value to the path specified before it
//
// ctx.Set("my", "string", "value")
//
// ctx.Get("my", "string") // -> "value"
func (ctx *Te4nickPokeCTX) Set(path ...string) {
	ctx.root.set(path...)
}

// Alias to Set(path, "")
func (ctx *Te4nickPokeCTX) SetKey(path ...string) {
	ctx.root.set(append(path, "")...)
}

// Get value by specified path
//
// ctx.Set("my", "string", "value")
//
// ctx.Get("my", "string") // -> "value"
func (ctx *Te4nickPokeCTX) Get(path ...string) (value string, found bool) {
	return ctx.root.get(path...)
}

// Get all available keys after specified path
//
// ctx.SetKey("my", "string", "value1")
//
// ctx.SetKey("my", "string", "value2")
//
// ctx.GetKeys("my", "string") // -> []string{"value1", "value2"}
func (ctx *Te4nickPokeCTX) GetKeys(path ...string) (value []string, found bool) {
	return ctx.root.getChildrenNames(path...)
}

type Ints interface {
	int | int8 | int16 | int32 | int64
}

type Uints interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Floats interface {
	float32 | float64
}

type Number interface {
	Ints | Uints | Floats
}

func GetOrDefaultNum[T Number](ctx *Te4nickPokeCTX, defaultVar T, path ...string) T {
	s, found := ctx.Get(path...)
	if !found {
		return defaultVar
	}

	switch interface{}(defaultVar).(type) {
	case int, int8, int16, int32, int64:
		u, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return defaultVar
		}

		return T(u)
	case uint, uint8, uint16, uint32, uint64:
		u, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return defaultVar
		}

		return T(u)
	case float32, float64:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return defaultVar
		}

		return T(f)
	default:
		return defaultVar
	}
}

func SetNum[T Number](ctx *Te4nickPokeCTX, number T, path ...string) {
	var s string

	switch interface{}(number).(type) {
	case int, int8, int16, int32, int64:
		s = strconv.FormatInt(int64(number), 10)
	case uint, uint8, uint16, uint32, uint64:
		s = strconv.FormatUint(uint64(number), 10)
	case float32, float64:
		s = strconv.FormatFloat(float64(number), 'f', -1, 64)
	}

	path = append(path, s)
	ctx.Set(path...)
}
