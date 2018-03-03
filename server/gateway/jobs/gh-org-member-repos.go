package jobs

import (
	"fmt"
	"github.com/andyblueyo/gitskill/server/gateway/services"
	"github.com/andyblueyo/gitskill/server/gateway/handlers"
)

func ListenForOrgToScrapeMembers(orgs *chan string, users *chan string, ctx *handlers.HandlerContext) {
	for {
		select {
		case org := <-*orgs:
			members, err := services.GetGithubMembers(org, ctx.GetNextToken)
			if err != nil {
				fmt.Printf("error getting members for org: %v with err: %v\n", org, err)
			} else {
				for i := range members {
					*users <- members[i]
				}
			}
		default: // nothing
		}
	}
}
