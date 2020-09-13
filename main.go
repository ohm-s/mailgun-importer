package main

import (
    "fmt"
    log "github.com/sirupsen/logrus"
    "os"
    "sync"
)

func main() {

    log.Print("Starting up")
    db, err := GetMysqlConnection(dbHost, dbPort, dbSchema, dbUser, dbPass, dbPoolSize, dbIdleSize)
    if err != nil {
        panic(err)
        os.Exit(1)
    }

    var mgChannel = make(chan MailgunChannelMessage, dbPoolSize)

    jobs := 0
    if fetchMailgunBounces {
        go ListBounces(mailgunDomain, mailgunKey, mgChannel)
        jobs++
    }
    if fetchMailgunComplaints {
        go ListComplaints(mailgunDomain, mailgunKey, mgChannel)
        jobs++
    }
    if fetchMailgunUnsubscribes {
        go ListUnsubscribes(mailgunDomain, mailgunKey, mgChannel)
        jobs++
    }

    var wg sync.WaitGroup
    hasErroredOut := false

    for jobs > 0 {
        batch := <-mgChannel

        fmt.Printf("+%o \n", batch)
        os.Exit(1)
        if batch.errored {
            hasErroredOut = true
            jobs--
        }  else if batch.done {
            jobs --
        } else if len(batch.bounces) > 0 {
            wg.Add(1)
            go SaveBouncesData(db, batch.bounces, &wg)
        } else if len(batch.complaints) > 0 {
            wg.Add(1)
            go SaveComplaintsData(db, batch.complaints, &wg)
        } else if len(batch.unsubscribes) > 0 {
            wg.Add(1)
            go SaveUnsubscribesData(db, batch.unsubscribes, &wg)
        }

    }
    wg.Wait()

    if hasErroredOut {
        log.Panic("Program has errored out during execution")
    }
}



