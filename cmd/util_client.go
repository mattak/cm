package cmd

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"log"
)

func PushNotifications(projectId string, msgs []*messaging.Message, dryrun bool) {
	ctx := context.Background()
	config := firebase.Config{ProjectID: projectId}
	app, err := firebase.NewApp(ctx, &config)
	if err != nil {
		log.Fatal(err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if !dryrun {
		res, err := client.SendAll(ctx, msgs)
		if res != nil {
			fmt.Printf("RESULT:RUN\tOK:%d\tNG:%d\n", res.SuccessCount, res.FailureCount)
			for i := 0; i < len(res.Responses); i++ {
				if !res.Responses[i].Success {
					fmt.Printf("%d\t%s\t%v\n", i, res.Responses[i].MessageID, res.Responses[i].Error)
				}
			}
		}
		if err != nil {
			log.Fatal(err)
		}
	} else {
		res, err := client.SendAllDryRun(ctx, msgs)
		if res != nil {
			fmt.Printf("RESULT:DRYRUN\tOK:%d\tNG:%d\n", res.SuccessCount, res.FailureCount)
			for i := 0; i < len(res.Responses); i++ {
				if !res.Responses[i].Success {
					fmt.Printf("%d\t%s\t%v\n", i, res.Responses[i].MessageID, res.Responses[i].Error)
				}
			}
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
