package cmd

import (
	"context"
	"encoding/json"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"log"
)

func PushNotifications(projectId string, msgs []messaging.Message, dryrun bool) {
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

	for i := 0; i < len(msgs); i++ {
		rawBytes, err := json.Marshal(msgs[i])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d/%d\t%s\n", i, len(msgs), string(rawBytes))

		if !dryrun {
			r, err := client.Send(ctx, &msgs[i])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(r)
		}
	}
}
