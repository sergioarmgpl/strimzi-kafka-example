package main

import (
    "fmt"
    "os"
    "net/http"

    "github.com/confluentinc/confluent-kafka-go/kafka"
    "github.com/gin-gonic/gin"
)


type Data struct {
    Key     string  `json:"key"`
    Value  string  `json:"value"`
}


func main() {
    router := gin.Default()
    router.POST("/data", postKeyValue)
    router.Run("0.0.0.0:3000")
}


func postKeyValue(c *gin.Context) {
    var data Data

    if err := c.BindJSON(&data); err != nil {
        return
    }

    topic := "mytopic"
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "my-cluster-kafka-0.my-cluster-kafka-brokers.kafka.svc:9092"})
    if err != nil {
        fmt.Printf("Failed to create producer: %s", err)
        os.Exit(1)
    }

    p.Produce(&kafka.Message{
        TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
        Key: []byte("data"),
        Value: []byte("{\""+data.Key+"\":\""+data.Value+"\"}"),
    }, nil)

    // Wait for all messages to be delivered
    p.Flush(1 * 1000)
    p.Close()
    c.IndentedJSON(http.StatusCreated, data)
}




