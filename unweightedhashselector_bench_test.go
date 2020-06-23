package gocrush

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
)

func BenchmarkUnwweightedHashSelector(b *testing.B) {
	b.StopTimer()
	memStats := new(runtime.MemStats)
	runtime.GC()
	runtime.ReadMemStats(memStats)
	node := new(CrushNode)
	node.childrens = make([]CNode, 10000)
	counter := make(map[string]int)
	for i := 0; i < 10000; i++ {
		child := new(CrushNode)
		child.weight = 1
		child.id = "Child" + strconv.Itoa(i)
		node.childrens[i] = child
		counter[child.id] = 0
	}
	b.StartTimer()
	for x := 0; x < b.N; x++ {

		selector := NewHashingSelector(node)

		for i := int64(0); i < 100000; i++ {
			// Get replicants
			for r := int64(0); r < 3; r++ {
				nn := selector.Select(i, r)
				counter[nn.GetID()]++
			}
		}

	}
	b.StopTimer()
	selectMemStats := new(runtime.MemStats)
	runtime.GC()
	runtime.ReadMemStats(selectMemStats)
	fmt.Printf("\nBefore alloc: %v; After selection: %v\n",
		memStats.Alloc/1000, selectMemStats.Alloc/1000)
	//for key, nn := range counter {
	//	log.Printf("Node: %s - %d", key, nn)
	//}

}
