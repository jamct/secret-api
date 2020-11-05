package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/zalando/go-keyring"
	"golang.org/x/crypto/ssh/terminal"
)

func login(server string) error {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Username: ")
	user, _ := reader.ReadString('\n')
	user = strings.Replace(user, "\n", "", -1)
	if len(user) < 2 {
		fmt.Println("please enter a username")
		return nil
	}

	service := "secret-api"
	pass := ""

	// get password from keychain
	pass, err := keyring.Get(service, user)
	if err != nil {
		//not found in keychain
		fmt.Print("Enter Password: ")
		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("no password entered.")
		}
		pass = string(bytePassword)
		//set password
		err = keyring.Set(service, user, pass)
		if err != nil {
			fmt.Println("could not save to keychain.")
			return nil
		}
	}

	fmt.Println("we have a password")

	//do http request here

	return nil
}
