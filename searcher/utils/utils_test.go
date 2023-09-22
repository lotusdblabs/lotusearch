package utils

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestExecTimeWithNanoseconds(t *testing.T)  {
	fmt.Println("TestExecTime ",ExecTimeWithNanoseconds(func() {
		fmt.Println("handle function")
	}))
}

func TestExecTimeWithError(t *testing.T){
	_,err := ExecTimeWithError(func() error {
		return fmt.Errorf("error handle function")
	})
	assert.Equal(t,fmt.Errorf("error handle function"),err)
}

