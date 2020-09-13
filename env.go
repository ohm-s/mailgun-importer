package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var dbHost string
var dbPort string
var dbUser string
var dbPass string
var dbSchema string
var dbPoolSize int
var dbIdleSize int
var mailgunDomain string
var mailgunKey string
var fetchMailgunBounces bool
var fetchMailgunComplaints bool
var fetchMailgunUnsubscribes bool


func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Warn("Could not load .env ")
	}
	dbHost = os.Getenv("db.host")
	if dbHost == "" {
		log.Panic("environment variable db.host not defined")
		os.Exit(1)
	}
	dbPort =  os.Getenv("db.port")
	dbUser = os.Getenv("db.user")
	dbPass = os.Getenv("db.pass")
	dbSchema = os.Getenv("db.schema")
	mailgunDomain = os.Getenv("mailgun.domain")
	mailgunKey = os.Getenv("mailgun.key")
	poolSize, _ := strconv.Atoi(os.Getenv("db.poolsize"))
	idleSize, _ := strconv.Atoi(os.Getenv("db.idlesize"))
	if poolSize == 0 {
		dbPoolSize = 1
	} else {
		dbPoolSize = poolSize
	}
	if idleSize == 0 {
		dbIdleSize = 1
	} else {
		dbIdleSize = idleSize
	}

	fetchComplaints, _ := strconv.Atoi(os.Getenv("fetcher.complaints"))
	fetchBounces, _ := strconv.Atoi(os.Getenv("fetcher.bounces"))
	fetchUnsubscribes, _ := strconv.Atoi(os.Getenv("fetcher.unsubscribes"))
	if fetchComplaints >  0 {
		fetchMailgunComplaints = true
	}
	if fetchBounces > 0 {
		fetchMailgunBounces = true
	}
	if fetchUnsubscribes > 0 {
		fetchMailgunUnsubscribes = true
	}

}

