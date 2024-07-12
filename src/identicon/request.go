package identicon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GithubUser struct {
	Id       uint   `json:"id"`
	UserName string `json:"login"`
}

func requestAccoundId(userName string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s", userName)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	if err != nil {
		return "", err
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var user GithubUser
	json.Unmarshal(resBody, &user)

	return fmt.Sprint(user.Id), nil
}
