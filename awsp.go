package main

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/fatih/color"
	"gopkg.in/ini.v1"
)

func check(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func profileErrorExit(out string) {
	color.HiRed("%s\n\n", out)
	color.Red("Refer to this guide for help on setting up a new AWS profile:")
	color.Red("https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-configure.html")
	os.Exit(1)
}

func getLocalAwsProfiles() (list []string, err error) {
	fname := config.DefaultSharedConfigFilename()
	f, err := ini.Load(fname)
	if errors.Is(err, fs.ErrNotExist) {
		profileErrorExit("Default shared config file not found.")
	} else if err != nil {
		return list, err
	}

	list = append(list, "default")
	for _, v := range f.Sections() {
		parts := strings.Split(v.Name(), " ")
		if len(parts) == 2 && parts[0] == "profile" { // skip default
			list = append(list, parts[1])
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
	profiles, err := getLocalAwsProfiles()
	check(err)

	currentProfileIndex := slices.Index(profiles, os.Getenv("AWS_PROFILE"))

	if currentProfileIndex == -1 {
		currentProfileIndex = 0
	}

	if len(profiles) == 0 {
		profileErrorExit("No profiles found.")
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
