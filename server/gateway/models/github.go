package models

import "gopkg.in/mgo.v2/bson"

type GithubUserType string

const (
	GHTypeUser         = "User"
	GHTypeOrganization = "Organization"
)

type User struct {
	ID                bson.ObjectId  `json:"id" bson:"_id"`
	GithubUsername    string         `json:"githubUsername"`
	UserType          GithubUserType `json:"userType"`
	WebsiteUrl        string         `json:"websiteUrl"`
	Email             string         `json:"email"`
	PublicRepos       int            `json:"publicRepos"`
	OwnedPrivateRepos int            `json:"ownedPrivateRepos"`
	TotalPrivateRepos int            `json:"totalPrivateRepos"`
	AvatarUrl         string         `json:"avatarUrl"`
}

type Language struct {
	Name  string `json:"name"`
	Lines int    `json:"lines"`
}

type Library struct {
	Name         string `json:"name"`
	LanguageName string `json:"languageName"`
}

type Repo struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	RepoOwnerID bson.ObjectId `json:"repoOwnerId"`
	FullName    string        `json:"fullName"`
	OwnerType   string        `json:"ownerType"`
	Name        string        `json:"name"`
	Url         string        `json:"url"`
	Languages   []*Language   `json:"languages"`
	Libraries   []*Library    `json:"libraries"`
}

// For public Repos from Github
type GithubOwner struct {
	AvatarUrl string `json:"avatar_url"`
	Login     string `json:"login"`
	Type      string `json:"type"`
}

type GithubRepo struct {
	Name     string      `json:"name"`
	FullName string      `json:"full_name"`
	Url      string      `json:"html_url"`
	Owner    GithubOwner `json:"owner"`
}

// For users that logged into the app
type GithubUser struct {
	Login             string         `json:"login"`
	UserType          GithubUserType `json:"type"`
	Email             string         `json:"email"`
	PublicRepos       int            `json:"public_repos"`
	TotalPrivateRepos int            `json:"total_private_repos"`
	OwnedPrivateRepos int            `json:"owned_private_repos"`
	AvatarUrl         string         `json:"avatar_url"`
}

func (gu *GithubUser) ToUser() *User {
	u := User{
		GithubUsername:    gu.Login,
		UserType:          gu.UserType,
		Email:             gu.Email,
		PublicRepos:       gu.PublicRepos,
		TotalPrivateRepos: gu.TotalPrivateRepos,
		OwnedPrivateRepos: gu.OwnedPrivateRepos,
		AvatarUrl:         gu.AvatarUrl,
	}

	return &u
}

func (gr *GithubRepo) ToRepo() *Repo {
	r := Repo{
		FullName:  gr.FullName,
		OwnerType: gr.Owner.Type,
		Name:      gr.Name,
		Url:       gr.Url,
	}

	return &r
}
