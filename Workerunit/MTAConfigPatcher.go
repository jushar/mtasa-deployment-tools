package main

import (
	"errors"
	"fmt"

	"github.com/beevik/etree"
)

type MTAConfigPatcher struct {
	Path        string
	XMLDocument *etree.Document
}

func NewMTAConfigPatcher(path string) (*MTAConfigPatcher, error) {
	patcher := new(MTAConfigPatcher)
	patcher.Path = path

	// Parse xml
	patcher.XMLDocument = etree.NewDocument()
	err := patcher.XMLDocument.ReadFromFile(path)
	if err != nil {
		return nil, err
	}

	return patcher, nil
}

func (patcher *MTAConfigPatcher) Patch(name string, value string) error {
	// Find element
	elem := patcher.XMLDocument.Root().SelectElement(name)
	if elem == nil {
		return errors.New("Could not find element")
	}

	// Patch it
	fmt.Printf("Patching %s to %s (old value: %s)\n", elem.Tag, value, elem.Text())
	elem.SetText(value)
	return nil
}

func (patcher *MTAConfigPatcher) Save() {
	// Write changes
	patcher.XMLDocument.WriteToFile(patcher.Path)
}
