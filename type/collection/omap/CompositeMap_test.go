package omap_test

import (
	"fmt"
	"testing"

	"github.com/SpeedVan/go-common/type/collection/omap"
)

func TestCompositeMap(t *testing.T) {
	config := &omap.CompositeMap{
		Data: make(omap.ComplexMap),
	}

	

	check(config)

	v := config.Get("123")
	fmt.Println(v)

	err := config.RecursionSet("A123BC_0_123ads_31_sfs_sa12d_33", 123)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(config)

	val, err := config.RecursionGet("Avcsasd_fdsafs_9_dsfas")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	val, err = config.RecursionGet("A123BC_0_123ads_31_sfs_sa12d")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}

	if m, ok := val.(*omap.CompositeMap); ok {
		fmt.Println(m.ToStringMap())
	}
	
}

func check(m omap.Map) omap.Map {
	return m
}
