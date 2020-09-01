package main

// Paste the executable in a directory which is in PATH, so as to use it from the terminal "C:\Go\bin"
import (
	"os"
	"net/http"
	"flag"
	"log"
	"io"
	"path/filepath"
)


func main() {
	var name = flag.String("n", "", "Name of the file to be saved as.")
	var link = flag.String("l", "", "The link of the file to be downloaded.")
	flag.Parse()

	if *name == "" || *link == ""{
		log.Fatal("Proper flags not found.")
	}

	home, _ := os.UserHomeDir()
	err := os.Chdir(filepath.Join(home, "Downloads"))
	if err != nil {
		log.Fatal("Couldn't get to the Downloads directory.")
	}

	out, err := os.Create(*name)
	defer out.Close()
	if err != nil {
		log.Fatal("Can not create file.")
	}

	resp, err := http.Get(*link)
	if err != nil {
		log.Fatal("Couldn't get the file.")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Couldn't get an OK response.")
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal("Couldn't copy the response.")
	}
}