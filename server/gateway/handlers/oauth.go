package handlers

import (
	"encoding/base64"
	"crypto/rand"
	"net/http"
	"github.com/patrickmn/go-cache"
	"fmt"
	"golang.org/x/oauth2"
	"io"
)

//newStateValue returns a base64-encoded crypto-random value
//suitable for using as the `state` parameter in an OAuth2
//authorization request
func newStateValue() string {
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		panic("error generating random bytes")
	}
	return base64.URLEncoding.EncodeToString(buf)
}

//OAuthSignInHandler handles requests for the oauth sign-on API
func (ctx *HandlerContext) OAuthSignInHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: implement this handler by:
	// - generating a new state value
	// - adding it to the cache (default timeout)
	// - redirecting the client to the authorization URL
	//   returned from the OAuth config
	state := newStateValue()
	ctx.StateCache.Add(state, nil, cache.DefaultExpiration)
	redirURL := ctx.OauthConfig.AuthCodeURL(state)
	http.Redirect(w, r, redirURL, http.StatusSeeOther)
}

//OAuthReplyHandler handles requests made after authenticating
//with the OAuth provider, and authorizing our application
func (ctx *HandlerContext) OAuthReplyHandler(w http.ResponseWriter, r *http.Request) {
	//This handler is called after the OAuth provider redirects the client
	//back to our server. The query string may contain either these parameters:
	// - code = authorization code
	// - state = state value we sent to the server
	//OR these params if there was an error:
	// - error = an error code: https://tools.ietf.org/html/rfc6749#section-4.1.2.1
	// - error_description (optional) = human-readable error message
	// - error_uri (optional) = human-readable web page

	//TODO: implement this handler by doing the following:
	// - if the query string contains an "error" parameter, handle the error
	// - if the "state" query string param is missing or is not found in
	//   the cache, respond with an error
	// - if it is found, delete it from the cache so that it can't be used again
	// - use the `.Exchange()` method on the OAuth config to get an access token
	// - use the token to get a new http.Client you can use to make requests on
	//   behalf of the authenticated user
	// - use that client to get the user's profile (see constants above)
	qsParams := r.URL.Query()
	if len(qsParams.Get("error")) > 0 {
		errorDescription := qsParams.Get("error_description")
		if len(errorDescription) == 0 {
			errorDescription = "Error signing in: " + qsParams.Get("error")
		}
		http.Error(w, fmt.Sprintf("error signing in: %s", errorDescription), http.StatusInternalServerError)
		return
	}

	stateReturned := qsParams.Get("state")
	if _, found := ctx.StateCache.Get(stateReturned); !found {
		http.Error(w, "invalid state value returned from OAuth Provider", http.StatusBadRequest)
		return
	}

	ctx.StateCache.Delete(stateReturned)
	token, err := ctx.OauthConfig.Exchange(oauth2.NoContext, qsParams.Get("code"))
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting access token: %v", err), http.StatusInternalServerError)
		return
	}
	client := ctx.OauthConfig.Client(oauth2.NoContext, token)
	profileRequest, _ := http.NewRequest(http.MethodGet, githubCurrentUserAPI, nil)
	profileRequest.Header.Add(headerAccept, acceptGitHubV3JSON)
	profileResponse, err := client.Do(profileRequest)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting profile:  %v", err), http.StatusInternalServerError)
		return
	}
	defer profileResponse.Body.Close()

	//After obtaining the current user's profile, this is where you
	//would typically create a new User record in your system,
	//and begin a new authenticated Session for that user.
	//For purposes of this demo, we will just stream the profile
	//to the client so that we can see what it contains
	w.Header().Add(headerContentType, profileResponse.Header.Get(headerContentType))
	io.Copy(w, profileResponse.Body)
}

