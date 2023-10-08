package command

type CommandState map[string]any

func GetStateData[T any](cs CommandState, key string) (T, bool) {
	var zeroVal T

	v, ok := cs[key]
	if !ok {
		return zeroVal, false
	}

	if result, ok := v.(T); ok {
		return result, true
	}

	return zeroVal, false
}

func SetStateData[T any](cs CommandState, key string, value T) {
	cs[key] = value
}
