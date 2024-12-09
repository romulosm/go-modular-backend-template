package messaging

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

func NewRabbitMQConnection() (*amqp.Connection, error) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar ao RabbitMQ: %w", err)
	}
	return conn, nil
}

func PublishMessage(conn *amqp.Connection, queueName, message string) error {
	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("falha ao abrir um canal: %w", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // nome
		false,     // durável
		false,     // excluir quando não utilizada
		false,     // exclusiva
		false,     // no-wait
		nil,       // argumentos
	)
	if err != nil {
		return fmt.Errorf("falha ao declarar uma fila: %w", err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return fmt.Errorf("falha ao publicar uma mensagem: %w", err)
	}

	return nil
}
