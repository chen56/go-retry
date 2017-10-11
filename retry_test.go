package retry

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"time"
	"fmt"
	"errors"
)

func Test_RetryMaxCount(t *testing.T) {
	assert := assert.New(t)
	retryer := New().
		WithWaitFixed(1 * time.Millisecond).
		WithRetryMaxCount(3)

    i:=0
	retry,err:=retryer.Run(func()error {
		i++
		return errors.New(fmt.Sprintf("sql timeout:%v",i))
	})

	assert.Equal("Cancel: RetryMaxCount: max(3) is reached",err.Error())
	assert.Equal(3,retry.Count)
	assert.Equal("sql timeout:3",retry.LastErr.Error())
	//fmt.Println(retry.elapsed)
	assert.Equal(true,retry.elapsed>=(3-1) * time.Millisecond)

}

func Ignore_Test_StopNever(t *testing.T) {
	retryer := New().
		WithWaitFixed(1 * time.Second).
		WithRetryForever()

	i:=0
	retryer.Run(func()error {
		i++
		fmt.Printf("Test_StopNever.try %d \n",i)
		return errors.New(fmt.Sprintf("sql timeout:%v",i))
	})
}