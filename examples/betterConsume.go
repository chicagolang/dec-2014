package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	//START1 OMIT
	url := "amqp://guest:guest@" + "192.168.59.103" + ":5672"
	connection, err := amqp.Dial(url)
	if err != nil {
		panic(err.Error())
	}
	defer connection.Close()

	//END1 OMIT

	//START2 OMIT
	channel, err1 := connection.Channel()
	if err1 != nil {
		panic(err1.Error())
	}
	defer channel.Close()
	//END2 OMIT

	//START3 OMIT
	name := "TransactionFirehose"
	durable := true
	autoDelete := false
	exclusive := false
	noWait := false
	queue, err2 := channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, nil)
	if err2 != nil {
		panic(err2.Error())
	}
	//END3 OMIT

	//START4 OMIT
	consumer := ""
	autoAck := false // HL
	cons_exclusive := false
	noLocal := false
	cons_noWait := false
	msgs, err3 := channel.Consume(queue.Name, consumer, autoAck, cons_exclusive, noLocal, cons_noWait, nil)
	if err3 != nil {
		panic(err3.Error())
	}

	for d := range msgs {
		log.Printf("Got a message: %s", d.Body)
		d.Ack(false) // HL
	}
	//END4 OMIT

}
