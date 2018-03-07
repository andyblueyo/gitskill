package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/andyblueyo/gitskill/server/gateway/handlers"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"github.com/patrickmn/go-cache"
	"time"
	"github.com/andyblueyo/gitskill/server/gateway/models/gh-user"
	"github.com/andyblueyo/gitskill/server/gateway/models/gh-repo"
	"github.com/andyblueyo/gitskill/server/gateway/jobs"
	"gopkg.in/mgo.v2"
	"strings"
	"sync"
)

const accountPath = "/account"

func requireEnv(name string) string {
	val := os.Getenv(name)
	if len(val) == 0 {
		log.Fatalf("please set the %s environment variable", name)
	}
	return val
}

func main() {

	//and use that as the address this server will listen on
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	token := requireEnv("GHTOKENS")
	tokens := strings.Split(token, ",")
	clientID := requireEnv("GIT_CLIENT_ID")
	clientSecret := requireEnv("GIT_CLIENT_SECRET")
	dbAddr := requireEnv("DBADDR")
	tlsKey := requireEnv("TLSKEY")
	tlsCert := requireEnv("TLSCERT")

	//Use the DBADDR to dial MongoDB server
	sess, err := mgo.Dial(dbAddr)
	if err != nil {
		log.Printf("Error dialing Mongo: %v", err)
		os.Exit(1)
	}

	reposStore := gh_repo.NewMongoStore(sess, "git", "repos")
	usersStore := gh_user.NewMongoStore(sess, "git", "users")

	usersToScrape := make(chan string)
	usersToUsers := make(chan string)
	ctx := &handlers.HandlerContext{
		OauthConfig: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       []string{"read:user"},
			RedirectURL:  "https://" + addr + apiReply,
			Endpoint:     github.Endpoint,
		},
		StateCache:    cache.New(5*time.Minute, 10*time.Second),
		Tokens:        tokens,
		TokenIndex:    0,
		AccountsQueue: usersToScrape,
		Mutex:         &sync.RWMutex{},
	}

	reposToScrape := make(chan gh_repo.Repo)
	reposWithLanguages := make(chan gh_repo.Repo)
	orgsToScrapeForMembers := make(chan string)

	go jobs.ListenForAccounts(&usersToScrape, &reposToScrape, &orgsToScrapeForMembers, ctx, usersStore)
	go jobs.ListenForUsersToUsers(&usersToScrape, &usersToUsers)
	go jobs.ListenForRepos(&reposToScrape, &reposWithLanguages, ctx)
	go jobs.ListenForReposWithLanguages(&reposWithLanguages, reposStore)
	go jobs.ListenForOrgToScrapeMembers(&orgsToScrapeForMembers, &usersToScrape, ctx)

	mux := http.NewServeMux()

	//create a new handlers.CityHandler struct
	//since that is in a different package, use the
	//package name as a prefix, and import the package above
	//add the handler to the mux using .Handle() instead
	//of .HandleFunc(). The former is used for structs that
	//implement the http.Handler interface, while the latter
	//is used for simple functions that conform tos the
	//http.HandlerFunc type.
	//see https://drstearns.github.io/tutorials/goweb/#sechandlers
	mux.HandleFunc(apiSignIn, ctx.OAuthSignInHandler)
	mux.HandleFunc(apiReply, ctx.OAuthReplyHandler)

	mux.HandleFunc(accountPath, ctx.AccountHandler)

	fmt.Printf("server is listening at https://%s\n", addr)
	//log.Fatal(http.ListenAndServe(addr, mux))
	log.Fatal(http.ListenAndServeTLS(addr, tlsCert, tlsKey, mux))
}
