package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Printer struct{}

func main() {
	fmt.Println("C'est parti pour le tri.")
	basePath := os.Args[1]
	f := make(map[string][]string)
	err := filepath.Walk(basePath+"\\atrier", func(path string, info os.FileInfo, err1 error) error {
		if err1 != nil {
			fmt.Println(err1)
			return err1
		}
		if !info.IsDir() {
			fmt.Println("-----------------------------------------")
			name := filepath.Base(path)
			fmt.Println(name)
			re := regexp.MustCompile(`\D+`)
			split := re.Split(name, -1)
			set := []string{}
			for i := range split {
				set = append(set, split[i])
			}
			for _, i := range set {
				r, err := time.Parse("20060102", i)
				if err == nil {
					month := strconv.Itoa(int(r.Month()))
					if len(month) == 1 {
						month = "0" + month
					}
					folder := strconv.Itoa(r.Year()) + "-" + month
					if f[folder] != nil {
						f[folder] = append(f[folder], path)
					} else {
						f[folder] = []string{path}

					}
				}
			}
		}

		return nil
	})
	for dossier := range f {
		dossierPath := basePath + "\\" + dossier
		if _, err := os.Stat(dossierPath); os.IsNotExist(err) {
			fmt.Println("Dossier:", dossier, " n'existe pas")
			os.Mkdir(dossierPath, os.ModeDevice)
		} else {
			fmt.Println("Dossier:", dossier, " existe")
		}
		for _, img := range f[dossier] {
			os.Rename(img, strings.Replace(img, "atrier", dossier, 1))
		}
	}
	if err != nil {
		fmt.Println("57:", err)
	}
}
