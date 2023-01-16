package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("git for-each-ref --count=10 --sort=-committerdate refs/heads/ --format='%(HEAD) %(color:yellow)%(refname:short)%(color:reset) - %(color:red)%(objectname:short)%(color:reset) - %(contents:subject) - %(authorname) (%(color:green)%(committerdate:relative)%(color:reset))'")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmd.String())

	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
			"Saturday", "Sunday"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}
