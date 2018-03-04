package jobs

import "fmt"

func ListenForUsersToUsers(users *chan string, usersToUsers *chan string) {
	for {
		select {
		case user := <-*usersToUsers:
			*users <- user
			fmt.Printf("adding user %s to the users chan\n", user)
		default: // nothing
		}
	}
}
