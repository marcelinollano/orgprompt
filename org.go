package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Sfdx struct {
	TargetDevHub string `json:"defaultdevhubusername"`
	TargetOrg    string `json:"defaultusername"`
}

func main() {
	var sfdx_path = "./.sfdx/sfdx-config.json"

	if _, err := os.Stat(sfdx_path); err == nil {
		output := log.New(os.Stdout, "", 0)

		content, err := ioutil.ReadFile(sfdx_path)
		if err != nil {
			output.Fatal("Error when opening file: ", err)
		}

		var payload Sfdx
		err = json.Unmarshal(content, &payload)
		if err != nil {
			output.Fatal("Error during Unmarshal(): ", err)
		}

		if payload.TargetDevHub != "" && payload.TargetOrg != "" {
			var org = payload.TargetDevHub + " → " + payload.TargetOrg
			output.Print(org)
		}

		if payload.TargetDevHub == "" && payload.TargetOrg != "" {
			var org = "? → " + payload.TargetOrg
			output.Print(org)
		}

		if payload.TargetDevHub != "" && payload.TargetOrg == "" {
			var org = payload.TargetDevHub + " → ?"
			output.Print(org)
		}

		if payload.TargetDevHub == "" && payload.TargetOrg == "" {
			var org = "? → ?"
			output.Print(org)
		}
	}
}
