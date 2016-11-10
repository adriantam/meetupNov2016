package main

type User struct {
	MyProvider Provider
}

func (user *User) ProcessInfo() (Information, error) {
	info, err := user.MyProvider.GetInfo()
	/* Do some work with the information */

	return info, err
}
