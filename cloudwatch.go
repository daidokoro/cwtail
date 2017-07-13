package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type logGroup struct {
	group   string
	streams []*cloudwatchlogs.LogStream
}

// event struct stores log event response to pass back to the channel
type event struct {
}

func newLogGroup(g string) *logGroup {
	return &logGroup{
		group: g,
	}
}

func (l *logGroup) getEvents(sess *session.Session, logs chan<- *cloudwatchlogs.OutputLogEvent, params func() *cloudwatchlogs.GetLogEventsInput) error {
	log.Debug(fmt.Sprintln("fetching events with params:", params()))
	svc := cloudwatchlogs.New(sess)
	args := params()

	t := time.NewTicker(time.Millisecond * 1500)
	for _ = range t.C {
		resp, err := svc.GetLogEvents(args)
		if err != nil {
			return err
		}

		// log.Debug(fmt.Sprintln("Received response:", resp))

		for _, evt := range resp.Events {
			logs <- evt
		}

		args.NextToken = resp.NextForwardToken
	}
	return nil
}

func (l *logGroup) getStreams(sess *session.Session, params func() *cloudwatchlogs.DescribeLogStreamsInput) error {
	log.Debug(fmt.Sprintln("fetching streams with params:", params()))
	svc := cloudwatchlogs.New(sess)
	resp, err := svc.DescribeLogStreams(params())
	if err != nil {
		return err
	}
	log.Debug(fmt.Sprintln("response:", resp))

	l.streams = resp.LogStreams
	return nil
}

func printEvents(done <-chan bool, logs <-chan *cloudwatchlogs.OutputLogEvent) {
	for {
		select {
		case l := <-logs:
			log.Info(fmt.Sprintf(
				"%s - %s",
				log.ColorString(time.Unix(*l.Timestamp/1000, 0).String(), "cyan"),
				*l.Message,
			))
		case <-done:
			return
			// default:
			//NOTHING
		}
	}
}
