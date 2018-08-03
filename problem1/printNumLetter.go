package problem1

import (
	"strconv"
	"sync"
)

func printNumLetter() string {
	res := []byte{}
	var wg sync.WaitGroup
	wg.Add(28)

	printNum, printLetter := make(chan bool), make(chan bool)

	go func() {
		for i := 1; i <= 28; i++ {
			<-printNum
			res = append(res, []byte(strconv.Itoa(i))...)
			wg.Done()

			i++

			res = append(res, []byte(strconv.Itoa(i))...)
			wg.Done()

			if i%2 == 0 {
				printLetter <- true
			}
		}
	}()

	go func() {
		for j := 65; j <= 90; j++ {
			<-printLetter
			res = append(res, byte(j))

			j++

			res = append(res, byte(j))

			if j%2 == 0 {
				printNum <- true
			}
		}
	}()

	printNum <- true

	wg.Wait()

	return string(res)
}
