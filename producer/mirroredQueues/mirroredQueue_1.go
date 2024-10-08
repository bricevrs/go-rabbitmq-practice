package producer_mirroredQueues

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func MirroredQueue1() {
	// Here we connect to RabbitMQ or send a message if there are any errors connecting.
	conn, err := amqp.Dial("amqp://guest:guest@haproxy:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	log.Println("Successfully connected to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// We create a quorum queue to send the message to.
	q, err := ch.QueueDeclare(
		"mirrored-queue-1", // name
		true,               // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare a queue")

	// We set the payload for the message.
	body := "Quorum Queues are awesome - Keep Moving Forward!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	// If there is an error publishing the message, a log will be displayed in the terminal.
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Congrats, sending message: %s", body)
}
