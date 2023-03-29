package kafka

import (
	"log"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type MessageFormat interface {
	Format(message []byte)
}

type Consume struct {
	partition sarama.PartitionConsumer
	Topic     string
	Consumer  sarama.Consumer
	MessageFormat
	*zap.Logger
}

func NewConsume(Topic string, brokers []string, Message MessageFormat, logger *zap.Logger) *Consume {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	Consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalln("sarama.NewConsume:err", err)
	}
	return &Consume{
		Topic:         Topic,
		Consumer:      Consumer,
		MessageFormat: Message,
		Logger:        logger,
	}
}

func (c *Consume) Start() {
	var err error
	c.partition, err = c.Consumer.ConsumePartition(c.Topic, 0, sarama.OffsetOldest)
	if err != nil {
		c.Error("sarama.NewConsume:err", zap.Error(err))
		panic(err)
	}
	for {
		select {
		case msg := <-c.partition.Messages():
			c.Info("message", zap.Int64("offset", msg.Offset), zap.String("key", string(msg.Key)), zap.String("value", string(msg.Value)))
			c.Format(msg.Value)
		case err := <-c.partition.Errors():
			c.Error("Consume:err", zap.Error(err))
		}
	}
}
func (c *Consume) Close() {
	if c.partition != nil {
		if err := c.partition.Close(); err != nil {
			c.Error("Consume:err", zap.Error(err))
			log.Fatalln(err)
		}
	}
}
