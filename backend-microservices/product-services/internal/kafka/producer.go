package kafka

import (
	"log"
	"os"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	producer sarama.SyncProducer
}

func NewKafkaProducer() *KafkaProducer {
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "kafka-broker:9092"
	}

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalf("❌ Failed to start Kafka producer: %v", err)
	}

	log.Printf("✅ Kafka producer connected to %s", broker)
	return &KafkaProducer{producer: producer}
}

func (p *KafkaProducer) Publish(topic string, key, value []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.ByteEncoder(key),
		Value: sarama.ByteEncoder(value),
	}

	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		log.Printf("❌ Failed to send message: %v", err)
		return err
	}

	log.Printf("📤 Sent message to topic %s [partition: %d, offset: %d]", topic, partition, offset)
	return nil
}

func (p *KafkaProducer) Close() error {
	return p.producer.Close()
}
