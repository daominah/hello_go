package main

import (
	"log"
	"time"
	"unsafe"
)

// this example accesses struct private fields
func NewDebugTime(wall uint64, ext int64, loc *time.Location) time.Time {
	ret := time.Time{}
	/* definition of Time struct in standard library "time"
	type Time struct {
		wall uint64
		ext  int64
		loc *Location
	}
	*/
	// since structs are organized in memory order, we can advance the pointer
	// by field size until we're at the desired member, int is 8 bytes.
	// if you wanted to alter the 1st field, you wouldn't advance the pointer
	// at all
	ptrToT0 := unsafe.Pointer(&ret)
	ptrToWall := (*uint64)(ptrToT0)
	*ptrToWall = wall
	ptrToT0 = unsafe.Pointer(uintptr(ptrToT0) + uintptr(8))
	ptrToExt := (*int64)(ptrToT0)
	*ptrToExt = ext
	ptrToT0 = unsafe.Pointer(uintptr(ptrToT0) + uintptr(8))
	ptrToLoc := (**time.Location)(ptrToT0)
	*ptrToLoc = loc
	return ret
}

func main() {
	now := NewDebugTime(0x0, 63740578600, nil)
	log.Printf("real: %v", now.Format(time.RFC3339))
	expected, _ := time.Parse(time.RFC3339, "2020-11-10T04:16:40Z")
	log.Printf("expected: %#v", expected)
}
