package MongoDB

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestMongoDbStart(t *testing.T) {
	// 控制初始化超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	monitor := &event.CommandMonitor{
		Started: func(ctx context.Context, startedEvent *event.CommandStartedEvent) {
			fmt.Println(startedEvent.Command)
		},
		Succeeded: nil,
		Failed:    nil,
	}

	opts := options.Client().
		ApplyURI("mongodb://root:example@localhost:27017").
		SetMonitor(monitor)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	defer func() {}()

	mdb := client.Database("webook")
	col := mdb.Collection("articles")

	res, err := col.InsertOne(ctx, Article{
		Id:      123,
		Title:   "我的标题",
		Content: "我的内容",
	})
	if err != nil {
		panic(err)
	}

	//他自己里面的id
	fmt.Println("id", res.InsertedID)
}

type Article struct {
	Id       int64  `bson:"id,omitempty"`
	Title    string `bson:"title,omitempty"`
	Content  string `bson:"content,omitempty"`
	AuthorId int64  `bson:"author_id,omitempty"`
	Status   uint8  `bson:"status,omitempty"`
	Ctime    int64  `bson:"ctime,omitempty"`
	Utime    int64  `bson:"utime,omitempty"`
}
