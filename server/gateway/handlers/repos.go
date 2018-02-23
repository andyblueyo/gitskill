package handlers

import (
	"net/http"
	"strings"
	"encoding/json"
	"github.com/andyblueyo/gitskill/server/gateway/api-service"
	"fmt"
)

type RepoHandler struct {
	Token string
}

func (rh *RepoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	accountName := r.URL.Query().Get("account")
	accountType := r.URL.Query().Get("type")

	accountName = strings.ToLower(accountName)

	if len(accountName) == 0 {
		http.Error(w, "please provide an accountName name", http.StatusBadRequest)
		return
	}

	//add the header `Content-Type: application/json`
	w.Header().Add(headerContentType, contentTypeJSON)
	//add the CORS header `Access-Control-Allow-Origin: *`
	w.Header().Add(headerAccessControlAllowOrigin, "*")
	account, err := api_service.GetGithubUser(accountName, accountType, rh.Token)
	if err != nil {
		http.Error(w, fmt.Sprintf("an error occured: %v", err), http.StatusBadRequest)
	}
	totalRepos := account.PublicRepos + account.TotalPrivateRepos
	fmt.Printf("total repos: %v\n", totalRepos)
	repos, err := api_service.GetGithubRepos(accountName, accountType, rh.Token, totalRepos)
	if err != nil {
		http.Error(w, fmt.Sprintf("an error occured: %v", err), http.StatusBadRequest)
	}

	//u, err := api_service.GetGithubUser("airbnb", models.GHTypeOrganization)
	//if err != nil {
	//	fmt.Printf("err: %v\n", err)
	//}
	//r, err := api_service.GetGithubRepos("airbnb", models.GHTypeOrganization, 0)
	//if err != nil {
	//	fmt.Printf("err %v\n", err)
	//}
	//
	//fmt.Printf("users")
	//fmt.Printf("%+v\n", u)
	//fmt.Printf("repos")
	//fmt.Printf("num repos: %v\n", len(r))
	//for i := range r {
	//	fmt.Printf("%+v\n", r[i])
	//}

	fmt.Printf("length of repos: %v\n", len(repos))
	json.NewEncoder(w).Encode(repos)
}
