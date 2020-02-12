package main

import "time"

type CloudBuildInfo struct {
	ID               string           `json:"id"`
	ProjectID        string           `json:"projectId"`
	Status           string           `json:"status"`
	Source           Source           `json:"source"`
	Steps            []Steps          `json:"steps"`
	Results          Results          `json:"results"`
	CreateTime       time.Time        `json:"createTime"`
	StartTime        time.Time        `json:"startTime"`
	FinishTime       time.Time        `json:"finishTime"`
	Timeout          string           `json:"timeout"`
	LogsBucket       string           `json:"logsBucket"`
	SourceProvenance SourceProvenance `json:"sourceProvenance"`
	BuildTriggerID   string           `json:"buildTriggerId"`
	Options          Options          `json:"options"`
	LogURL           string           `json:"logUrl"`
	Substitutions    Substitutions    `json:"substitutions"`
	Tags             []string         `json:"tags"`
	Timing           interface{}      `json:"timing"`
}
type StorageSource struct {
	Bucket string `json:"bucket"`
	Object string `json:"object"`
}
type Source struct {
	StorageSource StorageSource `json:"storageSource"`
}
type Timing struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
type PullTiming struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
type Steps struct {
	Name       string     `json:"name"`
	Args       []string   `json:"args"`
	ID         string     `json:"id"`
	WaitFor    []string   `json:"waitFor,omitempty"`
	Entrypoint string     `json:"entrypoint,omitempty"`
	Timing     Timing     `json:"timing,omitempty"`
	PullTiming PullTiming `json:"pullTiming,omitempty"`
	Status     string     `json:"status"`
	Dir        string     `json:"dir,omitempty"`
	Env        []string   `json:"env,omitempty"`
}
type Results struct {
	BuildStepImages []string `json:"buildStepImages"`
}
type ResolvedStorageSource struct {
	Bucket     string `json:"bucket"`
	Object     string `json:"object"`
	Generation string `json:"generation"`
}

type SourceProvenance struct {
	ResolvedStorageSource ResolvedStorageSource `json:"resolvedStorageSource"`
	FileHashes            interface{}           `json:"fileHashes"`
}
type Options struct {
	SubstitutionOption string `json:"substitutionOption"`
	Logging            string `json:"logging"`
}
type Substitutions struct {
	BRANCHNAME          string `json:"BRANCH_NAME"`
	COMMITSHA           string `json:"COMMIT_SHA"`
	REPONAME            string `json:"REPO_NAME"`
	REVISIONID          string `json:"REVISION_ID"`
	SHORTSHA            string `json:"SHORT_SHA"`
	BASEBRANCH          string `json:"_BASE_BRANCH"`
	DEPLOYERIMAGE       string `json:"_DEPLOYER_IMAGE"`
	FULFILLMENTIMAGE    string `json:"_FULFILLMENT_IMAGE"`
	GOOGLECLOUDSDK      string `json:"_GOOGLE_CLOUD_SDK"`
	GOIMAGE             string `json:"_GO_IMAGE"`
	HEADBRANCH          string `json:"_HEAD_BRANCH"`
	HEADREPOURL         string `json:"_HEAD_REPO_URL"`
	NAMESPACE           string `json:"_NAMESPACE"`
	NIFIIMAGE           string `json:"_NIFI_IMAGE"`
	PRNUMBER            string `json:"_PR_NUMBER"`
	SPARKJOBSERVERIMAGE string `json:"_SPARK_JOBSERVER_IMAGE"`
	SUPERSETIMAGE       string `json:"_SUPERSET_IMAGE"`
}

type GithubInfo struct {
	SHA          string       `json:"sha"`
	NodeID       string       `json:"node_id"`
	URL          string       `json:"url"`
	HTML_URL     string       `json:"html_url"`
	Author       PersonInfo   `json:"author"`
	Committer    PersonInfo   `json:"committer"`
	Tree         Tree         `json:"tree"`
	Message      string       `json:"message"`
	Parents      []Parent     `json:"parents"`
	Verification Verification `json:"verification"`
}

type PersonInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Date  time.Time
}

type Tree struct {
	SHA string `json:"sha"`
	URL string `json:"url"`
}

type Parent struct {
	SHA      string `json:"sha"`
	URL      string `json:"url"`
	HTML_URL string `json:"html_url"`
}

type Verification struct {
	Verified  bool        `json:"verified"`
	Reason    string      `json:"reason"`
	Signature interface{} `json:"signature"`
	Payload   interface{} `json:"payload"`
}
