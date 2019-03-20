package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

/*
	example to recursively scan folders and list all the files which have size great than 0KB
 */

func main(){

	err := filepath.Walk("c:\\test", func(path string, info os.FileInfo, err error) error {
		if err != nil{
			return err
		}

		if((!info.IsDir()) && info.Size() > 0){
			fmt.Println(path, info.Size())
		}
		return nil
	})

	if err !=nil{
		log.Println(err)
	}
}
