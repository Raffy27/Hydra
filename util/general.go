package util

import (
	"io"
	"os"
)

//RemoveDuplicates returns a new slice with duplicate elements removed.
func RemoveDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, i := range slice {
		if _, ok := keys[i]; !ok {
			keys[i] = true
			list = append(list, i)
		}
	}

	return list
}

//CopyFile attempts to copy a file from src to dst.
//Attributes are not preserved.
//Environment variables in paths are not supported.
func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}
