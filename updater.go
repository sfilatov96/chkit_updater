package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"os/exec"
	"github.com/fatih/color"
)




func main() {
	if len(os.Args) <= 1 {
		color.Red("Command Line Arguments unexpected")
		defer color.Unset()
		os.Exit(0)
	}
	URL := os.Args[1]
	color.Blue("Downloading file...")
	rawURL := URL

	fileURL, err := url.Parse(rawURL)

	if err != nil {
		panic(err)
	}

	check := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, err := check.Get(rawURL) // add a filter to check redirect

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		color.Red(resp.Status)
		defer color.Unset()
		os.Exit(0)
	} else {
		color.Green(resp.Status)
	}

	path := fileURL.Path

	segments := strings.Split(path, "/")

	fileName := segments[2] + ".tar.gz"// change the number to accommodate changes to the url.Path position



	if err != nil {
		panic(err)
	}



	remove_all_cmd := "ls | grep -v updater | xargs rm -rf"
	exec.Command("bash", "-c", remove_all_cmd).Output()

	file, err := os.Create(fileName)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer file.Close()

	size, err := io.Copy(file, resp.Body)

	color.Blue("%s with %v bytes downloaded", fileName, size)
	exec.Command("tar", "-xvf",fileName).Output()
	remove_archive_cmd := "rm *.tar.gz"
	exec.Command("bash", "-c", remove_archive_cmd).Output()
	color.Green("ChKit updated success!")
	defer color.Unset()




}
