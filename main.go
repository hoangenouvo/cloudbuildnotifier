package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("Failed to load env file")
	}
}

func main() {
	ctx := context.Background()
	proj := os.Getenv("PROJECT_ID")
	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}
	// Pull messages via the subscription.
	if err := pullMsgs(client, "cloudBuildSub"); err != nil {
		log.Fatal(err)
	}
}

func pullMsgs(client *pubsub.Client, name string) error {
	var mu sync.Mutex
	sub := client.Subscription(name)
	err := sub.Receive(context.Background(), func(ctx context.Context, msg *pubsub.Message) {
		msg.Ack()
		var cloudBuildInfo CloudBuildInfo
		err := json.Unmarshal(msg.Data, &cloudBuildInfo)
		if err != nil {
			log.Printf("Got err: %s\n", err)
		}
		if cloudBuildInfo.Status == "FAILURE" && cloudBuildInfo.Substitutions.REPONAME == "ProjectStrand" &&
			(cloudBuildInfo.Substitutions.BRANCHNAME == "dev" || cloudBuildInfo.Substitutions.BRANCHNAME == "master") {
			githubData, err := GetGithubInfo(cloudBuildInfo.Substitutions.COMMITSHA)
			if err != nil {
				log.Println(err)
			}
			buildType := func() string {
				if cloudBuildInfo.Substitutions.BRANCHNAME == "dev" {
					return "nightly"
				} else {
					return "production"
				}
			}()
			message := fmt.Sprintf("Cloud build for *%s* has been finished with status %s. Detail infomations: ```Repo: %s\nBranch: %s\nCommit message: %s\nCommit Url: %s\nAuthor: %s(%s)\nCommitter:%s(%s)\n```",
				buildType, cloudBuildInfo.Status, cloudBuildInfo.Substitutions.REPONAME, cloudBuildInfo.Substitutions.BRANCHNAME, githubData.Message, githubData.HTML_URL,
				githubData.Author.Name, githubData.Author.Email, githubData.Committer.Name, githubData.Committer.Email)
			err = PushMessageToChatHangout(message)
			if err != nil {
				log.Println(err)
			}
		}
		mu.Lock()
		defer mu.Unlock()
	})
	if err != nil {
		return err
	}
	return nil
}

func PushMessageToChatHangout(message string) error {
	url := os.Getenv("HANGOUT_URL")
	method := "POST"
	messageBody := make(map[string]string)
	messageBody["text"] = message
	payload, err := json.Marshal(messageBody)
	if err != nil {
		return err
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return errors.New("Push message to hangout failed ")
	}
	fmt.Println("A message has been sent to Cloud-build CI Room: ", message)
	return nil
}

func GetGithubInfo(commitRSA string) (githubData GithubInfo, err error) {
	url := fmt.Sprintf("https://api.github.com/repos/trunghlt/ProjectStrand/git/commits/%s", commitRSA)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return GithubInfo{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", os.Getenv("GITHUB_TOKEN")))
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return GithubInfo{}, err
	}
	err = json.Unmarshal(body, &githubData)
	if err != nil {
		return GithubInfo{}, err
	}
	return githubData, nil
}
