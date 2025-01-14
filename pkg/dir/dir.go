package dir

import "os"

// CreatePathIsNotExist create path is not exists
func CreatePathIsNotExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// create directory
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
	return false, err
}

// CheckPathExist path if exist
func CheckPathExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}
