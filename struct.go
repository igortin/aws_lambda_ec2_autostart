package main

/*
{
  "version": "1.0",
  "timestamp": "2020-07-15T08:31:35.452Z",
  "requestContext": {
    "requestId": "6bcbc196-2a44-4e96-bf6b-70ca241d442d",
    "functionArn": "arn:aws:lambda:us-east-1:625579157763:function:cloudWatchEventDebugLambda:$LATEST",
    "condition": "Success",
    "approximateInvokeCount": 1
  },
  "requestPayload": {
    "version": "0",
    "id": "c64bea0e-33f9-7fd7-fc1a-2d2ca2d45eec",
    "detail-type": "EC2 Instance State-change Notification",
    "source": "aws.ec2",
    "account": "625579157763",
    "time": "2020-07-15T08:31:34Z",
    "region": "us-east-1",
    "resources": [
      "arn:aws:ec2:us-east-1:625579157763:instance/i-0abdee38da9e51098"
    ],
    "detail": {
      "instance-id": "i-0abdee38da9e51098",
      "state": "stopped"
    }
  },
  "responseContext": {
    "statusCode": 200,
    "executedVersion": "$LATEST"
  },
  "responsePayload": null
}
*/

// EventDetails structure
type EventDetails struct {
	InstanceID string `json:"instance-id"`
	State      string `json:"state"`
}
