package cmd

import (
	"os/exec"
	"os"
	"bufio"
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
			fmt.Printf("error", cwdErr)
			return
		}
		os.Chdir(cwd)
		
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Enter pr title:")
		title,_ := reader.ReadString('\n')

		fmt.Println("Enter base branch to pull pr into:")
		base,_ := reader.ReadString('\n')

		featCmd := exec.Command("git", "branch","--show-current")
		output,err := featCmd.Output()
		if err != nil {
			fmt.Println("error",err)
		}
		feat := "remotes/origin/"+string(output)

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
			fmt.Printf("Error creating temporary file:", err)
			return
		}
		defer os.Remove(file.Name()) // Clean up the file after

		_, err = file.WriteString(body)
		if err != nil {
			fmt.Printf("Error writing to file:", err)
			return
		}
		file.Close()

		// Pass the body as a file to the gh command
		gp := exec.Command("gh", "pr", "create", "--base", base, "--head", feat, "--title", title, "--body-file", file.Name())
		gp.Stderr = os.Stderr
		gp.Stdout = os.Stdout
		gpErr := gp.Run()

		if gpErr != nil {
			fmt.Printf("Error on PR creation:", gpErr)
			return
		}
	},
}
