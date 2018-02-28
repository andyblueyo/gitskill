package handlers

import (
	"golang.org/x/oauth2"
	"github.com/patrickmn/go-cache"
)

//HandlerContext is the receiver for our handler methods
//and contains various global values our handlers will need
type HandlerContext struct {
	//oauthConfig is the OAuth configuration for GitHub
	OauthConfig *oauth2.Config
	//stateCache is a cache of previously-generated OAuth state values
	StateCache    *cache.Cache
	Token         string
	AccountsQueue chan string
}

