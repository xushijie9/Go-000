package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func SearchUser() error {
	return errors.New("user nil")
}

func Request() error {
	return errors.Wrap(SearchUser(), "user nil")
}

func main() {
	err := Request()
	fmt.Printf("%+v", err)

}