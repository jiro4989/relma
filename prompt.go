package main

import (
	"github.com/manifoldco/promptui"
)

func PromptYesNo(msg string) (bool, error) {
	p := promptui.Select{
		Label: msg,
		Items: []string{"yes", "no"},
	}
	_, result, err := p.Run()
	if err != nil {
		return false, err
	}
	return result == "yes", nil
}
