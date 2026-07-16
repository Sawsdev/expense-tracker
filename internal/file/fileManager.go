package file

import (
	"fmt"
	"log"
	"os"
)

var baseFileRoute string = "assets/db/"

func CreateFile(filename string) {
	file, err := os.Create(baseFileRoute + filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	fmt.Printf("File %s created at %s folder", filename, baseFileRoute)
}

func FileExists(filename string) bool{
	if _, err := os.Stat(baseFileRoute + filename); err == nil {
		return true
	}
	return false
}

func MakeDirectory(foldername string){
	err := os.Mkdir(foldername, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFile(filename string) ([]byte, error){
	file, err := os.ReadFile(baseFileRoute + filename)
	if err != nil {
		log.Fatal(err)
		return file, err
	}
	return file, err

}

func WriteFile(filename string, content []byte){
	err := os.WriteFile(baseFileRoute + filename, content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}