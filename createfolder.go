package createfolder

import (
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

func CreateFolder(path string, perm fs.FileMode) error {
	_, err := ioutil.ReadDir(path)

	if err == nil {
		return nil
	}

	if err.Error() != "open "+path+": no such file or directory" {
		return err
	}

	createErr := os.Mkdir(path, perm)

	if createErr == nil {
		return nil
	}

	paths := strings.Split(path, "/")
	createParentFolderErr := CreateFolder(strings.Join(paths[0:len(paths)-1], "/"), perm)

	if createParentFolderErr != nil {
		return createParentFolderErr
	}

	return os.Mkdir(path, perm)
}
