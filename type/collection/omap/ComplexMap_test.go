package omap_test

import (
	"fmt"
	"testing"

	"github.com/SpeedVan/go-common/type/collection/omap"
)

func Test(t *testing.T) {
	config := make(omap.ComplexMap)

	if v, ok := config["123"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("not exist")
	}

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

	val, err = config.RecursionGet("A123BC_0_123ads_31_sfs_sa12d_33")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
}

func TestRegex(t *testing.T) {

	fmt.Println(omap.Split("A123BC_0_123ads_31_sfs_sa12d_33"))

	fmt.Println(omap.Split("A123BC_0_123ads_31_sfs_sa12d_33_3a"))

	fmt.Println(omap.Split("abc_CHECK_CONFIG"))
}
