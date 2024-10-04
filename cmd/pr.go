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
		cwd,err := os.Getwd()
		if err != nil {
			fmt.Println("error", err)
		}

		fmt.Println(cwd)
		os.Chdir(cwd)
		
		gs := exec.Command("git", "status")
		gsOutput,_ := gs.Output()
		fmt.Println(string(gsOutput))

		ga := exec.Command("git", "add", "-A")
		gaOutput,_ := ga.Output()
		fmt.Println(string(gaOutput))
		
		var title string 
		fmt.Scanln(&title)

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

	gp := exec.Command("gh", "pr", "create", "--base", baseBranch, "--head", featBranch, "--title", title, "--body", body)
		gaOutput,_ := ga.Output()
		fmt.Println(string(gaOutput))
	},
}
