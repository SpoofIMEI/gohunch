package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	ntemplates []string = []string{
		"FIRST.LAST",
		"FIRSTLAST",
		"LAST.FIRST",
		"LASTFIRST",
	}

	emailps []string = []string{
		"gmail.com",
		"hotmail.com",
		"hotmail.co.uk",
		"hotmail.fr",
		"yahoo.com",
		"yandex.com",
		"aol.com",
		"protonmail.com",
		"proton.me",
		"tutanota.com",
		"zoho.com",
	}

	names  []string
	pnames []string
	emails []string
)

func mutate_name(nfirst string, nlast string) {
	for _, ntemplate := range ntemplates {
		ptemplate := strings.Replace(ntemplate, "FIRST", nfirst, -1)
		ptemplate  = strings.Replace(ptemplate, "LAST", nlast, -1)

		pnames = append(pnames, ptemplate)
	}
}

func main() {
	name      := flag.String("name", "", "Name/File of names to generate emails from.")
	delimeter := flag.String("del", " ", "Custom delimeter for first and last name.")
	nomutate  := flag.Bool("nomutate", false, "Doesn't mutate first and last names into usernames.")
	uonly     := flag.Bool("uonly", false, "Only returns mutated usernames.")

	flag.Parse()

	if *name == "" {
		log.Fatal("Please enter a name...")
	} else if _, err := os.Stat(*name); err == nil {
		nfh, err := os.Open(*name)
		if err != nil {
			log.Fatal("Could not open file", "ERROR", err)
		}
		defer nfh.Close()

		nfs := bufio.NewScanner(nfh)
		for nfs.Scan() {
			names = append(names, nfs.Text())
		}
	} else {
		names = append(names, *name)
	}

	for _, name := range names {
		if len(strings.Split(name, *delimeter)) == 2 && !*nomutate {
			nfirst := strings.Split(name, *delimeter)[0]
			nlast  := strings.Split(name, *delimeter)[1]

			mutate_name(nfirst, nlast)
		} else {
			pnames = append(pnames, name)
		}

		for _, pname := range pnames {
			for _, emailp := range emailps {
				emails = append(emails, pname+"@"+emailp)
			}
		}

		if !*uonly {
			for _, email := range emails {
				fmt.Println(email)
			}
		} else {
			for _, username := range pnames {
				fmt.Println(username)
			}
		}
	}
}
