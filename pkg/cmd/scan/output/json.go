package output

import (
	"encoding/json"
	"os"

	"github.com/cloudskiff/driftctl/pkg/analyser"
	"github.com/cloudskiff/driftctl/pkg/terraform"
)

const JSONOutputType = "json"
const JSONOutputExample = "json://PATH/TO/FILE.json"

type JSON struct {
	path string
}

func NewJSON(path string) *JSON {
	return &JSON{path}
}

func (c *JSON) Write(analysis *analyser.Analysis, _ *terraform.ProviderLibrary) error {
	file := os.Stdout
	if !isStdOut(c.path) {
		f, err := os.OpenFile(c.path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}
		defer f.Close()
		file = f
	}

	json, err := json.MarshalIndent(analysis, "", "\t")
	if err != nil {
		return err
	}
	if _, err := file.Write(json); err != nil {
		return err
	}
	return nil
}
