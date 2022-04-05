package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func urlHandler(url string) {
	// seperate domain and URIs
	splitUrl := strings.Split(url, "/")
	domain := splitUrl[2]
	// try root domain for open directory on http
	resp, err := http.Get("http://" + domain)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(string(body), "Index of") {
		fmt.Println("[x] Found potential open directory at http://" + domain)
	}

	// try root domain for open directory on https
	resp, err = http.Get("https://" + domain)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(string(body), "Index of") {
		fmt.Println("[x] Found potential open directory at https://" + domain)
	}
	//TODO - Add code to spider through a site using a wordlist?
	//TODO - Add code to check other directories for interesting files.
}

func main() {
	now := time.Now()
	wholeUrlPtr := flag.String("url", "https://example.com/214j9sa/example.php", "entire url")
	urlFilePtr := flag.String("file", "myfile.txt", "Text file")
  
	if !wholeUrlPtr.set && !urlFilePtr.set {
		fmt.Printf("[x] Error \n Please provide a URL (ex: dircheck.exe -url https://example.com), or a text file containing URLs (ex: dircheck.exe -file myfile.txt)")
	}
	if wholeUrlPtr.set && urlFilePtr.set {
		fmt.Printf("[x] Error \n Please only provide one flag")
	}
  // TODO: Find better way to check if flags are set 
  // TODO: Clean up file writer 
	println("Starting scan...")
	path := "dirscan_" + now.Format("2006-01-02 15:04:05")
	os.Mkdir(path, os.ModePerm)

	if wholeUrlPtr.set() {
		urlHandler(*wholeUrlPtr)
	}

	if urlFilePtr.set() {
		file, err := os.Open(*urlFilePtr)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			urlHandler(scanner.text)
		}
	}

}

//open dir finder,
