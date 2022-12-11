package main

import (
	"fmt"
	"os"
	"os/user"
)

type SomeObjectAccessProxy struct {
	someObjectAccess SomeObjectAccess
	ObjectAccess
}

func (accessorProxy SomeObjectAccessProxy) access(fileName string) (file *os.File, err error) {
	currentUser, err := user.Current()
	if err != nil {
		return
	}

	if !isAdmin(currentUser.Name) {
		return nil, fmt.Errorf("not permission")
	}

	file, err = accessorProxy.someObjectAccess.access(fileName)
	if err != nil {
		return
	}

	return
}

func isAdmin(userName string) bool {
	return userName == "Илья Уваркин"
}
