package main

import (
	"context"
	"fmt"
	"time"
	"os"
	"log"

	"github.com/joho/godotenv"
	"github.com/google/go-github/v64/github"
)

const ownerName = "git"
const repositoryName = "git"

func zeroFillDate(t time.Time) time.Time {
	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		0, 0, 0, 0, t.Location(),
	)
}

func hasCommit(date time.Time) (bool, error) {
	// https://docs.github.com/en/rest/commits/commits?apiVersion=2022-11-28
	client := github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_TOKEN"))
	commits, _, err := client.Repositories.ListCommits(context.Background(), ownerName, repositoryName, &github.CommitsListOptions{
		Until:       date,
		ListOptions: github.ListOptions{PerPage: 100},
	})
	if err != nil {
		return false, err
	}

	return len(commits) > 0, nil
}

func getFirstCommits(date time.Time) ([]*github.RepositoryCommit, error) {
	client := github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_TOKEN"))
	commits, _, err := client.Repositories.ListCommits(context.Background(), ownerName, repositoryName, &github.CommitsListOptions{
		Until:       date,
		ListOptions: github.ListOptions{PerPage: 100},
	})
	if err != nil {
		return nil, err
	}

	if len(commits) == 0 {
		return nil, nil
	}

	return commits, nil
}

func findBoundaryDate(start, end time.Time) (time.Time, error) {
	count := 0
	for start.Before(end) {
		mid := start.Add(end.Sub(start) / 2)
		mid = zeroFillDate(mid)

		hasCommit, err := hasCommit(mid)
		if err != nil {
			return time.Time{}, err
		}

		if hasCommit {
			end = mid
		} else {
			start = mid.Add(24 * time.Hour)
		}

		start = zeroFillDate(start)
		end = zeroFillDate(end)
		count++
		fmt.Println("in progress:", count, hasCommit, start, end)
	}

	fmt.Println("search count:", count)
	return start, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	startTime := time.Now()

	startDate := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 8, 25, 0, 0, 0, 0, time.UTC)

	boundaryDate, err := findBoundaryDate(startDate, endDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	commits, err := getFirstCommits(boundaryDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	fmt.Printf("execute time: %v\n", duration)

	fmt.Println("boundary date:", boundaryDate.Format("2006-01-02"))
	for i := 0; i < len(commits) && i < 5; i++ {
		fmt.Println(commits[len(commits)-(i+1)].GetHTMLURL())
	}
}
