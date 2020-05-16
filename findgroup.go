package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
)
//handle function for handling the errors encountered.
func handle(err error) {
	if err != nil {
		log.Println(err)
	}
}
//Entry point for the code.
func main() {
	args := os.Args
	var u *user.User
	var err error
//Gets the current user if not specified.
	if len(args) == 1 {
		u, err = user.Current()
		handle(err)
		return
	} else {
		username := args[1]
		u, err = user.Lookup(username)
		if err != nil {
			log.Println(err)
			return
		}
	}

	gids, err := u.GroupIds()
	handle(err)
	for _, gid := range gids {
		group, err := user.LookupGroup(gid)
		handle(err)
		fmt.Printf("Group is  %s and groupID is %s", group.Name, group.Gid)
	}

}
