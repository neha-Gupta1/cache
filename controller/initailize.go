package controller

import (
	"fmt"
	"log"
	"os"

	"github.com/cache/model"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

//Conn struct
type Conn struct {
	Channel *amqp.Channel
}

var (
	conn   Conn
	server model.Server
	bucket map[int]string
)

// InitializeApp to start application
func InitializeApp() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		log.Println("We are getting the env values")
	}
	initializeBucket()
	err = initalizeQueue()
	if err != nil {
		log.Println("Error found:", err)
	}
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

}

func initializeBucket() {
	bucket = make(map[int]string, 1)
}

func initalizeQueue() (err error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s", os.Getenv("AMQP_USERNAME"), os.Getenv("AMQP_PASSWORD"), os.Getenv("AMQP_URL"), os.Getenv("AMQP_PORT"))
	connection, err := amqp.Dial(url)
	if err != nil {
		log.Println("Error found:", err)
		return err
	}
	ch, err := connection.Channel()
	conn = Conn{Channel: ch}
	if err == nil {
		log.Println("Inialization successfull")
	}
	return err

}

// RunConsumer strart another thread
func RunConsumer() {
	err := conn.StartConsumer("add", "key")
	if err != nil {
		log.Println("err while connecting: ", err)
		return
	}
}
