package main

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	var err error

	url, err := createPullRequestURL()
	if err != nil {
		log.Fatal(err)
	}

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll, FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	case "default":
		log.Fatal("unsupported OS")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func createPullRequestURL() (string, error) {
	// get repository name from remote repository
	repositoryName, err := getRepositoryName()
	if err != nil {
		return "", err
	}

	// get branch name
	branchName, err := getBranchName()
	if err != nil {
		return "", err
	}

	// remove line break and create https git url
	url := strings.NewReplacer(
		"\r\n", "",
		"\r", "",
		"\n", "",
	).Replace(
		"https://github.com/" + repositoryName + "/pull/new/" + branchName,
	)

	return url, nil
}

func getRepositoryName() (string, error) {
	gitCommand := exec.Command("git", "config", "--get", "remote.origin.url")
	gitOutput, err := gitCommand.Output()
	if err != nil {
		return "", err
	}

	repoName := strings.NewReplacer(
		"git@github.com:", "",
		".git", "",
	).Replace(
		string(gitOutput),
	)

	return repoName, nil
}

func getBranchName() (string, error) {
	gitCommand := exec.Command("git", "symbolic-ref", "--short", "HEAD")

	gitOutput, err := gitCommand.Output()
	if err != nil {
		return "", err
	}

	return string(gitOutput), nil
}
