package services

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"strings"
	"github.com/andyblueyo/gitskill/server/gateway/models/gh-repo"
	"github.com/andyblueyo/gitskill/server/gateway/models/gh-user"
)

const githubApiBase = "https://api.github.com"
const orgsResource = "orgs"
const currentUserResource = "user"
const usersResource = "users"
const membersResource = "members"
const reposResource = "repos"
const languagesResource = "languages"
const defaultLimit = 100
const defaultLanguageLimit = 50 // if a proj has over 50 languages lul
const ownerQuery = "&affiliation=owner"

type getToken func() string

var limit = fmt.Sprintf("&per_page=%v", defaultLimit)

func GetGithubUser(accountName string, getTokenFn getToken) (*gh_user.User, error) {
	gu := &gh_user.GithubUser{}
	token := getTokenFn()
	url := getUserAccountUrl(accountName, token)

	//switch accountType {
	//case gh_user.GHTypeUser:
	//	url = getUserAccountUrl(accountName, token)
	//	break
	//case gh_user.GHTypeOrganization:
	//	url = getOrgAccountUrl(accountName, token)
	//	break
	//default:
	//	return nil, fmt.Errorf("need to pass in a recognized account type")
	//}

	if err := fetchAndDecodeJSON(url, gu); err != nil {
		return nil, fmt.Errorf("ran into an error getting data: %v", err)
	}

	u := gu.ToUser()

	return u, nil
}

func GetGithubRepos(accountName, accountType string, getTokenFn getToken, max int) ([]*gh_repo.GithubRepo, error) {
	url := ""
	ghJsonTotal := make([]*gh_repo.GithubRepo, 0, max)
	currentPage := 0
	for {
		token := getTokenFn()
		if currentPage*defaultLimit > max {
			break
		}

		currentPage += 1

		switch strings.ToLower(accountType) {
		case strings.ToLower(gh_user.GHTypeUser):
			url = getUserReposUrl(accountName, token, currentPage)
			break
		case strings.ToLower(gh_user.GHTypeAuthUser):
			url = getCurrentUserReposUrl(token, currentPage)
			break
		case strings.ToLower(gh_user.GHTypeOrganization):
			url = getOrgReposUrl(accountName, token, currentPage)
			break
		default:
			return nil, fmt.Errorf("need to pass in a recognized account type")
		}

		remainderOrHundred := defaultLimit
		if currentPage*defaultLimit > max {
			remainderOrHundred = max % defaultLimit
		}

		ghJson := make([]*gh_repo.GithubRepo, 0, remainderOrHundred)
		if err := fetchAndDecodeJSON(url, &ghJson); err != nil {
			return nil, fmt.Errorf("ran into an error getting data: %v", err)
		}

		ghJsonTotal = append(ghJsonTotal, ghJson...)
	}

	// Get repo languages
	return ghJsonTotal, nil
}

func GetGithubMembers(accountName string, getTokenFn getToken) ([]string, error) {
	url := ""
	ghJsonTotal := make([]*gh_user.GithubUser, 0, 0)
	currentPage := 0
	for {
		token := getTokenFn()
		currentPage += 1
		url = getOrgMembersUrl(accountName, token, currentPage)

		ghJson := make([]*gh_user.GithubUser, 0, 100)
		if err := fetchAndDecodeJSON(url, &ghJson); err != nil {
			return nil, fmt.Errorf("ran into an error getting data: %v", err)
		}

		ghJsonTotal = append(ghJsonTotal, ghJson...)
		if len(ghJson) < 100 {
			break
		}
	}
	memberStringSlice := make([]string, 0, len(ghJsonTotal))
	for i := range ghJsonTotal {
		memberStringSlice = append(memberStringSlice, ghJsonTotal[i].Login)
	}
	// Get repo languages
	return memberStringSlice, nil
}

func GetUserOrganizations(user string, getTokenFn getToken) ([]*gh_user.Organization, error) {
	// TODO
	token := getTokenFn()
	orgs := make([]*gh_user.Organization, 0, 100)
	url := getUserOrgsUrl(user, token)
	if err := fetchAndDecodeJSON(url, &orgs); err != nil {
		return nil, fmt.Errorf("ran into an error getting user orgs: %v", err)
	}
	return orgs, nil
}

func GetRepoLanguage(repo *gh_repo.Repo, getTokenFn getToken) (*gh_repo.Repo, error) {

	// is 50 languages sufficient? meh
	languages := make([]*gh_repo.Language, 0, defaultLanguageLimit)
	languagesRaw := map[string]int{}
	token := getTokenFn()

	languagesUrl := getLanguageUrlForRepo(repo.FullName, token)

	if err := fetchAndDecodeJSON(languagesUrl, &languagesRaw); err != nil {
		return nil, fmt.Errorf("ran into an error getting data: %v", err)
	}

	for k, v := range languagesRaw {
		l := &gh_repo.Language{
			Name:  k,
			Lines: v,
		}
		languages = append(languages, l)
	}

	repo.Languages = languages
	return repo, nil
}

func fetchAndDecodeJSON(url string, item interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error getting url: %v", err)
	}

	if resp.StatusCode >= 300 {
		return fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	// Get org repos
	if err := json.Unmarshal(body, item); err != nil {
		return fmt.Errorf("error decoding json: %v", err)
	}
	return nil
}

// Url helper functions
func getUserReposUrl(user, token string, offset int) string {
	offsetQuery := offsetQuery(offset)
	return githubApiBase + "/" + usersResource + "/" + user + "/" + reposResource + "?access_token=" + token + limit + offsetQuery
}

func getCurrentUserReposUrl(token string, offset int) string {
	offsetQuery := offsetQuery(offset)
	return githubApiBase + "/" + currentUserResource + "/" + reposResource + "?access_token=" + token + limit + offsetQuery + ownerQuery
}

func getOrgReposUrl(org, token string, offset int) string {
	offsetQuery := offsetQuery(offset)
	return githubApiBase + "/" + orgsResource + "/" + org + "/" + reposResource + "?access_token=" + token + limit + offsetQuery
}

func getOrgMembersUrl(org, token string, offset int) string {
	offsetQuery := offsetQuery(offset)
	return githubApiBase + "/" + orgsResource + "/" + org + "/" + membersResource + "?access_token=" + token + limit + offsetQuery
}

func getUserAccountUrl(user, token string) string {
	return githubApiBase + "/" + usersResource + "/" + user + "?access_token=" + token
}

func getOrgAccountUrl(org, token string) string {
	return githubApiBase + "/" + orgsResource + "/" + org + "?access_token=" + token
}

func getUserOrgsUrl(user, token string) string {
	return githubApiBase + "/" + usersResource + "/" + user + "/" + orgsResource + "?access_token=" + token
}

func getLanguageUrlForRepo(fullRepoName, token string) string {
	return githubApiBase + "/" + reposResource + "/" + fullRepoName + "/" + languagesResource + "?access_token=" + token
}

func offsetQuery(offset int) string {
	offsetQuery := ""
	if offset > 0 {
		offsetQuery = fmt.Sprintf("&page=%v", offset)
	}
	return offsetQuery
}
