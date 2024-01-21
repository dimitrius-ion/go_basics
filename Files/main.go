package main
import (
	"fmt" 
	"os"
	"github.com/dimitrius-ion/go_basics/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/go.mod"
	cont, err := fileutils.ReadTextFiles(filePath)
	if err != nil {
		fmt.Printf("ERROR PANIC: %v", err)
	}else {
		fmt.Println("Contents: \n", cont)
		banner := "____________________\n"
		newContent := fmt.Sprintf("Original: \n%s\n%v\nDouble Original:\n%s\n%v%v", banner, cont, banner, cont, cont)
		fileutils.WriteToFile(filePath + ".output.txt", newContent)
	}

}