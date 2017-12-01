package main

import (
	"encoding/json"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

func bootstrapLogger() *logrus.Logger {
	// init basic logger to stdout
	log := logrus.New()
	log.Out = os.Stdout
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.TextFormatter{FullTimestamp: true}

	// init logstash hook
	conn, err := net.Dial("tcp", "0.0.0.0:5044")
	if err != nil {
		log.Fatal(err)
	}

	hook := logrustash.New(conn, &logrus.JSONFormatter{FieldMap: logrus.FieldMap{
		logrus.FieldKeyTime:  "@timestamp",
		logrus.FieldKeyLevel: "@level",
		logrus.FieldKeyMsg:   "@message",
	}})

	log.Hooks.Add(hook)

	return log
}

func getChuckJoke() (string, error) {
	r, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	chuckJoke := struct {
		Value   string `json:"value"`
	}{}

	if err = json.NewDecoder(r.Body).Decode(&chuckJoke); err != nil {
		return "", err
	}

	return chuckJoke.Value, nil
}

func main() {
	log := bootstrapLogger()

	t := time.NewTimer(5 * time.Minute)
	for do := true; do; {
		joke, err := getChuckJoke()
		if err != nil {
			log.WithField("error", err).Errorln("failed to get joke")
		} else {
			log.Infoln(joke)
		}

		select {
		case <-t.C:
			do = false
		case <-time.After(500 * time.Millisecond):
		}
	}

	log.Info("end of example")
}
