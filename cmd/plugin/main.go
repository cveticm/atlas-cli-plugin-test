// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"

	"atlas-cli-plugin/internal/cli/echo"
	"atlas-cli-plugin/internal/cli/hello"
	"atlas-cli-plugin/internal/cli/listprofiles"
	"atlas-cli-plugin/internal/cli/printenv"
	"atlas-cli-plugin/internal/cli/stdinreader"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "example",
		Short: "Root command of the atlas cli plugin example",
	}

	rootCmd.AddCommand(
		hello.Builder(),
		echo.Builder(),
		printenv.Builder(),
		stdinreader.Builder(),
		listprofiles.Builder(),
	)

	execCmd := &cobra.Command{Use: "exec"}
	execCmd.AddCommand(rootCmd)

	if err := execCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
