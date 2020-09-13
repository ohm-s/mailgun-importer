package main

import (
	"github.com/mailgun/mailgun-go/v3"
	"log"
	"context"
)

type MailgunChannelMessage struct {
	bounces []mailgun.Bounce
	complaints []mailgun.Complaint
	unsubscribes []mailgun.Unsubscribe
	done bool
	errored bool
}

func ListBounces(domain, apiKey string, mgChannel chan MailgunChannelMessage)  {
	mg := mailgun.NewMailgun(domain, apiKey)
	it := mg.ListBounces(&mailgun.ListOptions{

	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var page []mailgun.Bounce
 	for it.Next(ctx, &page) {
 		if len(page) > 0 {
			mgChannel <- MailgunChannelMessage{bounces: page}
		}
	}

	if it.Err() != nil {
		log.Panic(it.Err())
		mgChannel <- MailgunChannelMessage{errored: true}
	} else {
		mgChannel <- MailgunChannelMessage{done: true}
	}
	log.Print("ListBounces End")
}

func ListComplaints(domain, apiKey string, mgChannel chan MailgunChannelMessage)  {
	mg := mailgun.NewMailgun(domain, apiKey)
	it := mg.ListComplaints(nil)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var page []mailgun.Complaint
	var counter = 0
	for it.Next(ctx, &page) {
		counter++
		mgChannel <- MailgunChannelMessage{complaints: page}
	}

	if it.Err() != nil {
		log.Panic(it.Err())
		mgChannel <- MailgunChannelMessage{errored: true}
	} else {
		mgChannel <- MailgunChannelMessage{done: true}
	}
	log.Print("ListComplaints End")
}


func ListUnsubscribes(domain, apiKey string, mgChannel chan MailgunChannelMessage)  {
	mg := mailgun.NewMailgun(domain, apiKey)
	it := mg.ListUnsubscribes(nil)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var page []mailgun.Unsubscribe
	var counter = 0
	for it.Next(ctx, &page) {
		counter++
		mgChannel <- MailgunChannelMessage{unsubscribes: page}
	}

	if it.Err() != nil {
		log.Panic(it.Err())
		mgChannel <- MailgunChannelMessage{errored: true}
	} else {
		mgChannel <- MailgunChannelMessage{done: true}
	}
	log.Print("ListUnsubscribes End")

}
