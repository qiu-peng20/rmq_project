package rabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"time"
)

const mqUrl = "amqp://guest:guest@127.0.0.1:5672/"

type RabbitMQ struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key : simple用不到
	Key   string
	MqUrl string
}

//实例化RabbitMQ
func NewRabbitMQ(queueName, exChange, key string) *RabbitMQ {
	rabbitMQ := &RabbitMQ{QueueName: queueName, Exchange: exChange, Key: key, MqUrl: mqUrl}
	var err error
	// 创建rabbitMQ 连接
	rabbitMQ.connection, err = amqp.Dial(rabbitMQ.MqUrl)
	rabbitMQ.FailError(err, "创建链接错误")
	rabbitMQ.channel, err = rabbitMQ.connection.Channel()
	rabbitMQ.FailError(err, "获取channel失败")
	return rabbitMQ
}

//停止RabbitMQ
func (r *RabbitMQ) StopRabbitMQ() {
	_ = r.connection.Close()
	_ = r.channel.Close()
}

//错误处理函数
func (r *RabbitMQ) FailError(err error, message string) {
	if err != nil {
		log.Fatalf("%s, %s", err, message)
	}
}

//简单模式：Step1 简单模式下创建rabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

// Step2 简单模式下生产代码
func (r *RabbitMQ) PublishSimple(message string) {
	// 1 申请队列，如果队列不存在则会创建队列
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//控制消息是否持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他
		false,
		//是否为阻塞
		false,
		//额外参数
		nil,
	)
	if err != nil {
		fmt.Print(err)
	}
	// 2 发送消息到队列
	_ = r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		//如果为true，当队列没有绑定消费者，消息会返回给生产者，而不会保存在队列里面
		false,
		amqp.Publishing{
			Headers:         nil,
			ContentType:     "text/plain",
			ContentEncoding: "",
			DeliveryMode:    0,
			Priority:        0,
			CorrelationId:   "",
			ReplyTo:         "",
			Expiration:      "",
			MessageId:       "",
			Timestamp:       time.Time{},
			Type:            "",
			UserId:          "",
			AppId:           "",
			Body:            []byte(message),
		})
}

// Step3 简单模式下消费代码
func (r *RabbitMQ) ConsumeSimple() {
	// 1 申请队列，如果队列不存在则会创建队列
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//控制消息是否持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他
		false,
		//是否为阻塞
		false,
		//额外参数
		nil,
	)
	if err != nil {
		fmt.Print(err)
	}
	message, err := r.channel.Consume(
		r.QueueName,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Print(err)
	}

	//启用协程处理消息
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		for d := range message {
			log.Printf("this is msg %s", d.Body)
		}
	}()
	log.Print("waiting for message")
	w.Wait()
}
