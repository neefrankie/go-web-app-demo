package main

// Recursively get all file paths in directory, including sub-directories.
func GetAllFilePathsInDirectory(dirPath string) ([]string, error) {
	var paths []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}

// Recursively parse all files in directory, including sub-directories.
func ParseDirectory(dirPath string) (*template.Template, error) {
	paths, err := GetAllFilePathsInDirectory(dirPath)
	if err != nil {
		return nil, err
	}
	return template.ParseFiles(paths...)
}
