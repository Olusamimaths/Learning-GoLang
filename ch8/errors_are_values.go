package main

import (
	"errors"
	"fmt"
)

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
}

// implicitly implement the error interface
func (se StatusErr) Error() string {
	return se.Message
}

func login(uid, pwd string) error {
	return errors.New("no pwd")
}

func getData(file string) ([]byte, error) {
	return []byte(""), errors.New("no file")
}

// even when using custom errors, return error interface
func LoginAndGetData(uid, pwd, file string) ([]byte, error) {
	err := login(uid, pwd)
	if err != nil {
		return nil, StatusErr{
			Status: InvalidLogin,
			Message: fmt.Sprintf("invalid credentials for %s", uid),
		}
	}

	data, err := getData(file)
	if err != nil {
		return nil, StatusErr{
			Status: NotFound,
			Message: fmt.Sprintf("file %s not found", file),
		}
	}

	return data, nil
}

// returnend genErr is never nil!
// the type of the interfaces is not nil and for an interface to be considered nil
// its type and value must be nil!
func GenerateError(flag bool) error {
	var genErr StatusErr
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

// FIX 1: Return nil explicitly
func GenerateErrorFix1(flag bool) error {
	if flag {
		return StatusErr{
			Status: NotFound,
		}
	}
	return nil
}

// FIX 2: define returned instance to be of tye error
func GenerateErrorFix2(flag bool) error {
var genErr error // 
	if flag {
		genErr = StatusErr{
			Status: NotFound,
		}
	}
	return genErr
}

func main() {
	err := GenerateError(true)
	fmt.Println(err != nil)
	err = GenerateError(false)
	fmt.Println(err != nil)
}
