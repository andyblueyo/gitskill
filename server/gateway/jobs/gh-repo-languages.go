package jobs

import (
	"github.com/andyblueyo/gitskill/server/gateway/models/gh-repo"
	"fmt"
)

func ListenForReposWithLanguages(repos *chan gh_repo.Repo, store *gh_repo.MongoStore) {
	for {
		select {
		case repo := <-*repos:
			go WriteRepoToDB(repo, store)
		default: // nothing
		}
	}
}

func WriteRepoToDB(repo gh_repo.Repo, store *gh_repo.MongoStore) {
	_, err := store.Insert(&repo)
	if err != nil {
		fmt.Printf("error inserting repo: %v\n", err)
	}
}