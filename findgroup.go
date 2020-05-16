package main

import (
	"fmt"
	"log"
	"os"
	"os/user"
)

func main() {
	args := os.Args
	var u *user.User
	var err error
//Get current user if no username is provided.
	if len(args) == 1 {
		u, err = user.Current()
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		username := args[1]
		u, err = user.Lookup(username)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	gids, err := u.GroupIds()
	handle(err)
	for _, gid := range gids {
		group, err := user.LookupGroup(gid)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Group is  %s and groupID is %s", group.Name, group.Gid)
	}

}
