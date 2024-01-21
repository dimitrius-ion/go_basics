package fileutils
import "os"


func ReadTextFiles (filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}else {
		return string(content), nil
	}

	return "This is the content of the file", nil
}