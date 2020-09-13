package main

import (
	"github.com/mailgun/mailgun-go/v3"
	"strings"
	"sync"
	log "github.com/sirupsen/logrus"
	"time"
)

func SaveBouncesData(db MysqlDB, data []mailgun.Bounce, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, row := range data {
		dt := time.Unix(row.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05")
		_, errSql := db.ExecuteNonQuery("INSERT INTO MailgunBounces ( domain, code, email, created_at, error, updatedon )" +
			" VALUES (?, ?, ?, ?, ?, now()) ", mailgunDomain, row.Code, row.Address, dt, row.Error)
		if errSql != nil {
			log.Error(errSql)
		}
	}
}


func SaveComplaintsData(db MysqlDB, data []mailgun.Complaint, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, row := range data {
		dt := time.Unix(row.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05")
		_, errSql := db.ExecuteNonQuery("INSERT INTO MailgunComplaints ( domain, complaints_count, email, created_at, updatedon )" +
			" VALUES (?, ?, ?, ?, now()) ", mailgunDomain, row.Count, row.Address, dt)
		if errSql != nil {
			log.Error(errSql)
		}
	}
}

func SaveUnsubscribesData(db MysqlDB, data []mailgun.Unsubscribe, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, row := range data {
		dt := time.Unix(row.CreatedAt.Unix(), 0).Format("2006-01-02 15:04:05")
		_, errSql := db.ExecuteNonQuery("INSERT INTO MailgunUnsubscribes ( domain, tags, email, created_at, updatedon ) VALUES (?, ?, ?, ?, now()) ", mailgunDomain, strings.Join(row.Tags, ","), row.Address, dt)
		if errSql != nil {
			log.Error(errSql)
		}
	}
}



