package dowork

import (
	"context"
	"log"
	"time"
)

func DoWork(ctx context.Context, resChan chan int) {
	log.Println("[doWork] launch the doWork")
	sum := 0
	for {
		log.Println("[doWork] one iteration")
		time.Sleep(time.Millisecond)
		select {
		case <-ctx.Done():
			log.Println("[doWork] ctx Done is received inside doWork")
			return
		default:
			sum++
			if sum > 1000 {
				log.Println("[doWork] sum has reached 1000")
				resChan <- sum
				return
			}
		}
	}
}
