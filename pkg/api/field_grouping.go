package api

import "reflect"

/**
 * With filed grouping, the event are routed to downstream
 * instances by its value. This implementation has many limitations
 * as it only supports mapping events by string.
 */
type FieldGrouping struct {
	cnt int
	m   map[string]int
}

func NewFieldGrouping() *FieldGrouping {
	return &FieldGrouping{
		cnt: 0,
		m:   make(map[string]int),
	}
}

// Get key from an event. Child class can override this function to calculate key in different ways.
// For example, calculate the key from some specific fields.
func (f *FieldGrouping) GetKey(event Event) any {
	return valuesFromStruct(event)
}

// Get target instance id from an event and component parallelism.
func (f *FieldGrouping) GetInstance(event Event, parallelism int) int {
	s, ok := f.GetKey(event).(string)
	if !ok {
		panic("only support map for string currently")
	}

	val, ok := f.m[s]
	if !ok {
		f.m[s] = f.cnt
		f.cnt++
		return f.m[s] % parallelism
	}
	return val % parallelism
}

// helper function to convert the data of event into it's own type, but currently only supports string mapping in GetInstance
func valuesFromStruct(data any) []any {
	v := reflect.ValueOf(data)
	out := make([]any, 0)
	for i := 0; i < v.NumField(); i += 1 {
		field := v.Field(i)
		fieldType := field.Type()
		switch fieldType.Name() {
		case "int64":
			out = append(out, field.Interface().(int64))
			break
		case "float64":
			out = append(out, field.Interface().(float64))
			break
		case "string":
			out = append(out, field.Interface().(string))
			break
		// And all your other types (here) ...
		default:
			out = append(out, field.Interface())
			break
		}
	}
	return out
}
