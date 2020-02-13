package orderedmap

import "sort"

type Map map[sort.Interface]interface{}

type OrderedMap struct {
	Map
}
