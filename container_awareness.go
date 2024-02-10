package go_container_awareness

import (
	_ "github.com/KimMachineGun/automemlimit"
	_ "go.uber.org/automaxprocs"
	"runtime/debug"
)

func init() {
	debug.SetGCPercent(-1)
}
