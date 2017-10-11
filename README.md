# go-retry

## go-retry hello world

retry three times, and wait 1 second.

```bash
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
		WithRetryMaxCount(3)

	i :=0
	result,err:=retryer.Run(func()error {
		fmt.Printf("try %d \n", i)
		i++
		return errors.New(fmt.Sprintf("sql timeout: %d", i))
	})
	fmt.Printf("result= %+v \n", result)
	fmt.Printf("err= %v \n",err)
}
```

output:

```bash
try 0
try 1
try 2
result= {Count:3 Elapsed:2.007613586s LastErr:sql timeout 3}
err= Cancel: RetryMaxCount: max(3) is reached
```

## examples

<https://github.com/chen56/go-retry/tree/master/examples>

