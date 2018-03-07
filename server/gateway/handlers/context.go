package handlers

import (
	"golang.org/x/oauth2"
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
	"fmt"
)

const fudge = 10000
const limit = 5000

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
	ctx.TokenIndex = ctx.TokenIndex + 1
	idx := ctx.TokenIndex % len(ctx.Tokens)
	if ctx.TokenIndex % (limit * len(ctx.Tokens) - fudge) == 0 {
		fmt.Printf("waiting an hour\n")
		time.Sleep(time.Hour + 5 * time.Minute)
	}
	ctx.Mutex.Unlock()
	return ctx.Tokens[idx]
}
