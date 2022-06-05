package interrupt

import (
	"os"
	"sync/atomic"
	"time"

	"go-api-server-example/internal/logging"
)

var (
	running = true

	runningGoroutines int32
)

func applicationStop() {
	running = false
}

func IsApplicationReady() bool {
	return running
}

func IncreaseActiveTasksCounter() func() {
	atomic.AddInt32(&runningGoroutines, 1)

	return func() {
		DecreaseActiveTasksCounter()
	}
}

func DecreaseActiveTasksCounter() {
	atomic.AddInt32(&runningGoroutines, -1)
}

func GetCountRunningGoroutines() int32 {
	return atomic.LoadInt32(&runningGoroutines)
}

func WaitInterruptSignal(c chan os.Signal) {
	var (
		log = logging.GetLogger()
	)
	for val := range c {
		if IsApplicationReady() {
			applicationStop()
			go func() {
				for {
					if GetCountRunningGoroutines() == 0 {
						os.Exit(0)
					}

					log.Infof("Waiting for tasks to finish before closing the program. Running tasks: %d", GetCountRunningGoroutines())
					time.Sleep(time.Second)
				}
			}()
		}

		log.Infof("Got signal: %s. Running tasks: %d", val, GetCountRunningGoroutines())
	}
}
