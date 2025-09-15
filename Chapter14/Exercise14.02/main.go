package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

var ErrorWorkingFileNotFount = errors.New("the working file is not found")

func CreateBackup(working, backup string) error {
	_, err := os.Stat(working)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrorWorkingFileNotFount
		}
		return err
	}
	workfile, err := os.Open(working)
	if err != nil {
		return err
	}
	defer workfile.Close()
	content, err := io.ReadAll(workfile)
	if err != nil {
		return err
	}
	err = os.WriteFile(backup, content, 0o644)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func AddNotes(workingFile, notes string) error {
	notes += "\n"
	f, err := os.OpenFile(workingFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write([]byte(notes)); err != nil {
		return err
	}
	return nil
}

func main() {
	backupFile := "backupFile.txt"
	workingFile := "note.txt"
	data := "note"
	err := CreateBackup(workingFile, backupFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 1; i <= 10; i++ {
		note := data + " " + strconv.Itoa(i)
		err := AddNotes(workingFile, note)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
