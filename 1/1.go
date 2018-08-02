package problem1

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(28)

	printNum, printLetter := make(chan bool), make(chan bool)

	go func() {
		for i := 1; i <= 28; i++ {
			<-printNum
			fmt.Print(i)
			wg.Done()

			i++

			fmt.Print(i)
			wg.Done()

			if i%2 == 0 {
				printLetter <- true
			}
		}
	}()

	go func() {
		for j := 65; j <= 90; j++ {
			<-printLetter
			fmt.Print(string(byte(j)))

			j++

			fmt.Print(string(byte(j)))

			if j%2 == 0 {
				printNum <- true
			}
		}
	}()

	printNum <- true

	wg.Wait()
}
