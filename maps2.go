package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"log"
)

const (
	UPDATE_STATUS_PENDING = "PENDING"
	UPDATE_STATUS_SUCCESS = "SUCCESS"
	UPDATE_STATUS_ERROR   = "ERROR"
)

// ProductSettingUpdate An individual product setting update, encapsulating a globally unique ID and update status
type ProductSettingUpdate struct {
	Uuid   string
	Status string
}

type Topic string
type Partition int64
type Offset int64

// MessageId A globally unique ID for a Kafka message: Topic name + Partition number + Offset number
type MessageId struct {
	Topic     Topic
	Partition Partition
	Offset    Offset
}

// Message A DSOM message with one or more product settings updates and/or deletes
type Message struct {
	MessageId  MessageId
	CommitFunc func(error)
	Updates    map[string]*ProductSettingUpdate
}

// MessageLog Map of messages: message id to updates
type MessageLog map[MessageId]*Message

// BulkUpdates Map of BulkProcessor request id to its associated updates.
type BulkUpdates map[int64][]string

// UpdateToMessageMap Update to message map
type UpdateToMessageMap map[string]MessageId

func main() {

	// Initialize the misc. data structs
	messagesLog := make(MessageLog)
	bulkUpdates := make(BulkUpdates)
	updateToMessageMap := make(UpdateToMessageMap)

	// New message is read

	topic := "topic_a"
	partition := 0
	offset := 0
	commitFunc := func(err error) {
		fmt.Println("commitFunc", err)
	}
	msg1ID := MessageId{
		Topic:     Topic(topic),
		Partition: Partition(partition),
		Offset:    Offset(offset),
	}
	msg1 := Message{
		MessageId:  msg1ID,
		CommitFunc: commitFunc,
		Updates:    make(map[string]*ProductSettingUpdate),
	}

	update1Uuid := uuid.NewString()
	update1 := ProductSettingUpdate{
		Uuid:   update1Uuid,
		Status: UPDATE_STATUS_PENDING,
	}

	msg1.Updates[update1Uuid] = &update1

	messagesLog[msg1ID] = &msg1
	updateToMessageMap[update1Uuid] = msg1ID

	fmt.Println("Messages Log:", messagesLog)
	fmt.Println("update-to-message log map:", updateToMessageMap)

	// Before()
	bulkUpdate1ID := int64(1000)
	bulkUpdates[bulkUpdate1ID] = make([]string, 0)
	bulkUpdates[bulkUpdate1ID] = append(bulkUpdates[bulkUpdate1ID], update1Uuid)

	fmt.Println("Bulk updated:", bulkUpdates)

	// After()
	// For each update in bulkUpdate 'bulkUpdate1ID', find the corresponding update in the MessagesLog
	fmt.Println("Processing bulk update ", bulkUpdate1ID)
	for _, updateUuid := range bulkUpdates[bulkUpdate1ID] {
		fmt.Println("Processing update ", updateUuid)
		messageId, ok := updateToMessageMap[updateUuid]
		if !ok {
			fmt.Println(fmt.Errorf("No message is associated with update", updateUuid))
		} else {
			fmt.Println("Found message id", messageId)

			message, ok := messagesLog[messageId]
			if !ok {
				log.Fatalln("No message associated with message id", messageId)
			} else {
				fmt.Println("Found message in messages log", message)

				update, ok := message.Updates[updateUuid]
				if !ok {
					log.Fatalln("No update associated with update id", updateUuid)
				} else {
					fmt.Println("Found update id", updateUuid)

					update.Status = UPDATE_STATUS_ERROR
				}
			}
		}
	}

	fmt.Println("Messages Log:", messagesLog)

	for messageId, message := range messagesLog {
		fmt.Println("\tmessage id", messageId)
		commitReady := true
		success := true
		for updateId, update := range message.Updates {
			fmt.Println("\t\tupdate", updateId, update.Status)
			if update.Status == UPDATE_STATUS_PENDING {
				commitReady = false
			} else if update.Status != UPDATE_STATUS_SUCCESS {
				success = false
			}
		}

		fmt.Println("\tCommit ready", commitReady)
		if commitReady {
			var result error
			if !success {
				result = errors.New("An error occurred")
			}
			message.CommitFunc(result)
		}
	}
}
