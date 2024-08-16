package sarama2

import (
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSyncProducer(t *testing.T) {
	//config设置1
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true

	//config设置1
	//通过这个来设置使用何种分区算法
	cfg.Producer.Partitioner = sarama.NewHashPartitioner

	//创建一个生产者
	producer, err := sarama.NewSyncProducer(addrs2, cfg)
	assert.NoError(t, err)

	//发送消息
	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: "test_topic2",

		//消息本体:
		Value: sarama.StringEncoder("hello, 这是一条消息"),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("trace_id"),
				Value: []byte("123456"),
			},
		},
		//只作用于发送过程
		Metadata: "这是metadata",
	})
	assert.NoError(t, err)

}
