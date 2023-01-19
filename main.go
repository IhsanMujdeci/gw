package main

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func main() {

	//var gitOutPut []string

	gitBranchCommand := exec.Command(
		"git",
		"for-each-ref",
		"--count=10",
		"--sort=-committerdate",
		"refs/heads/",
		"--format='%(HEAD) %(color:yellow)%(refname:short)%(color:reset) - %(color:red)%(objectname:short)%(color:reset) - %(contents:subject) - %(authorname) (%(color:green)%(committerdate:relative)%(color:reset))'",
	)

	gitOutPut, err := cmdRun(gitBranchCommand)
	if err != nil {
		log.Fatal(err)
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

	// Strip any stars and spaces after for current branch
	var re = regexp.MustCompile(`[*]* +`)
	branch := re.ReplaceAllString(splitString[0], "$1")

	switchBranch := exec.Command("git", "checkout", branch)

	switchBranchOutput, err := cmdRun(switchBranch)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Join(switchBranchOutput[:], "\n"))
}

func cmdRun(cmd *exec.Cmd) (output []string, err error) {
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return output, err
	}
	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		return output, err
	}
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		output = append(output, strings.Trim(in.Text(), "' "))
	}
	return output, in.Err()
}
