package api_service

import (
	"net/http"
	"github.com/andyblueyo/gitskill/server/gateway/models"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

const githubApiBase = "https://api.github.com"
const orgsResource = "orgs"
const currentUserResource = "user"
const usersResource = "users"
const reposResource = "repos"
const languagesResource = "languages"
const defaultLimit = 100
const defaultLanguageLimit = 50 // if a proj has over 50 languages lul
var limit = fmt.Sprintf("&per_page=%v", defaultLimit)
const ownerQuery = "&affiliation=owner"

func GetGithubUser(accountName, accountType, token string) (*models.User, error) {
	gu := &models.GithubUser{}

	url := ""

	switch accountType {
	case models.GHTypeUser:
		url = getUserAccountUrl(accountName, token)
		break
	case models.GHTypeOrganization:
		url = getOrgAccountUrl(accountName, token)
		break
	default:
		return nil, fmt.Errorf("need to pass in a recognized account type")
	}

	if err := fetchAndDecodeJSON(url, gu); err != nil {
		return nil, fmt.Errorf("ran into an error getting data: %v", err)
	}

	u := gu.ToUser()

	return u, nil
}

func GetGithubRepos(accountName, accountType, token string, max int) ([]*models.Repo, error) {
	url := ""
	fmt.Printf("max: %v\n", max)
	ghJsonTotal := make([]models.GithubRepo, 0, max)
	currentPage := 0
	for {
		if currentPage * defaultLimit > max {
			break
		}
		currentPage += 1
		switch accountType {
		case models.GHTypeUser:
			url = getCurrentUserReposUrl(token, currentPage)
			fmt.Printf("url: %v\n", url)
			break
		case models.GHTypeOrganization:
			url = getOrgReposUrl(accountName, token, currentPage)
			break
		default:
			return nil, fmt.Errorf("need to pass in a recognized account type")
		}
		remainderOrHundred := defaultLimit
		if currentPage * defaultLimit > max {
			remainderOrHundred = max % defaultLimit
		}
		fmt.Printf("remainderOrHundred: %v\n", remainderOrHundred)
		ghJson := make([]models.GithubRepo, 0, remainderOrHundred)
		if err := fetchAndDecodeJSON(url, &ghJson); err != nil {
			return nil, fmt.Errorf("ran into an error getting data: %v", err)
		}
		ghJsonTotal = append(ghJsonTotal, ghJson...)
	}

	repos := make([]*models.Repo, 0, max)

	for i := range ghJsonTotal {

		currentRepo := ghJsonTotal[i].ToRepo()

		// is 50 languages sufficient? meh
		languages := make([]*models.Language, 0, defaultLanguageLimit)
		languagesRaw := map[string]int{}
		languagesUrl := getLanguageUrlForRepo(currentRepo.FullName, token)

		if err := fetchAndDecodeJSON(languagesUrl, &languagesRaw); err != nil {
			return nil, fmt.Errorf("ran into an error getting data: %v", err)
		}

		for k, v := range languagesRaw {
			l := &models.Language{
				Name:  k,
				Lines: v,
			}
			languages = append(languages, l)
		}

		currentRepo.Languages = languages

		repos = append(repos, currentRepo)
	}

	// Get repo languages
	return repos, nil
}

func fetchAndDecodeJSON(url string, item interface{}) error {
	fmt.Printf("url: %v\n", url)
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

func getUserAccountUrl(user, token string) string {
	return githubApiBase + "/" + usersResource + "/" + user + "?access_token=" + token
}

func getOrgAccountUrl(org, token string) string {
	return githubApiBase + "/" + orgsResource + "/" + org + "?access_token=" + token
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