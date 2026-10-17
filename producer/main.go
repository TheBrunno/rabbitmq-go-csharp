package main

import (
	"fmt"
	"log"
	"time"
	"math/rand/v2"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	// conectando ao rabbitmq
	conn, err := amqp.Dial("amqp://admin:admin@rabbitmq:5672/")
	if err != nil {
		log.Fatal(err) // log se der erro
	}
	defer conn.Close()

	fmt.Println("conectado ao rabbitmq com sucesso.")

	// criando um canal de comunicação
	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err) // log se der erro
	}
	defer channel.Close()

	// criando fila se ela nao existir
	queueName := "fila-pizzas"
	durable := true
	autoDelete := false
	exclusive := false
	noWait := false

	queue, err := channel.QueueDeclare(
		queueName,
		durable,
		autoDelete,
		exclusive,
		noWait,
		nil,
	)

	// enviar uma pizza a cada 10 segundos

	pizzas := []string{"Portuguesa", "Calabresa", "Frango com Catupiry", "Quatro Queijos", "Pepperoni", "Marguerita", "Napolitana", "Vegetariana", "Havaiana", "Baiana"}
	for {

		mensagem := pizzas[rand.IntN(len(pizzas))]

		exchange := ""
		routingKey := queue.Name
		mandatory := false
		immediate := false

		contentType := "text/plain"
		body := []byte(mensagem)

		err = channel.Publish(
			exchange,
			routingKey,
			mandatory,
			immediate,
			amqp.Publishing{
				ContentType: contentType,
				Body:        body,
			},
		)
		
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("pizza enviada:", mensagem)

		// espera 10 segundos antes de enviar a próxima pizza
		time.Sleep(10 * time.Second)
	}
}