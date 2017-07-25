package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/spf13/cobra"
)

// root command (calls all other commands)
var cwtailCmd = &cobra.Command{
	Use:     "cwtail",
	Short:   "Tail AWS CloudWatch Logs.",
	Example: "cwtail /aws/log/group",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}

		// create sessoin
		sess, err := manager.GetSess(profile)
		handleError(err)

		// make channels
		logs := make(chan *cloudwatchlogs.OutputLogEvent)
		done := make(chan bool)

		// print event
		go printEvents(done, logs)

		for _, group := range args {
			wg.Add(1)
			go func(group string) {
				defer wg.Done()
				lg := newLogGroup(group)
				if err := lg.getStreams(sess, func() *cloudwatchlogs.DescribeLogStreamsInput {
					return &cloudwatchlogs.DescribeLogStreamsInput{
						OrderBy:      aws.String("LastEventTime"),
						LogGroupName: &group,
						Limit:        aws.Int64(1),
						Descending:   aws.Bool(true),
					}
				}); err != nil {
					handleError(err)
				}

				// for _, streams := range lg.streams {
				if err := lg.getEvents(sess, logs, func() *cloudwatchlogs.GetLogEventsInput {
					return &cloudwatchlogs.GetLogEventsInput{
						LogGroupName: &group,
						// tail the most recent stream
						LogStreamName: lg.stream.LogStreamName,
					}
				}); err != nil {
					handleError(err)
				}
				// }
				return
			}(group)
		}

		wg.Wait()
		done <- true
		<-logs

	},
}

// define flags
func init() {
	cwtailCmd.PersistentFlags().StringVarP(&profile, "profile", "p", "default", "configured AWS profile")
	cwtailCmd.PersistentFlags().BoolVarP(&debug, "debug", "", false, "debug mode")
	cwtailCmd.PersistentFlags().BoolVarP(&colors, "no-colors", "", false, "disable color output")
}
