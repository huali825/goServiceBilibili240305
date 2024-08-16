package sarama2

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var addrs2 = []string{"localhost:9094"}
var addrs3 = []string{"localhost:9094", "localhost:9095"}

type ConsumerHanlder struct {
}

func (c ConsumerHanlder) Setup(session sarama.ConsumerGroupSession) error {
	log.Println("handler setup")
	return nil
}

func (c ConsumerHanlder) Cleanup(session sarama.ConsumerGroupSession) error {
	log.Println("handler cleanup")
	return nil
}

// ConsumeClaim 消费者 分区
// 打印?
func (c ConsumerHanlder) ConsumeClaim(
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim) error {

	//从channel 接收的消息
	ch := claim.Messages()
	for msg := range ch {
		log.Println(string(msg.Value))

		//标记已经消费了msg
		session.MarkMessage(msg, "")
	}
	return nil
}

func TestConsumer2(t *testing.T) {
	cfgC2 := sarama.NewConfig() //新建cfg

	//新建消费者组
	cg2, err := sarama.NewConsumerGroup(addrs2,
		"test_group2", cfgC2)
	assert.NoError(t, err)

	//新建ctx
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Consumer加入给定主题列表的消费者集群，
	// 并通过 ConsumerGroupHandler 启动一个阻塞的ConsumerGroupSession。
	err = cg2.Consume(ctx,
		[]string{"test_topic2"}, &ConsumerHanlder{})
	assert.NoError(t, err)
}
