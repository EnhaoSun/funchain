package funchain

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var mu sync.Mutex
var m = make(map[uint64]int)
var t = make(map[uint64]string)

var fontColor = []int{31, 32, 33, 34, 35, 36, 37}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func printTrace(id uint64, name, typ string, indent int, c1, c2 int) string {
	indents := ""
	for i := 0; i < indent; i++ {
		indents += "\t"
	}
	gid := fmt.Sprintf("g[%02d]:", id)
	gtrace := fmt.Sprintf("%s%s%s", indents, typ, name)
	//fmt.Printf("\033[1;%dm%s\033[0m %c[0;40;%dm%s%c[0m\n", c1, gid, 0x1B, c2, gtrace, 0x1B)
	return fmt.Sprintf("\033[1;%dm%s\033[0m %c[0;40;%dm%s%c[0m\n", c1, gid, 0x1B, c2, gtrace, 0x1B)
}

func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	// 为当前的goroutine获取一个id
	id := getGID()
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	// id号打印的颜色
	f1 := 37
	if id%2 == 0 {
		f1 = 35
	}

	// 在同一个协程当中，v一直递增
	mu.Lock()
	v := m[id]
	m[id] = v + 1
	// 函数名字打印的颜色
	f2 := fontColor[v%(len(fontColor))]

	// 函数入口打印
	t[id] += printTrace(id, name, "->", v+1, f1, f2)
	mu.Unlock()

	// 函数出口打印
	return func() {
		mu.Lock()
		v := m[id]
		m[id] = v - 1
		t[id] += printTrace(id, name, "<-", v, f1, f2)
		mu.Unlock()
		if v == 1 {
			// 当函数调用完毕，打印整体的
			fmt.Println(t[id])
		}
	}
}
