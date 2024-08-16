package sarama

import (
	"github.com/IBM/sarama"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var addrs = []string{"localhost:9094"}

func TestSyncProducer(t *testing.T) {
	//config设置1
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true

	//config设置1
	//通过这个来设置使用何种分区算法
	cfg.Producer.Partitioner = sarama.NewHashPartitioner

	//创建一个生产者
	producer, err := sarama.NewSyncProducer(addrs, cfg)
	assert.NoError(t, err)

	//发送消息
	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: "first_topic",

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

func TestAsyncProducer(t *testing.T) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Errors = true
	cfg.Producer.Return.Successes = true
	//
	////通过这个来设置使用何种分区算法
	//cfg.Producer.Partitioner = sarama.NewHashPartitioner

	producer, err := sarama.NewAsyncProducer(addrs, cfg)
	assert.NoError(t, err)
	require.NoError(t, err)
	msgCh := producer.Input()
	msgCh <- &sarama.ProducerMessage{
		Topic: "test_topic",

		Key:   sarama.StringEncoder("oid-123"),
		Value: sarama.StringEncoder("hello, 这是一条消息"),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("trace_id"),
				Value: []byte("123456"),
			},
		},
		//只作用于发送过程
		Metadata: "这是metadata",
	}
	errCh := producer.Errors()
	succCh := producer.Successes()

	select {
	case err := <-errCh:
		t.Log("发送出问题了", err.Err)
	case <-succCh:
		t.Log("发送成功")
	}

}
