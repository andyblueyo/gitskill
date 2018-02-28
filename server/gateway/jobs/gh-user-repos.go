package jobs

import (
	"github.com/andyblueyo/gitskill/server/gateway/services"
	"github.com/andyblueyo/gitskill/server/gateway/models/gh-repo"
	"fmt"
)

func ListenForRepos(repoChannel *chan gh_repo.Repo, reposWithLanguages *chan gh_repo.Repo, token string) {
	for {
		select {
		case repo := <-*repoChannel:
			rwl, err := services.GetRepoLanguage(&repo, token)
			if err != nil {
				fmt.Printf("error getting github repo langugaes: %v\n", err)
				return
			}
			*reposWithLanguages <- *rwl
		default: // nothing
		}
	}
}


	