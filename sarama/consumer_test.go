package sarama

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestConsumer(t *testing.T) {
	cfg := sarama.NewConfig()

	consumer, err := sarama.NewConsumerGroup(
		addrs, "test_consumer", cfg)
	//if err != nil {
	//	return
	//}
	require.NoError(t, err)

	err = consumer.Consume(context.Background(),
		[]string{"test_topic"},
		testConsumerGroupHandler{})
	t.Log(err)
}

type testConsumerGroupHandler struct {
}

func (t testConsumerGroupHandler) Setup(session sarama.ConsumerGroupSession) error {
	log.Println("setup")
	return nil
}

func (t testConsumerGroupHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	log.Println("cleanup")
	return nil
}

func (t testConsumerGroupHandler) ConsumeClaim(
	//从建立连接到链接彻底断掉那一段时间
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim) error {
	msgs := claim.Messages()
	for msg := range msgs {
		//var bizMsg MyBizMsg
		//err := json.Unmarshal(msg.Value, &bizMsg)
		//if err != nil {
		//	//这里消费消息出错
		//	//大多数时候是重试 且 记录日志
		//	continue
		//}

		log.Println(string(msg.Value))

		session.MarkMessage(msg, "")

	}

	return nil
}

type MyBizMsg struct {
	Name string
}
