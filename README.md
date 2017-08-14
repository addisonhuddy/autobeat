# Autobeat
A tool for making the creation of update emails a little easier

## Installation
To install from source, simply `go install github.com/addisonhuddy/autobeat`. This assumes you have the Go toolchain installed.

Or use the OS X or Linux binaries on the Github release page.

## Usage

`autobeat -project=[PROJECT ID] -key=[TRACKER API KEY] -jira=[url to JIRA issues]`

Default JIRA url is https://issues.apache.org/jira/browse/.

The above command will produce `autobeat.md` in the same directory.  Edit and use a tool like [Markdown Here](http://markdown-here.com/) to make sending update emails a breeze.