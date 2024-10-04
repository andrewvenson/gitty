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
		fmt.Scanf("%s", &commitMsg)

		gc := exec.Command("git", "commit", "-m", "'"+commitMsg+"'")
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
		fmt.Scanf("%s",&title)

		fmt.Println("Enter base branch to pull pr into:")
		fmt.Scanf("%s",&base)

		fmt.Println("Enter feat branch name to pull pr into:")
		fmt.Scanf("%s",&feat)

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

		// Create a temporary file to store the PR body
		file, err := os.CreateTemp("", "pr_body_*.md")
		if err != nil {
			fmt.Println("Error creating temporary file:", err)
			return
		}
		defer os.Remove(file.Name()) // Clean up the file after

		_, err = file.WriteString(body)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		file.Close()

		// Pass the body as a file to the gh command
		gp := exec.Command("gh", "pr", "create", "--base", base, "--head", feat, "--title", title, "--body-file", file.Name())
		gp.Stderr = os.Stderr
		gp.Stdout = os.Stdout
		gpErr := gp.Run()

		if gpErr != nil {
			fmt.Println("Error on PR creation:", gpErr)
			return
		}
	},
}
