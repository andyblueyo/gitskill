package jobs

import (
	"github.com/andyblueyo/gitskill/server/gateway/services"
	"github.com/andyblueyo/gitskill/server/gateway/models/gh-repo"
	"github.com/andyblueyo/gitskill/server/gateway/models/gh-user"
	"github.com/andyblueyo/gitskill/server/gateway/handlers"
	"fmt"
)

func ListenForAccounts(
	accounts *chan string,
	repoChannel *chan gh_repo.Repo,
	orgsToScrapeUsers *chan string,
	ctx *handlers.HandlerContext,
	store *gh_user.MongoStore,
) {
	for {
		select {
		case accountName := <-*accounts:
			go ProcessUser(accountName, repoChannel, orgsToScrapeUsers, ctx, store)
		default: // nothing
		}
	}
}

func ProcessUser(
	accountName string,
	repoChannel *chan gh_repo.Repo,
	orgsToScrapeUsers *chan string,
	ctx *handlers.HandlerContext,
	store *gh_user.MongoStore,
) {
	fmt.Printf("getting account: %v\n", accountName)
	gu, err := services.GetGithubUser(accountName, ctx.GetNextToken)
	if err != nil {
		fmt.Printf("error getting github user: %v\n", err)
	} else {
		if gu.UserType == gh_user.GHTypeOrganization {
			*orgsToScrapeUsers <- accountName
		} else {
			orgs, err := services.GetUserOrganizations(accountName, ctx.GetNextToken)
			if err != nil {
				fmt.Printf("error getting user's orgs :(\n")
			}
			orgsStr := make([]string, 0, len(orgs))
			for i := range orgs {
				orgsStr = append(orgsStr, orgs[i].Login)
			}
			gu.Orgs = orgsStr
		}

		ngu, err := store.Insert(gu)
		if err != nil {
			fmt.Printf("some error happened saving github user: %v", err)
		} else {
			totalRepos := gu.PublicRepos + gu.TotalPrivateRepos
			uts := string(gu.UserType)
			repos, err := services.GetGithubRepos(accountName, uts, ctx.GetNextToken, totalRepos)
			if err != nil {
				fmt.Printf("error getting github repos: %v\n", err)
			}
			for _, i := range repos {
				repo := i.ToRepo()
				repo.RepoOwnerID = ngu.ID
				*repoChannel <- *repo
			}
		}
	}
}

//func ProcessUserOrgs(u2u *chan string, orgs []*gh_user.Organization, store *gh_user.MongoStore) {
//	for i := range orgs {
//		fmt.Printf("org: %v\n", orgs[i].Login)
//		count := store.GetOrgByGithubNameCount(orgs[i].Login)
//		if count == 0 {
//			fmt.Printf("count: %v\n", count)
//			*u2u <- orgs[i].Login
//		}
//	}
//}
