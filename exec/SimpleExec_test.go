package exec

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	fmt.Println(SimpleExec("sh", "login_k8s_fg.sh"))
	fmt.Println(SimpleExec("kube", "get", "pods", "-n", "fission"))
}
