package githubimages

import (
	"fmt"
	"context"
	"os"
	"net/http"
	"io/ioutil"

	// github "github.com/google/go-github/github"

	github "github.com/google/go-github/v50/github"
)

type Image struct {
	Name	string `json:"name"`
}

type Creds struct {
	Ghubkey string
	Ghuborg string 
	Ghubrepo string
	Ghubname string
	Ghubmail string
}

const folderName string = "./storage/images"

func UploadImages(image []byte) bool {
	fmt.Println("images uploaded for safe keeping")
	return false
}

func Push(list []string) bool {

	for _, img := range list {
        fmt.Println(img)
    }

// https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#create-a-file
// https://pkg.go.dev/github.com/google/go-github/github#example-RepositoriesService.CreateFile

	fmt.Println("pushing images to github")
	return false
}

// https://pkg.go.dev/github.com/google/go-github/github#example-RepositoriesService.CreateFile
// https://docs.github.com/en/rest/repos/contents?apiVersion=2022-11-28#create-a-file
func PushTest() bool {
	// In this example we're creating a new file in a repository using the
	// Contents API. Only 1 file per commit can be managed through that API.

	// Note that authentication is needed here as you are performing a modification
	// so you will need to modify the example to provide an oauth client to
	// github.NewClient() instead of nil. See the following documentation for more
	// information on how to authenticate with the client:
	// https://godoc.org/github.com/google/go-github/github#hdr-Authentication

	fmt.Println("pushing to github - start")

	creds := getGithubCreds()

	fmt.Println(creds)

	client := github.NewClient(nil)

	ctx := context.Background()
	fileContent := []byte("This is the content of my file\nand the 2nd line of it")

	// Note: the file needs to be absent from the repository as you are not
	// specifying a SHA reference here.
	opts := &github.RepositoryContentFileOptions{
		Message:   github.String("test"),
		Content:   fileContent,
		Branch:    github.String("master"),
		Committer: &github.CommitAuthor{Name: github.String(creds.Ghubname), Email: github.String(creds.Ghubmail)},
	}

	fmt.Println(opts.Branch)

	// for loop to get the file names then run through loop for every upload... looks like we will not be able to batch this, so will need to use queues instead. 
	// QUEUES, this is a test for now.
	_, _, err := client.Repositories.CreateFile(ctx, creds.Ghuborg, creds.Ghubrepo, "juicy.raw", opts)
	if err != nil {
		fmt.Println(err)
		return false	
	}

	return true
}

func CurlGetGit() bool {

	url := "https://api.github.com/octocat"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	return true
}

func CurlUploadToGit() bool {

	url := "https://api.github.com/octocat"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	return true
}

func getGithubCreds() Creds {
	return Creds{
		Ghubkey: 	os.Getenv("GHUBKEY"),
		Ghuborg: 	os.Getenv("GHUBORG"),
		Ghubrepo: 	os.Getenv("GHUBREPO"),
		Ghubname: 	os.Getenv("GHUBNAME"),
		Ghubmail: 	os.Getenv("GHUBMAIL"),
	} 
}