package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), 0); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func printTab(buf *bytes.Buffer, n int) {
	for i := 0; i < n; i++ {
		fmt.Fprintf(buf, "%s", "\t")
	}
}

func encode(buf *bytes.Buffer, v reflect.Value, tabCount int) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")
	case reflect.Bool:
		if v.Bool() {
			buf.WriteString("true")
		} else {
			buf.WriteString("false")
		}

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Complex64, reflect.Complex128:
		c := v.Complex()
		fmt.Fprintf(buf, "#C(%f %f)", real(c), imag(c))

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), tabCount)

	case reflect.Array, reflect.Slice: // (value ...)
		tabCount++
		buf.WriteString("[\n")
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteString(",\n")
			}
			printTab(buf, tabCount)
			if err := encode(buf, v.Index(i), tabCount); err != nil {
				return err
			}
		}
		tabCount--
		buf.WriteByte('\n')
		printTab(buf, tabCount)
		buf.WriteByte(']')

	case reflect.Struct: // ((name value) ...)
		tabCount++
		buf.WriteString("{\n")
		isFirst := true
		for i := 0; i < v.NumField(); i++ {
			if isNil(v.Field(i)) {
				continue
			}
			if !isFirst {
				buf.WriteString(",\n")
			}
			isFirst = false
			printTab(buf, tabCount)
			fmt.Fprintf(buf, "\"%s\": ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), tabCount); err != nil {
				return err
			}
		}

		buf.WriteString("\n}")
		tabCount--

	case reflect.Map: // ((key value) ...)
		tabCount++
		buf.WriteString("{\n")
		isFirst := true
		for _, key := range v.MapKeys() {
			if isNil(v.MapIndex(key)) {
				continue
			}
			if !isFirst {
				buf.WriteString(",\n")
			}
			isFirst = false
			printTab(buf, tabCount)
			if err := encode(buf, key, tabCount); err != nil {
				return err
			}
			buf.WriteString(": ")
			if err := encode(buf, v.MapIndex(key), tabCount); err != nil {
				return err
			}
		}
		tabCount--
		buf.WriteByte('\n')
		printTab(buf, tabCount)
		buf.WriteByte('}')

	case reflect.Interface:
		// TODO
	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Map, reflect.Slice, reflect.Interface:
		return v.IsNil()
	default:
		return false
	}
}
