package controllers

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

//ValidateUserInput -- used to Validate the JSON supplied by the user
func ValidateUserInput(userInput map[string]string, keys []string) (err error) {

	for _, key := range keys {
		if strings.EqualFold(userInput[key], "") {
			err = fmt.Errorf("%s should not be an empty string", key)
			log.Println(err)
			return
		}
		if strings.HasSuffix(key, "id") {
			tempt, err := strconv.Atoi(userInput[key])
			if err != nil {
				err = fmt.Errorf("%s needs to be a string representation of a postive integer", key)
				log.Println(err)
				return err
			}
			if tempt <= 0 {
				err = fmt.Errorf("%s should not be less than zero", key)
				log.Println(err)
				return err
			}
		}
	}
	return
}
