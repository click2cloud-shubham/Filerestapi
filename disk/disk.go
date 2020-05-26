package disk

import (
	"errors"
	"io/ioutil"
	"os"
	"path"
)

type Disk struct {
	// this is the folder you will use to all file operations.
	workspace string
}

// NewDisk creates a new instance of Disk (constructor)
//basically here we are checking if the directory exist or not
func NewDisk(workspace string) (Disk, error) {
	if workspace == "" {
		return Disk{}, errors.New("workspace is required")
	}

	// open directory workspace
	d, err := os.Open(workspace)
	if err != nil {
		return Disk{}, err
	}
	// close directory at the end of this function.
	defer d.Chdir()

	// get Stat of directory, when error occurs the file (directory) probably doesn't exist.
	s, err := d.Stat()
	if err != nil {
		return Disk{}, err
	}

	// check if the file is directory
	if !s.IsDir() {
		return Disk{}, errors.New("workspace is not a directory")
	}

	// create a new Disk structure with workspace set.
	return Disk{workspace: workspace}, nil
}

// CreateFile creates a new empty file.
func (d Disk) CreateFile(fileName string) error {
	_, err := os.Stat(path.Join(d.workspace, fileName))
	if err == nil {
		return errors.New("file exists")
	}

	f, err := os.Create(path.Join(d.workspace, fileName))
	if err != nil {
		return err
	}
	return f.Close()
}

// RenameFile changes the file name.
func (d Disk) RenameFile(oldFileName, newFileName string) error {
	// os.Rename renames oldFileName to newFileName, otherwise returns error
	return os.Rename(path.Join(d.workspace, oldFileName), path.Join(d.workspace, newFileName))
}

// DeleteFile removes the file.
func (d Disk) DeleteFile(fileName string) error {
	// os.Remove will remove file if exists, otherwise returns error.
	return os.Remove(path.Join(d.workspace, fileName))
}

// ListWorkspace gets all file names in the workspace.
func (d Disk) ListWorkspace() ([]string, error) {
	fi, err := ioutil.ReadDir(d.workspace)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, f := range fi {
		fileNames = append(fileNames, f.Name())
	}

	return fileNames, nil
}
