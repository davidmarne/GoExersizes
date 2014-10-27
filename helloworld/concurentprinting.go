import(
    "fmt"
)

responses := [5]string{"hello","hi","how are you","good you?","good thanks for asking"}

func chPrinter(ch1<-chan string, ch2 chan<-string) {
	for {
		msg, more := <-ch2
		indx := responses.IndexOf(msg)
		fmt.Println(responses[indx])
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go chPrinter(ch1, ch2)
	go chPrinter(ch2, ch1)

	ch1 <- 'hello'
}