package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func main() {

	var gitOutPut []string

	cmd := exec.Command("git", "for-each-ref", "--count=10", "--sort=-committerdate", "refs/heads/", "--format='%(HEAD) %(color:yellow)%(refname:short)%(color:reset) - %(color:red)%(objectname:short)%(color:reset) - %(contents:subject) - %(authorname) (%(color:green)%(committerdate:relative)%(color:reset))'")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	in := bufio.NewScanner(stdout)

	for in.Scan() {
		gitOutPut = append(gitOutPut, strings.Trim(in.Text(), "' "))
	}
	if err := in.Err(); err != nil {
		log.Printf("error: %s", err)
	}

	prompt := promptui.Select{
		Label: "Select Branch",
		Items: gitOutPut,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	splitString := strings.Split(result, " - ")

	// Remove the * from current branch
	var re = regexp.MustCompile(`[*]* +`)
	branch := re.ReplaceAllString(splitString[0], "$1")

	switchBranch := exec.Command("git", "checkout", branch)
	if errors.Is(switchBranch.Err, exec.ErrDot) {
		cmd.Err = nil
	}
	if err := switchBranch.Run(); err != nil {
		log.Fatal(err)
	}

}
