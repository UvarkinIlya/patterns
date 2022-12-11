package main

import "os"

type ObjectAccess interface {
	access(fileName string) (file *os.File, err error)
}

type SomeObjectAccess struct {
	ObjectAccess
}

func (accessor SomeObjectAccess) access(fileName string) (file *os.File, err error) {
	file, err = os.Open(fileName)
	if err != nil {
		return
	}

	return
}
