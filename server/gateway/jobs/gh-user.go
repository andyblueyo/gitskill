package jobs

import (
	"github.com/andyblueyo/gitskill/server/gateway/services"
	"github.com/andyblueyo/gitskill/server/gateway/models/gh-repo"
	"github.com/andyblueyo/gitskill/server/gateway/models/gh-user"
	"fmt"
)

func ListenForAccounts(accounts *chan string, repoChannel *chan gh_repo.Repo, token string, store *gh_user.MongoStore) {
	for {
		select {
		case accountName := <-*accounts:
			if len(accountName) == 0 {
				fmt.Printf("account name no length")
				return
			}
			gu, err := services.GetGithubUser(accountName, token)
			if err != nil {
				fmt.Printf("error getting github user: %v\n", err)
				return
			}
			ngu, err := store.Insert(gu)
			if err != nil {
				fmt.Printf("some error happened saving github user: %v", err)
				return
			}
			totalRepos := gu.PublicRepos + gu.TotalPrivateRepos
			uts := string(gu.UserType)
			repos, err := services.GetGithubRepos(accountName, uts, token, totalRepos)
			if err != nil {
				fmt.Printf("error getting github repos: %v\n", err)
				return
			}
			for _, i := range repos {
				repo := i.ToRepo()
				repo.RepoOwnerID = ngu.ID
				*repoChannel <- *repo
			}
		default: // nothing
		}
	}
}
