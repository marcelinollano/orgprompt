package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
)

type Sfdx struct {
	TargetDevHub string `json:"defaultdevhubusername"`
	TargetOrg    string `json:"defaultusername"`
}

func prompt(message string) {
	prompt := log.New(os.Stdout, "", 0)
	prompt.Print("(" + message + ") ")
}

func arrange(payload Sfdx, devhub *bool) {
	var divider string = "|"
	var unknown string = "?"

	if *devhub {
		if payload.TargetDevHub != "" && payload.TargetOrg != "" {
			prompt(payload.TargetDevHub + divider + payload.TargetOrg)
		}

		if payload.TargetDevHub == "" && payload.TargetOrg != "" {
			prompt(unknown + divider + payload.TargetOrg)
		}

		if payload.TargetDevHub != "" && payload.TargetOrg == "" {
			prompt(payload.TargetDevHub + divider + unknown)
		}

		if payload.TargetDevHub == "" && payload.TargetOrg == "" {
			prompt(unknown + divider + unknown)
		}
	} else {
		if payload.TargetOrg != "" {
			prompt(payload.TargetOrg)
		}

		if payload.TargetOrg == "" {
			prompt(unknown)
		}
	}
}

func main() {
	var sfdx_path string = "./.sfdx/sfdx-config.json"

	devhub := flag.Bool("devhub", false, "displays both Default Dev Hub and Default Org")
	flag.Parse()

	_, err := os.Stat(sfdx_path)
	var exists bool
	if !os.IsNotExist(err) {
		exists = true
	} else {
		exists = false
	}

	if exists {
		content, err := ioutil.ReadFile(sfdx_path)
		if err != nil {
			log.Fatal("Error when opening file: ", err)
		}

		var payload Sfdx
		err = json.Unmarshal(content, &payload)
		if err != nil {
			log.Fatal("Error during Unmarshal(): ", err)
		}

		arrange(payload, devhub)
	}
}
