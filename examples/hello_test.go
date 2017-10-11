package examples

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/chen56/retry"
	"time"
	"fmt"
	"errors"
)

func TestStop(t *testing.T) {
	assert := assert.New(t)
	assert.True(true)
	retryer := retry.New().
		WithWaitFixed(1 * time.Second).
		WithRetryMaxCount(3)
	try :=0
	result,err:=retryer.Run(func()error {
		fmt.Printf("try \n")
		try++
        if try ==2 {
        	//return nil
		}
		return errors.New("sql timeout")
	})
	fmt.Printf("result= %v \n", result)
	fmt.Printf("err= %v \n",err)
}

