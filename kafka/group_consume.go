package kafka

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
	"go.uber.org/zap"
)

type GroupHandler struct {
	MessageFormat
	*zap.Logger
}

func NewGroupHandler(message MessageFormat, logger *zap.Logger) *GroupHandler {
	return &GroupHandler{
		MessageFormat: message,
		Logger:        logger,
	}
}

func (h GroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	h.Info("Setup")
	return nil
}
func (h GroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	h.Info("Cleanup")
	return nil
}
func (h GroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	h.Info("ConsumeClaim: ", zap.String("topic", claim.Topic()), zap.Int32("Partition", claim.Partition()))
	for msg := range claim.Messages() {
		h.Info("message", zap.Int64("offset", msg.Offset), zap.String("key", string(msg.Key)), zap.String("value", string(msg.Value)))
		session.MarkMessage(msg, "")
		h.Format(msg.Value)
		session.Commit()
	}
	return nil
}

type GroupConsume struct {
	Group sarama.ConsumerGroup
	*GroupHandler
}

func NewGroupConsume(brokers []string, gorupId string, message MessageFormat, logger *zap.Logger) *GroupConsume {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	group, err := sarama.NewConsumerGroup(brokers, gorupId, config)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for err := range group.Errors() {
			logger.Error("group error", zap.Error(err))
		}
	}()
	return &GroupConsume{
		Group:        group,
		GroupHandler: NewGroupHandler(message, logger),
	}
}

func (g *GroupConsume) Start(topic string) {
	defer g.Close()
	g.Info("start topic", zap.String("topic", topic))
	for {
		err := g.Group.Consume(context.Background(), []string{topic}, g.GroupHandler)
		if err != nil {
			g.Error("GroupConsume.Start.err", zap.Error(err))
			return
		}
	}
}

func (g *GroupConsume) Close() {
	if err := g.Group.Close(); err != nil {
		g.Error("GroupConsume.Close.err", zap.Error(err))
		log.Fatalln(err)
	}
}
