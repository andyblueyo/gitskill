package handlers

import (
	"golang.org/x/oauth2"
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

//HandlerContext is the receiver for our handler methods
//and contains various global values our handlers will need
type HandlerContext struct {
	//oauthConfig is the OAuth configuration for GitHub
	OauthConfig *oauth2.Config
	//stateCache is a cache of previously-generated OAuth state values
	StateCache    *cache.Cache
	Tokens        []string
	TokenIndex    int
	AccountsQueue chan string
	Mutex         *sync.RWMutex
}

func (ctx *HandlerContext) GetNextToken() string {
	ctx.Mutex.Lock()
	idx := ctx.TokenIndex % len(ctx.Tokens)
	ctx.TokenIndex = ctx.TokenIndex + 1
	wait := 150 * time.Millisecond
	time.Sleep(wait)
	ctx.Mutex.Unlock()
	return ctx.Tokens[idx]
}
