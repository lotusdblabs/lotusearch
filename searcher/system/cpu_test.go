package system

import (
	"fmt"
	"testing"

	"github.com/shirou/gopsutil/v3/cpu"
)

func TestCPU(t *testing.T) {
	fmt.Println(GetCPUStatus())
	c, _ := cpu.Info()
	fmt.Println(c)
}
