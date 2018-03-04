package gh_repo

import (
	"gopkg.in/mgo.v2/bson"
)

type Store interface {
	GetByID(id bson.ObjectId) (*Repo, error)

	GetByGithubRepoName(name string) (*Repo, error)

	GetByGithubUser(name string) ([]*Repo, error)

	Insert(repo *Repo) error

	Update(repoUpdate *RepoUpdates) (*Repo, error)
}
