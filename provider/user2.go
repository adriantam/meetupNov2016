package main

type User2 struct {
	MyProvider ProviderInterface // Use the interface instead of provider
}

func (user *User2) ProcessInfo() (Information, error) {
	info, err := user.MyProvider.GetInfo()
	/* Do some work with the information */

	return info, err
}
