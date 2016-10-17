package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	//"github.com/tdhite/q3-training-journal/journal"
)

type SimpleConsumer struct {
}

func (sc *SimpleConsumer) consume(url string, topic string, listen chan bool) (<-chan *Message, error) {
	messages := make(chan *Message, 1)
	tl := fmt.Sprintf("%s/api/topic/%s", url, topic)

	go func() {
		for <-listen {
			fmt.Printf("Checking queue at %s for topic %s\n", url, topic)

			res, err := http.Get(tl)
			if err != nil {
				log.Print(err)
				continue
			}

			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				//log.Print(err)
				continue
			}

			//fmt.Printf("raw: %s", b)
			msg := &Message{}

			msg.FromJson(b)
			if msg.Base64 == nil || len(msg.Base64) == 0 {
				//log.Print("no message")
				res.Body.Close()
				continue
			}
			messages <- msg
			res.Body.Close()
			if err != nil {
				log.Fatal(err)
			}

		}
		fmt.Println("Exit consumer")

		close(messages)
	}()

	return messages, nil
}

func (sc *SimpleConsumer) ConsumeMessages(url string, topic string) error {

	listen := make(chan bool, 1)

	messages, _ := sc.consume(url, topic, listen)
	go func() {
		for msg := range messages {
			fmt.Printf("topic: %s", msg.ToJson())
			fmt.Println()
		}
	}()

	for {
		time.Sleep(time.Second * 15)
		listen <- true
	}

	listen <- false
	time.Sleep(time.Second * 3)
	return nil

}
