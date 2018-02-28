package gh_repo

import "gopkg.in/mgo.v2/bson"

type Language struct {
	Name  string `json:"name"`
	Lines int    `json:"lines"`
}

type Library struct {
	Name         string `json:"name" bson:"name"`
	LanguageName string `json:"languageName" bson:"languageName"`
}

type Repo struct {
	ID          bson.ObjectId `json:"id" bson:"_id"`
	RepoOwnerID bson.ObjectId `json:"repoOwnerId" bson:"repoOwnerId"`
	FullName    string        `json:"fullName" bson:"fullName"`
	OwnerName   string        `json:"ownerName" bson:"ownerName"`
	OwnerType   string        `json:"ownerType" bson:"ownerType"`
	Name        string        `json:"name" bson:"name"`
	Url         string        `json:"url" bson:"url"`
	Languages   []*Language   `json:"languages" bson:"languages"`
	Libraries   []*Library    `json:"libraries,omitempty" bson:"libraries,omitempty"`
	Stars       int           `json:"stars" bson:"stars"`
	Description string        `json:"description" bson:"description"`
	Forks       int           `json:"forks" bson:"forks"`
	UpdatedAt   string        `json:"updatedAt" bson:"updatedAt"`
}

// For public Repos from Github
type GithubOwner struct {
	AvatarUrl string `json:"avatar_url"`
	Login     string `json:"login"`
	Type      string `json:"type"`
}

type GithubRepo struct {
	Name        string      `json:"name"`
	FullName    string      `json:"full_name"`
	Url         string      `json:"html_url"`
	Owner       GithubOwner `json:"owner"`
	Stars       int         `json:"stargazers_count"`
	Description string      `json:"description"`
	Forks       int         `json:"forks_count"`
	UpdatedAt   string      `json:"updated_at"`
}

func (gr *GithubRepo) ToRepo() *Repo {
	r := Repo{
		FullName:    gr.FullName,
		OwnerName:   gr.Owner.Login,
		OwnerType:   gr.Owner.Type,
		Name:        gr.Name,
		Url:         gr.Url,
		Stars:       gr.Stars,
		Description: gr.Description,
		Forks:       gr.Forks,
		UpdatedAt:   gr.UpdatedAt,
	}

	return &r
}

type RepoUpdates struct {
	Languages []*Language `json:"languages"`
	Libraries []*Library  `json:"libraries"`
}
