package file

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test(t *testing.T) {

	maxRound := 5

	maxTry := 7 * maxRound
	for i := 0; i < maxTry; i = i % 7 {
		lockedFile := "/Users/admin/mnt/10-121-128-101/lock/version1/" + strconv.Itoa(i)
		fmt.Println("try lockfile:" + lockedFile)
		flock, err := New(lockedFile)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer flock.Unlock()
		err = flock.Lock()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("get file lock:" + strconv.Itoa(i))
			break
		}
		i++
	}
	fmt.Println("lock")
	// wg := sync.WaitGroup{}

	// for i := 0; i < 10; i++ {
	//     wg.Add(1)
	//     go func(num int) {
	//         flock := New(lockedFile)
	//         err := flock.Lock()
	//         if err != nil {
	//             wg.Done()
	//             fmt.Println(err.Error())
	//             return
	//         }
	//         fmt.Printf("output : %d\n", num)
	//         wg.Done()
	//     }(i)
	// }
	// wg.Wait()
	time.Sleep(30 * time.Second)
}
