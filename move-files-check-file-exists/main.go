package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

/*
	example to move files from one location to another location
	replace file names with '_' with '-'
 */

func main(){

	// target location where we want to move files
	var movepath = "c:\\test1"

	// source location from where we to move files
	var sourcepath = "c:\\test"

	err := filepath.Walk(sourcepath, func(path string, info os.FileInfo, err error) error {
		if err != nil{
			return err
		}

		if((!info.IsDir()) && info.Size() > 0){
			if _, err := os.Stat(movepath + info.Name()); os.IsNotExist(err){
				fmt.Println(path, info.Size())
				filename := strings.Replace(info.Name(),"_", "-",-1)
				err := os.Rename(path, movepath + "\\" + filename)
				if err != nil{
					log.Fatal(err)
				}
			}
		}
		return nil
	})

	if err !=nil{
		log.Println(err)
	}
}
