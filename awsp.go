package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/aws/aws-sdk-go-v2/config"
	"gopkg.in/ini.v1"
)

func check(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func getLocalAwsProfiles() (list []string, err error) {
	fname := config.DefaultSharedConfigFilename()
	f, err := ini.Load(fname)
	if err != nil {
		return list, err
	}

	list = append(list, "default")
	for _, v := range f.Sections() {
		if len(v.Keys()) != 0 { // Get only the sections having Keys
			parts := strings.Split(v.Name(), " ")
			if len(parts) == 2 && parts[0] == "profile" { // skip default
				list = append(list, parts[1])
			}
		}
	}

	return
}

func writeAwspFile(profile string) (err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	awspFile := path.Join(homeDir, ".awsp")
	err = os.WriteFile(awspFile, []byte(profile), 0644)
	if err != nil {
		return err
	}

	return
}

func main() {
	fmt.Println("AWS Profile Switcher")

	profiles, err := getLocalAwsProfiles()
	check(err)

	currentProfileIndex := slices.Index(profiles, os.Getenv("AWS_PROFILE"))

	if currentProfileIndex == -1 {
		currentProfileIndex = 0
	}

	if len(profiles) == 0 {
		log.Println("No profiles found.")
		log.Println("Refer to this guide for help on setting up a new AWS profile:")
		log.Println("https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html")
		os.Exit(1)
	}

	prompt := &survey.Select{
		Message: "Choose a profile:",
		Options: profiles,
		Default: currentProfileIndex,
	}

	var selection string
	err = survey.AskOne(prompt, &selection, survey.WithPageSize(7))
	check(err)

	writeAwspFile(selection)
}
