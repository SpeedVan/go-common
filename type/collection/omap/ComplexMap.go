package omap

import (
	"bytes"
	"errors"
)

type ComplexMap map[string]interface{}

func (s ComplexMap) Get(key string) interface{} {
	return s.GetWithDefault(key, nil)
}

func (s ComplexMap) GetWithDefault(key string, def interface{}) interface{} {
	if v, ok := s[key]; ok {
		return v
	}
	return def
}

func (s ComplexMap) Set(key string, val interface{}) {
	s[key] = val
}

func (s ComplexMap) RecursionGet(keyStr string) (interface{}, error) {
	return s.RecursionGetByKeyArr(Split(keyStr))
}

func (s ComplexMap) RecursionSet(keyStr string, val interface{}) error {
	return s.RecursionSetByKeyArr(Split(keyStr), val)
}

func (s ComplexMap) RecursionGetByKeyArr(keys []string) (interface{}, error) {
	if len(keys) == 1 {
		return s[keys[0]], nil
	}
	if subMap, ok := s[keys[0]].(ComplexMap); ok {
		return subMap.RecursionGetByKeyArr(keys[1:])
	}
	return nil, errors.New("errorget:" + keys[0])
}

func (s ComplexMap) RecursionSetByKeyArr(keys []string, val interface{}) error {
	if len(keys) == 1 {
		s[keys[0]] = val
		return nil
	}
	if s[keys[0]] == nil {
		s[keys[0]] = make(ComplexMap)
	}
	if subMap, ok := s[keys[0]].(ComplexMap); ok {
		return subMap.RecursionSetByKeyArr(keys[1:], val)
	}
	return errors.New("errorset:" + keys[0])
}

func (s ComplexMap) ForEach(handler func(string, interface{})) {
	for k, v := range s {
		handler(k, v)
	}
}

func Split(str string) []string {
	result := []string{}
	buf := new(bytes.Buffer)
	hasChar, hasNumber := false, false
	for _, c := range []byte(str) {
		if c == '_' {
			l := len(result)
			if hasNumber && !hasChar {
				result = append(result, string(buf.Bytes()), "")
			} else {
				if l == 0 {
					result = append(result, string(buf.Bytes()))
				} else {
					if len(result[l-1]) == 0 {
						result[l-1] = string(buf.Bytes())
					} else {
						result[l-1] = result[l-1] + "_" + string(buf.Bytes())
					}

				}
			}
			buf.Reset()
			hasChar, hasNumber = false, false
			continue
		}
		if '0' <= c && c <= '9' {
			hasNumber = true
		} else {
			hasChar = true
		}
		buf.WriteByte(c)
	}
	l := len(result)
	if hasNumber && !hasChar {
		result = append(result, string(buf.Bytes()))
	} else {
		if l == 0 {
			result = append(result, string(buf.Bytes()))
		} else {
			if len(result[l-1]) == 0 {
				result[l-1] = string(buf.Bytes())
			} else {
				result[l-1] = result[l-1] + "_" + string(buf.Bytes())
			}

		}
	}
	return result
}
