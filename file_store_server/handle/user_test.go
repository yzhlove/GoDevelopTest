package handle

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_genToken(t *testing.T) {
	ts := fmt.Sprintf("%x", time.Now().Unix())
	fmt.Println(ts, " - ", time.Now().Unix())

	fmt.Println("tns => ", fmt.Sprintf("%x", time.Now().UnixNano()))

	nt, _ := strconv.ParseInt(ts, 16, 64)
	fmt.Println("nt = ", nt)

	a := "1235d99d1f4"
	b := a[3:]
	fmt.Println("b = ", b)

}
