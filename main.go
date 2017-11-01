package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type Story struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	URL        string `json:"url"`
	ExternalID string `json:"external_id, omitempty"`
	StoryType  string `json:"story_type"`
}

type StoryResponse []struct {
	ProjectID int     `json:"project_id"`
	StoryList []Story `json:"stories"`
}

func main() {

	// Part project flags
	projectPtr := flag.String("project", "default", "project_id")
	keyPtr := flag.String("key", "default", "tracker api key")
	jiraPtr := flag.String("jira", "https://issues.apache.org/jira/browse/", "linked jira url")
	flag.Parse()

	//Print if default flags are used
	if *projectPtr == "default" || *keyPtr == "default" {
		PrintUsageAndExit()
	}

	fmt.Println("Fetching project with id", *projectPtr)

	// create the url
	done_url := fmt.Sprintf("https://www.pivotaltracker.com/services/v5/projects/%s/iterations?scope=done&limit=1&offset=-1", *projectPtr)
	current_url := fmt.Sprintf("https://www.pivotaltracker.com/services/v5/projects/%s/iterations?scope=current&offset=2", *projectPtr)

	// Make request to tracker
	client := &http.Client{}

	f, err := os.OpenFile("autobeat.md", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	now := time.Now().Local()
	then := now.AddDate(0, 0, -7)
	updateStr := fmt.Sprintf("# Update: %s - %s\n", then.Format("2006-01-02"), now.Format("2006-01-02"))
	f.WriteString(updateStr)

	f.WriteString("## Done\n")
	stories := MakeStoryRequest(done_url, keyPtr, client)
	buffer := PrintStories(stories, jiraPtr)
	f.Write(buffer)

	f.WriteString("## Next\n")
	stories = MakeStoryRequest(current_url, keyPtr, client)
	buffer = PrintStories(stories, jiraPtr)
	f.Write(buffer)

	fmt.Println("\nUpdate written to ./autobeat.md")
	os.Exit(0)

}

func MakeStoryRequest(done_url string, keyPtr *string, client *http.Client) StoryResponse {
	req, _ := http.NewRequest("GET", done_url, nil)
	req.Header.Set("X-TrackerToken", *keyPtr)
	res, _ := client.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	s, _ := ReadStoryResponse(body)
	return s
}

func ReadStoryResponse(body []byte) (StoryResponse, error) {
	s := StoryResponse{}
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("Error parsing tracker reponse", err)
	}
	return s, err
}

func PrintStories(s StoryResponse, jiraPtr *string) (buffer []byte) {

	for i := 0; i < len(s[0].StoryList); i++ {

		story := s[0].StoryList[i]

		if len(story.ExternalID) == 0 {
			writeStr := fmt.Sprintf("* %s [Tracker](%s)\n", story.Name, story.URL)
			buffer = append(buffer, writeStr...)
		} else {
			writeStr := fmt.Sprintf("* %s [JIRA](%s/%s) | [Tracker](%s)\n", story.Name, *jiraPtr, story.ExternalID, story.URL)
			buffer = append(buffer, writeStr...)
		}
	}
	return buffer
}
