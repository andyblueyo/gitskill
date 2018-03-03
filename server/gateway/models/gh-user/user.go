package gh_user

import "gopkg.in/mgo.v2/bson"

type GithubUserType string

const (
	GHTypeUser         = "User"
	GHTypeAuthUser     = "AuthUser"
	GHTypeOrganization = "Organization"
)

type User struct {
	ID                bson.ObjectId  `json:"id" bson:"_id"`
	GithubUsername    string         `json:"githubUsername" bson:"githubUsername"`
	UserType          GithubUserType `json:"userType" bson:"userType"`
	WebsiteUrl        string         `json:"websiteUrl" bson:"websiteUrl"`
	Email             string         `json:"email" bson:"email"`
	PublicRepos       int            `json:"publicRepos" bson:"publicRepos"`
	OwnedPrivateRepos int            `json:"ownedPrivateRepos" bson:"ownedPrivateRepos"`
	TotalPrivateRepos int            `json:"totalPrivateRepos" bson:"totalPrivateRepos"`
	AvatarUrl         string         `json:"avatarUrl" bson:"avatarUrl"`
	Orgs              []string       `json:"orgs,omitempty" bson:"orgs,omitempty"`
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

type Organization struct {
	Login string `json:"login"`
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
