package retry

import (
	"time"
	"fmt"
	"errors"
)
type Retrer struct {
	canRetry func(attempt Attempt) error
	wait     func(attempt Attempt)
}

type Attempt struct {
	Count   int
	elapsed time.Duration
	LastErr error
}

func (this Retrer) WithWaitFixed(d time.Duration) Retrer {
	this.wait = func (attempt Attempt){
		time.Sleep(d)
	}
	return this
}

func (this Retrer) WithRetryMaxCount(max int) Retrer {
	this.canRetry = func(attempt Attempt) error {
		if attempt.Count < max {
			return nil
		} else {
			return errors.New(fmt.Sprintf("Cancel: RetryMaxCount: max(%d) is reached",max))
		}
	}
	return this
}

func (this Retrer) WithRetryForever() Retrer {
	this.canRetry = func(attempt Attempt) error {
		return nil
	}
	return this
}

func (this Retrer) Run(retryFunc func()(err error)) (attempt Attempt,err error) {
	start:=time.Now()

	for{
		attempt.LastErr = retryFunc()
        attempt.Count++
        attempt.elapsed=time.Now().Sub(start)

		//no err ,return
		if attempt.LastErr ==nil{
			return attempt,nil
		}

		//can not retry,return err
		if err = this.canRetry(attempt);err!=nil{
			return attempt,err
		}

		//wait
		this.wait(attempt)
	}
}

func New() Retrer {
	return Retrer{
		canRetry: func(attempt Attempt) error {
			return nil
		},
		wait: func(attempt Attempt) {},

	}
}