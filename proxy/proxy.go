package proxy

import (
	// "errors"
	"fmt"
)

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type User struct {
	ID int32
}

type UserList []User

type UserListProxy struct {
	SomeDatabase              *UserList
	StackCapacity             int
	StackCache                UserList
	DidDidLastSearchUsedCache bool
}

func (u *UserList) FindUser(id int32) (User, error) {
	for i := 0; i < len(*u); i++ {
		if (*u)[i].ID == id {
			return (*u)[i], nil
		}
	}
	return User{}, fmt.Errorf("User %d could not be found\n", id)
}

func (u *UserListProxy) FindUser(id int32) (User, error) {
	user, err := u.StackCache.FindUser(id)
	if err == nil {
		fmt.Println("Returning user from cache")
		u.DidDidLastSearchUsedCache = true
		return user, nil
	}
	user, err = u.SomeDatabase.FindUser(id)
	if err != nil {
		return User{}, err
	}
	fmt.Println("Returning user from database")
	u.DidDidLastSearchUsedCache = false
	u.addUserToStack(user)
	return user, nil
}

func (u *UserListProxy) addUserToStack(user User) {
	if len(u.StackCache) >= u.StackCapacity {
		u.StackCache = append(u.StackCache[1:], user)
	} else {
		u.StackCache.addUser(user)
	}
}
func (u *UserList) addUser(newUser User) {
	*u = append(*u, newUser)
}
