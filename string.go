package main

import "fmt"

func (i queueIndex) String() string {
	return fmt.Sprintf("lastI: %d | length: %d | entryP: %d | hwm: %d", i.lastInsert, i.length, i.entryPoint, i.highWatermark)
}
