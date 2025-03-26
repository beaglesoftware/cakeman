package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/google/go-github/v45/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// func open_github_pr() {}

func copyFile(src string, dest string) {
	data, err := os.ReadFile(src)
	if err != nil {
		printerror("failed to copy file '" + src + "': " + err.Error())
		os.Exit(1)
	}
	err = os.WriteFile(dest, data, 0755)
	if err != nil {
		printerror("failed to copy file '" + src + "': " + err.Error())
		os.Exit(1)
	}
}

func read_gh_token() string {
	token, err := loadToken()
	if err != nil {
		printerror("Failed to read token: " + err.Error())
	}
	return token.AccessToken
}

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish your cake to Cakeman confectionary",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Check for arguments. If the number arguments was less than 1, print an error and exit with error code '1'
		if len(args) < 1 {
			printerror("Not enough argument. Usage: cman publish [CAKE NAME]. Note that cake must be a library")
			os.Exit(1)
		}
		// Cake name
		arg1 := args[0]

		// Just for debugging
		// fmt.Println(arg1)

		// Read Cakefile, read its error in 'err' variable
		file, err := os.ReadFile(arg1 + ".cman")
		// print error if there was an error
		if err != nil {
			printerror("Error reading file: " + err.Error())
			hint("Currently, you can just publish a library, so make sure you are sharing a library cake, not binary")
			os.Exit(2)
		}
		// Read GitHub token
		token := read_gh_token()
		owner := "beaglesoftware"
		repo := "cakes"
		var config Config
		err = json.Unmarshal(file, &config)
		if err != nil {
			printerror("Failed to unmarshal JSON: " + err.Error())
		}
		cakeName := config.Package.Name

		// Format the new branch name as "{owner}.{repo}.{name}"
		// 		// fmt.Println(newBranchName)

		// Authenticate using the token
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		client := oauth2.NewClient(context.Background(), ts)
		ghClient := github.NewClient(client)

		user, _, err := ghClient.Users.Get(context.Background(), "")
		if err != nil {
			printerror("Failed to get user: " + err.Error())
		}
		username := user.Login
		// os.Exit(0)

		forkRepo, _, err := ghClient.Repositories.Get(context.Background(), *username, repo)
		if err != nil {
			printerror("Failed to get beaglesoftware/cakes repo: " + err.Error())
		}

		newBranchName := fmt.Sprintf("%s.%s.%s", *username, repo, cakeName)

		// fmt.Println(forkRepo)

		// os.Exit(0)

		if err != nil {
			if githubErr, ok := err.(*github.ErrorResponse); ok && githubErr.Response.StatusCode == 404 {
				info(fmt.Sprintf("Repository %s/%s not found\n", owner, forkRepo))
			} else {
				printerror("Error getting repository: " + err.Error())
			}
		} else {
			// fmt.Printf("Repository %s/%s exists\n", owner, "")
			// fmt.Println(repo)
		}

		// Fork the repository
		fork, _, err := ghClient.Repositories.CreateFork(context.Background(), owner, repo, nil)
		if err != nil {
			if _, ok := err.(*github.AcceptedError); ok {
				info("Forking was successful")
			} else {
				printerror("Error forking repository: " + err.Error())
			}
		}

		homedir, err := os.UserHomeDir()
		if err != nil {
			printerror("Failed to get home directory")
		}

		path := homedir + "/.cman/CakeIndex/"
		_, err = os.Stat(homedir + "/.cman/CakeIndex")
		if os.IsNotExist(err) {
			os.MkdirAll(homedir+"/.cman/CakeIndex", 0755)
		}

		_, err = git.PlainClone(path, false, &git.CloneOptions{
			URL:      fork.GetCloneURL(),
			Progress: os.Stdout,
		})
		if err != nil {
			printerror("Failed to clone repo: " + err.Error())
		}

		var firstchar string

		for _, char := range args[0] {
			firstchar = string(char)
			break
		}

		// Open the repository
		indexRepo, err := git.PlainOpen(path)
		if err != nil {
			printerror("Failed to open repository: " + err.Error())
		}

		// Get the working tree
		worktree, err := indexRepo.Worktree()
		if err != nil {
			printerror("Failed to get worktree: " + err.Error())
		}

		// Check if the branch already exists
		branchRef := plumbing.NewBranchReferenceName(newBranchName)
		_, err = indexRepo.Reference(branchRef, true)
		if err == nil {
			// Delete the branch if it exists
			err = indexRepo.Storer.RemoveReference(branchRef)
			if err != nil {
				printerror("Failed to delete existing branch: " + err.Error())
			}
			info("Deleted existing branch: " + newBranchName)
		}

		// Create a new branch
		headRef, err := indexRepo.Head()
		if err != nil {
			printerror("Failed to get HEAD reference: " + err.Error())
		}

		newBranchRef := plumbing.NewHashReference(branchRef, headRef.Hash())
		err = indexRepo.Storer.SetReference(newBranchRef)
		if err != nil {
			printerror("Failed to create new branch: " + err.Error())
		}
		info("Branch " + newBranchName + " created successfully")

		// Checkout the newBranchName branch
		err = worktree.Checkout(&git.CheckoutOptions{
			Branch: branchRef,
			// Create: true,
		})
		if err != nil {
			printerror("Failed to checkout " + newBranchName + " branch: " + err.Error())
			os.Exit(1)
		}
		info("Checked out to " + newBranchName + " branch")

		copyFile(arg1+".cman", path+"manifests/"+firstchar+"/"+arg1+".cman")

		// Create a new branch
		ref, _, err := ghClient.Git.GetRef(context.Background(), owner, repo, "refs/heads/main")
		if err != nil {
			printerror("Error getting reference: " + err.Error())
		}

		newRef := &github.Reference{
			Ref: github.String("refs/heads/" + newBranchName),
			Object: &github.GitObject{
				SHA: ref.Object.SHA,
			},
		}

		// Add the file to the staging area
		_, err = worktree.Add("manifests/" + firstchar + "/" + arg1 + ".cman")
		if err != nil {
			printerror("Failed to add file to staging area: " + err.Error())
			hint("Path: " + "manifests/" + firstchar + "/" + arg1 + ".cman")
			os.Exit(1)
		}

		// Commit the changes
		commitMsg := fmt.Sprintf("🍰 Add cake: %s", cakeName)
		_, err = worktree.Commit(commitMsg, &git.CommitOptions{
			Author: &object.Signature{
				Name:  user.GetName(),
				Email: user.GetEmail(),
				When:  time.Now(),
			},
		})
		if err != nil {
			printerror("Failed to commit changes: " + err.Error())
		}
		info("Commit was successful")

		// Push the changes to the remote repository
		err = indexRepo.Push(&git.PushOptions{
			RemoteName: "origin",
			Auth: &http.BasicAuth{
				Username: "x-access-token", // This can be anything except an empty string
				Password: token,            // GitHub token
			},
		})
		if err != nil {
			printerror("Failed to push changes: " + err.Error())
		}
		info("Push was successful")

		_, _, err = ghClient.Git.CreateRef(context.Background(), fork.GetOwner().GetLogin(), fork.GetName(), newRef)
		if err != nil {
			printerror("Error creating reference: " + err.Error())
		}
		info("Branch creation was successful")

		// Create a new pull request
		newPR := &github.NewPullRequest{
			Title: github.String("[Cake Request] Add cake: " + cakeName),
			Head:  github.String(fork.GetOwner().GetLogin() + ":" + newBranchName),
			Base:  github.String("main"),
			Body:  github.String("Created with `cman publish`\nThis PR adds " + cakeName + " cake."),
		}

		pr, _, err := ghClient.PullRequests.Create(context.Background(), owner, repo, newPR)
		if err != nil {
			printerror("Error creating pull request: " + err.Error())
			os.Exit(1)
		}
		info("Pull request creation was successful: " + pr.GetHTMLURL())

	},
}

func init() {
	rootCmd.AddCommand(publishCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// publishCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// publishCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
