package cmd

import (
	"os/exec"
	"os"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(prCmd)
}

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Creates pull request",
	Long: `Create a pull request with custom template`,
	Run: func(cmd *cobra.Command, args []string){
		cwd,cwdErr := os.Getwd()
		if cwdErr != nil {
			fmt.Println("error", cwdErr)
			return
		}

		fmt.Println(cwd)
		os.Chdir(cwd)
		
		gs := exec.Command("git", "status")
		_,gsErr := gs.Output()
		if gsErr != nil {
			fmt.Println("error",gsErr)
			return
		}

		ga := exec.Command("git", "add", "-A")
		_,gaErr := ga.Output()
		if gaErr != nil {
			fmt.Println("error",gaErr)
			return
		}

		var commitMsg string 
		fmt.Println("Enter commit message:")
		fmt.Scanln(&commitMsg)

		gc := exec.Command("git", "commit", "-m",commitMsg)
		_,gcErr := gc.Output()
		if gcErr != nil {
			fmt.Println("error",gcErr)
			return
		}

		gpush := exec.Command("git", "push")
		_,gpushErr := gpush.Output()
		if gpushErr != nil {
			fmt.Println("error",gpushErr)
			return
		}

		var title string 
		var base string 
		var feat string 

		fmt.Println("Enter pr title:")
		fmt.Scanf("\n")
		fmt.Scanln(&title)

		fmt.Println("Enter base branch to pull pr into:")
		fmt.Scanf("\n")
		fmt.Scanln(&base)

		fmt.Println("Enter feat branch name to pull pr into:")
		fmt.Scanf("\n")
		fmt.Scanln(&feat)

		body := `
		## Description

		<!-- Provide a short summary of the changes. Why are they necessary? -->

		## Related Issue

		<!-- If this PR is related to an open issue, link it here. -->
		Fixes #[issue-number]

		## Type of Change

		<!-- Please delete options that are not relevant. -->
		- [ ] Bug fix (non-breaking change which fixes an issue)
		- [ ] New feature (non-breaking change which adds functionality)
		- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
		- [ ] Documentation update
		- [ ] Refactor (code improvements with no change in functionality)

		## How Has This Been Tested?

		<!-- Describe the tests that you ran to verify your changes. Provide instructions so reviewers can reproduce. -->
		- [ ] Unit Tests
		- [ ] Integration Tests
		- [ ] Manual Testing

		## Checklist:

		- [ ] My code follows the style guidelines of this project
		- [ ] I have performed a self-review of my code
		- [ ] I have commented my code, particularly in hard-to-understand areas
		- [ ] I have added necessary documentation (if applicable)
		- [ ] My changes generate no new warnings
		- [ ] I have added tests that prove my fix is effective or that my feature works
		- [ ] New and existing unit tests pass locally with my changes

		## Screenshots (if applicable)

		<!-- If your PR includes visual changes, include screenshots here. -->
	`

		gp := exec.Command("gh", "pr", "create", "--base", base, "--head", feat, "--title", title, "--body", body)
		gpOutput,gpErr := gp.Output()
		if gpErr != nil {
			fmt.Println("error on pr creation", gpErr)
			fmt.Println("error on pr creation", gpOutput)
			return
		}
		fmt.Println(string(gpOutput))
	},
}
