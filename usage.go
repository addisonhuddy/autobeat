package main

import (
	"fmt"
	"os"
)

func PrintUsageAndExit() {
	fmt.Println(`autobeat v1.0
Usage:
	autobeat -project=[PROJECT ID] -key=[TRACKER API KEY] -jira=[url of private JIRA]"

	Output will be to a markdown filwe in the same directory call autobeat.md. Copy this markdown into your
	email client of choice and use something like http://markdown-here.com/ to get nicely formatted html.

	Default Jira URL: https://issues.apache.org/jira/browse/

	Happy updating!
	`)

	os.Exit(1)
}
