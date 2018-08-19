package psu

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"
)

func stringFromFile(fileName string) (string, error) {
	fh, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(fh)
	if err != nil {
		return "", err
	}
	return strings.TrimRight(bytes.NewBuffer(b).String(), " \n"), nil
}
