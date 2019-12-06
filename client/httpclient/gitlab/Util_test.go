package gitlab

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(URLParse("http://gitlab.puhuitech.cn/shizhiyin/temporary/raw/bf1c72c26a9e80a9ac41d6860df482cd52dfcefd/indicator/python/ZHJ.py"))

	fmt.Println(URLParse("http://gitlab.puhuitech.cn/lidongchen/FF_PB_workflow/raw/master/"))
	fmt.Println(URLParse("http://gitlab.puhuitech.cn/lidongchen/FF_PB_workflow/raw/master"))

	fmt.Println(URLParse("http://gitlab.puhuitech.cn/lidongchen/FF_PB_workflow/"))
	fmt.Println(URLParse("http://gitlab.puhuitech.cn/lidongchen/FF_PB_workflow"))
}
