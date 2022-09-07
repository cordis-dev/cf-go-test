package abomaB

import "sync/atomic"

func AtomicTests2() {
	x := uint64(1)
	x = atomic.AddUint64(&x, 1)        // ERROR "direct assignment to atomic value"
}
