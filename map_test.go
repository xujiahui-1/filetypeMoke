package filetypeMoke

import (
	"fmt"
	"sync"
	"testing"
)

var tt sync.Map

func TestMap(t *testing.T) {
	tt.Store("gou", "xiaobai")

	tt.Store("gou", "xiaohong")
	tt.Store("gou", "xiaohuang")
	fmt.Println(tt.Load("gou"))
}
