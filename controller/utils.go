package controller

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/cache/model"
	"github.com/streadway/amqp"
)

// Publish to queue
func (conn Conn) Publish(routingKey string, data []byte) (err error) {
	err = conn.Channel.Publish("events", routingKey, false, false, amqp.Publishing{ContentType: "application/json", Body: data, DeliveryMode: amqp.Persistent})

	return
}

// StartConsumer listen to queue
func (conn Conn) StartConsumer(queueName, routingKey string) error {
	log.Println("starting consumer")
	// create the queue if it doesn't already exist
	_, err := conn.Channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Println("Error while declaring queue:", err)
		return err
	}

	// bind the queue to the routing key
	err = conn.Channel.QueueBind(queueName, routingKey, "events", false, nil)
	log.Println("binding to queue")
	if err != nil {
		log.Println("Error while declaring queue:", err)
		return err
	}

	err = conn.Channel.Qos(1, 0, false)
	if err != nil {
		log.Println("Error while setting qos:", err)
		return err
	}

	msgs, err := conn.Channel.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		log.Println("Error while consume: ", err)
		return err
	}
	stopChan := make(chan bool)
	go func() {
		var data model.Data
		for msg := range msgs {
			err := json.Unmarshal(msg.Body, &data)
			log.Println("Msg recieved:", msg.Body)
			if err != nil {
				log.Printf("Error decoding JSON: %s", err)
			}
			_, _, err = insertInCacheFromDB(data)
			if err != nil {
				log.Println("Got err: ", err)
			}
			log.Println("Msg recieved: body: ", data)
			if err := msg.Ack(false); err != nil {
				log.Printf("Error acknowledging message : %s", err)
			} else {
				log.Printf("Acknowledged message")
			}
		}
	}()
	<-stopChan
	return nil
}

func insertIntoqueue(data model.Data) (err error) {
	log.Println("Func called: insertIntoqueue: ", data)
	var buffer bytes.Buffer
	enc := json.NewEncoder(&buffer)
	err = enc.Encode(data)
	if err != nil {
		log.Println(err)
		return
	}

	err = conn.Publish("key", buffer.Bytes())
	if err != nil {
		log.Println("unable to connect to mq. Err: ", err)
		return
	}
	return nil
}
