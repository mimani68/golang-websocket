package log

import "fmt"

func Log(a ...interface{}) {
	fmt.Println(a[0])
}
