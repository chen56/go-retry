package main

import (
	"github.com/chen56/go-retry"
	"time"
	"fmt"
	"errors"
)

func main() {
	retryer := retry.NewRetryer().
		WithWaitFixed(1 * time.Second).
		WithRetryForever()

	i :=0
	result,err:=retryer.Run(func()error {
		fmt.Printf("try %d \n", i)
		i++
		return errors.New(fmt.Sprintf("sql timeout %d", i))
	})
	fmt.Printf("result= %#v \n", result)
	fmt.Printf("err= %v \n",err)
}
