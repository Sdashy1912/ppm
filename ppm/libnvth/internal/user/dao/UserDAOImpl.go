package dao

import (
	"ppm/libnvth/internal/database"
	"ppm/libnvth/internal/user/bean"
	"gopkg.in/mgo.v2/bson"
)

// UserDAOImpl implements UserDAO
type UserDAOImpl struct {
	session *database.DBSession
}

var _ UserDAO = (*UserDAOImpl)(nil)

// NewUserDAOImpl return a new UserDAOImpl instance
func NewUserDAOImpl(session *database.DBSession) UserDAOImpl {
	return UserDAOImpl{session}
}

// List select all
func (dao UserDAOImpl) List() ([]bean.UserBean, error) {
	collection := dao.session.Collection("Users")
	defer collection.Close()
	users := []bean.UserBean{}
	err := collection.Find(nil).All(&users)
	return users, err
}

// Get find by id
func (dao UserDAOImpl) Get(ID bson.ObjectId) (bean.UserBean, error) {
	collection := dao.session.Collection("Users")
	defer collection.Close()
	user := bean.UserBean{}
	err := collection.FindId(ID).One(&user)
	return user, err
}

// GetByEmail get an user by email
func (dao UserDAOImpl) GetByEmail(email string) (bean.UserBean, error) {
	collection := dao.session.Collection("Users")
	defer collection.Close()
	user := bean.UserBean{}
	err := collection.Find(bson.M{"email": email}).One(&user)
	return user, err
}

// GetByEmailAndPassword get an user by username and password (for login)
func (dao UserDAOImpl) GetByEmailAndPassword(email string, password string) (bean.UserBean, error) {
	collection := dao.session.Collection("Users")
	defer collection.Close()
	user := bean.UserBean{}
	err := collection.Find(bson.M{"email": email, "password": password}).One(&user)
	return user, err
}

// Insert persist a new object
func (dao UserDAOImpl) Insert(param UserCreateDAOParam) error {
	collection := dao.session.Collection("Users")
	defer collection.Close()
	return collection.Insert(param)
}

// Update change details
func (dao UserDAOImpl) Update(param UserUpdateDAOParam) error {
	collection := dao.session.Collection("Users")
	defer collection.Close()
	return collection.UpdateId(param.ID, bson.M{"$set": bson.M{"name": param.Name}, "phone_number": param.PhoneNumber})
}

// Remove remove object by id
func (dao UserDAOImpl) Remove(ID bson.ObjectId) error {
	collection := dao.session.Collection("Users")
	defer collection.Close()
	return collection.RemoveId(ID)
}

// IsSafeToDelete check before delete an user
func (dao UserDAOImpl) IsSafeToDelete(ID bson.ObjectId) (bool, error) {
	collection := dao.session.Collection("Targets")
	defer collection.Close()
	result := []bson.M{}
	err := collection.Find(bson.M{"assignee_id": ID}).All(&result)
	if err != nil {
		return false, err
	}
	return len(result) == 0, nil
}
