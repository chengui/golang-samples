package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type NewIssue struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type SearchQuery struct {
	Filter string `json:"filter"` // assigned,created,mentioned, subscribed, repos, all
	State  string `json:"state"`  // open,closed,all
	Sort   string `json:"sort"`   // created, updated, comments
	Since  string `json:"since"`  // YYYY-MM-DDTHH:MM:SSZ
}

const ApiURL = "https://api.github.com"
const ApiToken = ""

type Github struct {
	IssueURL string
}

func New(repo string) *Github {
	return &Github{
		IssueURL: fmt.Sprintf("%s/repos/%s/issues", ApiURL, repo),
	}
}

func (g *Github) SetHeader(req *http.Request) {
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+ApiToken)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
}

func (g *Github) SearchIssues(query map[string]string) (*[]NewIssue, error) {
	var param string
	for k, v := range query {
		if v != "" {
			if param == "" {
				param += "?" + k + "=" + v
			} else {
				param += "&" + k + "=" + v
			}
		}
	}
	fmt.Println(param)
	req, err := http.NewRequest("GET", g.IssueURL+param, nil)
	if err != nil {
		return nil, err
	}
	g.SetHeader(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get failed: %s", resp.Status)
	}
	issues := []NewIssue{}
	if err := json.NewDecoder(resp.Body).Decode(&issues); err != nil {
		return nil, err
	}
	return &issues, nil
}

func (g *Github) CreateIssue(issue *NewIssue) (bool, error) {
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(issue); err != nil {
		return false, err
	}
	req, err := http.NewRequest("POST", g.IssueURL, body)
	if err != nil {
		return false, err
	}
	g.SetHeader(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("get failed: %s", resp.Status)
	}
	return true, nil
}

func (g *Github) GetIssue(id int) (*NewIssue, error) {
	issueURL := fmt.Sprintf("%s/%d", g.IssueURL, id)
	req, err := http.NewRequest("GET", issueURL, nil)
	if err != nil {
		return nil, err
	}
	g.SetHeader(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("get failed: %s", resp.Status)
	}
	var issue NewIssue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func (g *Github) UpdateIssue(id int, issue *NewIssue) (bool, error) {
	body, err := json.Marshal(issue)
	if err != nil {
		return false, err
	}
	issueURL := fmt.Sprintf("%s/%d", g.IssueURL, id)
	req, err := http.NewRequest("PATCH", issueURL, bytes.NewReader(body))
	if err != nil {
		return false, err
	}
	g.SetHeader(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("get failed: %s", resp.Status)
	}
	return true, nil
}

func (g *Github) CloseIssue(id int) (bool, error) {
	body, err := json.Marshal(map[string]string{"state": "closed"})
	if err != nil {
		return false, err
	}
	issueURL := fmt.Sprintf("%s/%d", g.IssueURL, id)
	req, err := http.NewRequest("PATCH", issueURL, bytes.NewReader(body))
	if err != nil {
		return false, err
	}
	g.SetHeader(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("get failed: %s", resp.Status)
	}
	return true, nil
}
