package gh_user

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

var ErrUserNotFound = errors.New("error: user not found")
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

type UpdatedUser struct {
	GithubUsername    string         `json:"githubUsername" bson:"githubUsername"`
	UserType          GithubUserType `json:"userType" bson:"userType"`
	Orgs              []string       `json:"orgs" bson:"orgs"`
	WebsiteUrl        string         `json:"websiteUrl" bson:"websiteUrl"`
	Email             string         `json:"email" bson:"email"`
	PublicRepos       int            `json:"publicRepos" bson:"publicRepos"`
	OwnedPrivateRepos int            `json:"ownedPrivateRepos" bson:"ownedPrivateRepos"`
	TotalPrivateRepos int            `json:"totalPrivateRepos" bson:"totalPrivateRepos"`
	AvatarUrl         string         `json:"avatarUrl" bson:"avatarUrl"`
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

func (s *MongoStore) GetByID(id bson.ObjectId) (*User, error) {
	u := &User{}
	if err := s.col.FindId(id).One(u); err != nil {
		return nil, fmt.Errorf("error finding user: %v", err)
	}
	return u, nil
}

func (s *MongoStore) GetByGithubName(name string) (*User, error) {
	u, err := s.getByFieldCaseInsensitive("githubUsername", name)
	if err != nil {
		return nil, fmt.Errorf("error finding user: %v", err)
	}
	return u, nil
}

func (s *MongoStore) GetOrgByGithubNameCount(name string) (int) {
	count, err := s.col.Find(
		bson.M{
			"githubUsername": bson.M{"$regex": bson.RegEx{Pattern: name, Options: "i"}},
		}).Count()
	if err != nil {
		return 0
	}
	return count
}

func (s *MongoStore) GetAllOrgs() ([]*User, error) {
	us, err := s.getManyByField("userType", GHTypeOrganization)
	if err != nil {
		return nil, fmt.Errorf("error finding users: %v", err)
	}
	return us, nil
}

func (s *MongoStore) Insert(user *User) (*User, error) {
	id := bson.NewObjectId()
	existing, err := s.getByFieldCaseInsensitive("githubUsername", user.GithubUsername)
	if err != nil {
		user.ID = id
		if err := s.col.Insert(user); err != nil {
			return nil, fmt.Errorf("error inserting user: %v", err)
		}
		return user, nil
	} else {
		change := mgo.Change{
			Update: bson.M{
				"$set": &UpdatedUser{
					GithubUsername:    user.GithubUsername,
					Orgs:              user.Orgs,
					UserType:          user.UserType,
					WebsiteUrl:        user.WebsiteUrl,
					Email:             user.Email,
					PublicRepos:       user.PublicRepos,
					OwnedPrivateRepos: user.OwnedPrivateRepos,
					TotalPrivateRepos: user.TotalPrivateRepos,
					AvatarUrl:         user.AvatarUrl,
				}},
			ReturnNew: false,
		}
		updatedUser := &User{}
		if _, err := s.col.FindId(existing.ID).Apply(change, updatedUser); err != nil {
			return nil, fmt.Errorf("error updating user: %v", err)
		}
		return updatedUser, nil
	}
}

//GetByField returns the first user for which the given field
//has the given found
//Returns an error if no user matches the query
func (s *MongoStore) getByField(field string, value string) (*User, error) {
	u := &User{}
	if err := s.col.Find(bson.M{field: value}).One(u); err != nil {
		return nil, ErrUserNotFound
	}
	return u, nil
}

//GetByField returns the first user for which the given field
//has the given found
//Returns an error if no user matches the query
func (s *MongoStore) getManyByField(field string, value string) ([]*User, error) {
	us := make([]*User, 0, 100)
	if err := s.col.Find(bson.M{field: value}).All(&us); err != nil {
		return nil, ErrUserNotFound
	}
	return us, nil
}

func (s *MongoStore) getByFieldCaseInsensitive(field string, value string) (*User, error) {
	u := &User{}
	if err := s.col.Find(bson.M{field: bson.M{"$regex": bson.RegEx{Pattern: value, Options: "i"}}}).One(u); err != nil {
		return nil, ErrUserNotFound
	}
	return u, nil
}
