package gonfig

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Configuration interface is implemented by configuration types
// that are used in this package
type Configuration interface {
	File() string
}

// Read reads the JSON file and Marshals it in to the Configuration.
func Read(cfg Configuration) (err error) {
	var in []byte
	if in, err = ioutil.ReadFile(cfg.File()); err != nil {
		return
	}
	return json.Unmarshal(in, cfg)
}

// Write the configuration to file. When indent is true, the file's contents will
// be indented with four spaces, making it more human-friendly.
func Write(cfg Configuration, indent bool) (err error) {

	// Marshal to JSON...
	var body []byte
	if indent {
		body, err = json.MarshalIndent(cfg, "", "    ")
	} else {
		body, err = json.Marshal(cfg)
	}
	if err != nil {
		return
	}

	// Open the file...
	var f *os.File
	if f, err = os.OpenFile(cfg.File(), os.O_RDWR|os.O_CREATE, 0660); err != nil {
		return
	}
	defer f.Close()

	// Finally, write body to the file.
	_, err = f.Write(body)
	return
}
