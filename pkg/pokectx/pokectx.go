package pokectx

import (
	"strconv"
)

type PokeCTX struct {
	root *node
}

func New(path ...string) *PokeCTX {
	ctx := &PokeCTX{
		root: &node{
			children: make(map[string]*node),
		},
	}

	ctx.Set(path...)
	return ctx
}

func (ctx *PokeCTX) Set(path ...string) {
	ctx.root.set(path...)
}

func (ctx *PokeCTX) Get(path ...string) (string, bool) {
	return ctx.root.get(path...)
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

func GetDefaultNum[T Number](ctx *PokeCTX, defaultVar T, path ...string) T {
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

func SetNum[T Number](ctx *PokeCTX, number T, path ...string) {
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
