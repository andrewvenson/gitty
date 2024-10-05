package cmd

import (
	"errors"
	"os/exec"
	"os"
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	rootCmd.AddCommand(prCmd)
}

func getTitle(reader *bufio.Reader) (string, error) {
	fmt.Println("Enter pr title:")
	title,err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input for title")
		return "", errors.New("error")
	}
	return strings.TrimSuffix(title,"\n"), nil
}

func getBase(reader *bufio.Reader) (string, error) {
	fmt.Println("Enter base branch to pull pr into:")
	base,err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input for base")
		return "", errors.New("error")
	}
	return strings.TrimSuffix(base,"\n"), nil
}

func getFeat(reader *bufio.Reader) (string, error)  {
	featCmd := exec.Command("git", "branch","--show-current")
	output,err := featCmd.Output()
	if err != nil {
		fmt.Printf("Error executing command for showing current branch")
		return "", errors.New("error")
	}
	feat := string(output)
	return strings.TrimSuffix(feat, "\n"), nil
}

func createTempFile() (*os.File, error) {
	file, err := os.CreateTemp("", "pr_body_*.md")
	if err != nil {
		fmt.Printf("Error creating temporary file:", err)
		return nil, errors.New("error")
	}
	return file, nil
}

func writePrTemplateToTempFile(file *os.File) error {
		body := `
## Description

<!-- Provide a short summary of the changes. Why are they necessary? -->
...

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
	_, err := file.WriteString(body)
	if err != nil {
		fmt.Printf("Error writing to file:", err)
		return errors.New("error")
	}
	file.Close()
	return nil
}

func createPr(file *os.File, base string, feat string, title string) error {
	gp := exec.Command("gh", "pr", "create", "--base", base, "--head", feat, "--title", title, "--body-file", file.Name())
	gp.Stderr = os.Stderr
	gp.Stdout = os.Stdout
	gpErr := gp.Run()

	if gpErr != nil {
		fmt.Printf("Error on PR creation:", gpErr)
		os.Remove(file.Name())
		return errors.New("error")
	}
	os.Remove(file.Name())
	return nil
}

func changeDir() error {
	cwd,cwdErr := os.Getwd()
	if cwdErr != nil {
		fmt.Printf("error", cwdErr)
		return errors.New("error")
	}
	os.Chdir(cwd)
	return nil
}

var prCmd = &cobra.Command{
	Use:   "pr",
	Short: "Creates pull request",
	Long: `Create a pull request with custom template`,
	Run: func(cmd *cobra.Command, args []string){
		changeDirErr := changeDir()
		if changeDirErr != nil {
			return
		}
		
		reader := bufio.NewReader(os.Stdin)

		title,titleErr := getTitle(reader)
		if titleErr != nil {
			return
		}

		base,baseErr := getBase(reader)
		if baseErr != nil {
			return
		}

		feat,featErr := getFeat(reader)
		if featErr != nil {
			return
		}
		
		file,fileErr := createTempFile()
		if fileErr != nil {
			return
		}

		writeErr := writePrTemplateToTempFile(file)
		if writeErr != nil {
			return
		}
		
		createPrErr := createPr(file,base,feat,title)
		if createPrErr != nil {
			return
		}
	},
}
