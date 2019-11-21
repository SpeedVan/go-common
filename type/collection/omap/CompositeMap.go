package omap

import (
	"errors"
	"fmt"
)

type CompositeMap struct {
	Data ComplexMap
	Map
}

func New() Map {
	return &CompositeMap{
		Data: make(ComplexMap),
	}
}

func (s *CompositeMap) Get(key string) interface{} {
	return s.GetWithDefault(key, nil)
}

func (s *CompositeMap) GetWithDefault(key string, def interface{}) interface{} {
	if v, ok := s.Data[key]; ok {
		return v
	}
	return def
}

func (s *CompositeMap) Set(key string, val interface{}) Map {
	s.Data[key] = val
	return s
}

func (s *CompositeMap) RecursionGet(keyStr string) (interface{}, error) {
	return s.RecursionGetByKeyArr(Split(keyStr))
}

func (s *CompositeMap) RecursionSet(keyStr string, val interface{}) error {
	return s.RecursionSetByKeyArr(Split(keyStr), val)
}

func (s *CompositeMap) RecursionGetByKeyArr(keys []string) (interface{}, error) {
	if len(keys) == 1 {
		return s.Get(keys[0]), nil
	}
	if subMap, ok := s.Get(keys[0]).(*CompositeMap); ok {
		return subMap.RecursionGetByKeyArr(keys[1:])
	}
	return nil, errors.New("errorget:" + keys[0])
}

func (s *CompositeMap) RecursionSetByKeyArr(keys []string, val interface{}) error {
	if len(keys) == 1 {
		s.Set(keys[0], val)
		return nil
	}
	if s.Get(keys[0]) == nil {
		s.Set(keys[0], &CompositeMap{
			Data: make(ComplexMap),
		})
	}
	if subMap, ok := s.Get(keys[0]).(*CompositeMap); ok {
		return subMap.RecursionSetByKeyArr(keys[1:], val)
	}
	return errors.New("errorset:" + keys[0])
}

func (s *CompositeMap) Keys() []string {
	result := []string{}
	for k := range s.Data {
		result = append(result, k)
	}
	return result
}

func (s *CompositeMap) Values() []interface{} {
	result := []interface{}{}
	for _, v := range s.Data {
		result = append(result, v)
	}
	return result
}

func (s *CompositeMap) String() string {
	return fmt.Sprint(s.Data)
}

func (s *CompositeMap) ForEach(handler func(string, interface{})) {
	s.Data.ForEach(handler)
}

func (s *CompositeMap) F_map(handlerRet func(string, interface{}) (string, interface{})) Map {
	result := New()
	s.Data.ForEach(func(k string, v interface{}) {
		result.Set(handlerRet(k, v))
	})
	return s
}

func (s *CompositeMap) ToSourceMap() map[string]interface{} {
	return s.Data
}

func (s *CompositeMap) ToStringMap() map[string]string {
	result := make(map[string]string)
	for k, v := range s.Data {
		result[k]= fmt.Sprintf("%v", v)
	}
	return result
}