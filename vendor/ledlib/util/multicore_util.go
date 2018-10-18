package util

import (
	"runtime"
	"sync"
)

type EnumCallback func(i int)

func ConcurrentEnum(start, end int, callback EnumCallback) {
	var wg sync.WaitGroup
	usingCore := runtime.NumCPU()
	wg.Add(usingCore)
	xloop := func(start, end int) {
		defer wg.Done()
		for i := start; i < end; i++ {
			callback(i)
		}
	}

	work := (end - start) / usingCore
	for c := start; c < usingCore; c++ {
		if c == usingCore-1 {
			go xloop(start+c*work, end)
		} else {
			go xloop(start+c*work, (c+1)*work)
		}
	}
	wg.Wait()
}
