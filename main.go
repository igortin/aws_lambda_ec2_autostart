package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

// handler GO
func handler(ctx context.Context, event events.CloudWatchEvent) {
	message := &EventDetails{}
	log.Println("INFO: Trying to Unmarshal ...")

	err := json.Unmarshal(event.Detail, message)
	if err != nil {
		log.Panic(err)
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := ec2.New(sess)

	// log.Printf("Paradise env: %v\n", message.InstanceID)
	isTrue, err := ec2StartInstance(svc, message)

	if err != nil {
		log.Fatal(err)
	}

	if isTrue {
		log.Println("INFO: Succesfully completed")
	} else {
		log.Println("Error: Logs should be viewed")
	}
}

func main() {
	lambda.Start(handler)
}

func ec2StartInstance(svc *ec2.EC2, msg *EventDetails) (bool, error) {
	input := &ec2.StartInstancesInput{
		DryRun: aws.Bool(false),
		InstanceIds: []*string{
			aws.String(msg.InstanceID),
		},
	}

	output, err := svc.StartInstances(input)

	if err != nil {
		return false, err
	}
	for _, node := range output.StartingInstances {
		log.Printf("Node %s was started\n", *node.InstanceId)
	}
	return true, nil
}
