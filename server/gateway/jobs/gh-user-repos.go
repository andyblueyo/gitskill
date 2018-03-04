package jobs

import (
	"github.com/andyblueyo/gitskill/server/gateway/services"
	"github.com/andyblueyo/gitskill/server/gateway/models/gh-repo"
	"fmt"
	"github.com/andyblueyo/gitskill/server/gateway/handlers"
)

func ListenForRepos(repoChannel *chan gh_repo.Repo, reposWithLanguages *chan gh_repo.Repo, ctx *handlers.HandlerContext) {
	for {
		select {
		case repo := <-*repoChannel:
			rwl, err := services.GetRepoLanguage(&repo, ctx.GetNextToken)
			if err != nil {
				fmt.Printf("error getting github repo langugaes: %v for repo: %s\n", err, repo)
			} else {
				*reposWithLanguages <- *rwl
			}
		default: // nothing
		}
	}
}
