package srcinfo

import (
	"bufio"
	"io"
	"reflect"
	"strings"
)

func ParseSrcInfo(r io.Reader) (*SrcInfo, error) {
	var srcinfo SrcInfo
	var current *applier
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		trimmed := strings.TrimSpace(line)
		keyValue := strings.SplitN(trimmed, "=", 2)
		if len(keyValue) < 2 {
			continue
		}
		key := strings.TrimSpace(keyValue[0])
		value := strings.TrimSpace(keyValue[1])
		switch key {
		case "pkgbase":
			current = newApplier(&srcinfo.Global)
		case "pkgname":
			srcinfo.Packages = append(srcinfo.Packages, Package{})
			current = newApplier(&srcinfo.Packages[len(srcinfo.Packages)-1])
		}
		current.apply(key, value)
	}
	if sc.Err() != nil {
		return nil, sc.Err()
	}
	return &srcinfo, nil
}

type applier struct {
	fields map[string]reflect.Value
}

func newApplier(v interface{}) *applier {
	a := &applier{
		fields: make(map[string]reflect.Value),
	}
	a.discoverFields(reflect.ValueOf(v).Elem()) // is pointer
	return a
}

func (a *applier) discoverFields(r reflect.Value) {
	t := r.Type()
	for fn := 0; fn < t.NumField(); fn++ {
		if t.Field(fn).Anonymous {
			a.discoverFields(r.Field(fn))
		} else {
			a.fields[strings.ToLower(t.Field(fn).Name)] = r.Field(fn)
		}
	}
}

func (a *applier) apply(key, value string) {
	field, ok := a.fields[key]
	if ok {
		switch {
		case field.Type().Kind() == reflect.String:
			field.SetString(value)
		case field.Type().Kind() == reflect.Slice:
			switch {
			case field.Type().Elem().Kind() == reflect.String:
				field.Set(reflect.Append(field, reflect.ValueOf(value)))
			}
		}
	}
}
