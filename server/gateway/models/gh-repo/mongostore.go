package gh_repo

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

var ErrRepoNotFound = errors.New("error: repo not found")

//MongoStore implements Store for MongoDB
type MongoStore struct {
	//mongo session
	session *mgo.Session
	//database name
	dbname string
	//collection name
	colname string
	//Collection object
	col *mgo.Collection
}

type UpdatedRepo struct {
	FullName    string      `json:"fullName" bson:"fullName"`
	OwnerName   string      `json:"ownerName" bson:"ownerName"`
	OwnerType   string      `json:"ownerType" bson:"ownerType"`
	Name        string      `json:"name" bson:"name"`
	Url         string      `json:"url" bson:"url"`
	Languages   []*Language `json:"languages" bson:"languages"`
	Libraries   []*Library  `json:"libraries,omitempty" bson:"libraries,omitempty"`
	Stars       int         `json:"stars" bson:"stars"`
	Description string      `json:"description" bson:"description"`
	Forks       int         `json:"forks" bson:"forks"`
	UpdatedAt   string      `json:"updatedAt" bson:"updatedAt"`
}

//NewMongoStore creates a new Mongo store
func NewMongoStore(sess *mgo.Session, dbName string, colName string) *MongoStore {
	return &MongoStore{
		session: sess,
		dbname:  dbName,
		colname: colName,
		col:     sess.DB(dbName).C(colName),
	}
}

func (s *MongoStore) Insert(repo *Repo) (*Repo, error) {
	existing, err := s.getByFieldCaseInsensitive("fullName", repo.FullName)
	if err != nil {
		id := bson.NewObjectId()
		repo.ID = id
		if err := s.col.Insert(repo); err != nil {
			return nil, fmt.Errorf("error inserting repo: %v", err)
		}
		return repo, nil
	} else {
		change := mgo.Change{
			Update: bson.M{
				"$set": &UpdatedRepo{
					FullName:    repo.FullName,
					OwnerName:   repo.OwnerName,
					OwnerType:   repo.OwnerType,
					Name:        repo.Name,
					Url:         repo.Url,
					Languages:   repo.Languages,
					Libraries:   repo.Libraries,
					Stars:       repo.Stars,
					Description: repo.Description,
					Forks:       repo.Forks,
					UpdatedAt:   repo.UpdatedAt,
				}},
			ReturnNew: false,
		}
		updatedRepo := &Repo{}
		if _, err := s.col.FindId(existing.ID).Apply(change, updatedRepo); err != nil {
			return nil, fmt.Errorf("error updating repo: %v", err)
		}
		return repo, nil
	}
}

func (s *MongoStore) GetByID(id bson.ObjectId) (*Repo, error) {
	r := &Repo{}
	if err := s.col.FindId(id).One(r); err != nil {
		return nil, fmt.Errorf("error finding repo: %v", err)
	}
	return r, nil
}

func (s *MongoStore) GetByRepoName(name string) (*Repo, error) {
	r, err := s.getByField("name", name)
	if err != nil {
		return nil, fmt.Errorf("error finding repo: %v", err)
	}
	return r, nil
}

func (s *MongoStore) GetByGithubUser(ghUser string) ([]*Repo, error) {
	repoLimit := 100
	rs := make([]*Repo, 0, repoLimit)
	if err := s.col.Find(bson.M{"": ghUser}).Limit(repoLimit).All(rs); err != nil {
		return nil, ErrRepoNotFound
	}
	return rs, nil
}

func (s *MongoStore) Update(id bson.ObjectId, updates *RepoUpdates) (*Repo, error) {
	change := mgo.Change{
		Update:    bson.M{"$set": updates},
		ReturnNew: true,
	}

	col := s.session.DB(s.dbname).C(s.colname)
	updatedRepo := &Repo{}

	// update changes to UserSettings store using userSetting.ID
	_, err := col.FindId(id).Apply(change, updatedRepo)
	if err != nil {
		return nil, fmt.Errorf("error updating repo: %v", err)
	}

	return updatedRepo, nil
}

//GetByField returns the first user for which the given field
//has the given found
//Returns an error if no user matches the query
func (s *MongoStore) getByField(field string, value string) (*Repo, error) {
	r := &Repo{}
	if err := s.col.Find(bson.M{field: value}).One(r); err != nil {
		return nil, ErrRepoNotFound
	}
	return r, nil
}

func (s *MongoStore) getByFieldCaseInsensitive(field, value string) (*Repo, error) {
	u := &Repo{}
	if err := s.col.Find(bson.M{field: bson.M{"$regex": bson.RegEx{Pattern: value, Options: "i"}}}).One(u); err != nil {
		return nil, ErrRepoNotFound
	}
	return u, nil
}
