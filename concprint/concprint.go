package main

import(
    "fmt"
)

var responses = []string{"hello","hi","how are you","good you?","good thanks for asking", "np", "bye","cya"}
var done = false

func chPrinter(chName string, ch1<-chan string, ch2 chan<-string, chDone chan<-bool) {
	for {
		msg, more := <-ch1
		
		if more == false{
			if done == false {
				done = true
				close(ch2)
				return
			}else {
				chDone <- true
				return 
			}
		}

		indx := indexOf(responses, msg)

		fmt.Println(chName, "says" ,responses[indx])

		if indx < len(responses) - 1{
			ch2 <- responses[indx + 1]
		} else {
			close(ch2)
		}
	}
}

func indexOf(arr [] string, toFind string) int{
	for i := 0; i < len(arr); i++ {
		if arr[i] == toFind {
			return i
		}
	}
	return -1
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	chDone := make(chan bool)

	go chPrinter("channel 1",ch1, ch2, chDone)
	go chPrinter("channel 2",ch2, ch1, chDone)

	ch1 <- "hello"
	<-chDone
}