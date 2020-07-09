/*
Copyright 2020 The kconnect Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"

	"github.com/fidelity/kconnect/internal/version"
)

var versionCmd = &cli.Command{
	Name:  "version",
	Usage: "Display version & build information",
	Action: func(c *cli.Context) error {
		return doVersion(c)
	},
}

func init() {
	// TODO: add any additional flags
	App.Commands = append(App.Commands, versionCmd)
}

func doVersion(c *cli.Context) error {
	v := version.Get()
	outYaml, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Println(string(outYaml))

	return nil
}
