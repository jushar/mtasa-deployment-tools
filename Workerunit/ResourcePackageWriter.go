package main

import (
	"io"
	"mime/multipart"
	"os"
	"os/exec"
	"path/filepath"
)

type ResourcePackageWriter struct {
	InputStream *multipart.File
	ArchivePath string
}

func NewResourcePackageWriter(file *multipart.File, archivePath string) *ResourcePackageWriter {
	// Convert to absolute path (required for the tar command to work correctly)
	archivePath, _ = filepath.Abs(archivePath)

	return &ResourcePackageWriter{InputStream: file, ArchivePath: archivePath}
}

func (writer *ResourcePackageWriter) Write() error {
	outFile, err := os.OpenFile(writer.ArchivePath, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Write to outFile
	io.Copy(outFile, *writer.InputStream)
	return nil
}

func (writer *ResourcePackageWriter) Extract(targetDir string) error {
	// Call Unix tar (golang is able to untar/gunzip out-of-the-box, but this requires a lot more lines of code)
	cmd := exec.Command("tar", "xfz", writer.ArchivePath)

	// Set working dir to target dir
	cmd.Dir = targetDir

	// Go!
	return cmd.Run()
}
