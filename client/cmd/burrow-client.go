// Copyright 2017 Monax Industries Limited
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"os"
	"strconv"

	"fmt"

	"github.com/hyperledger/burrow/client"
	"github.com/hyperledger/burrow/project"
	"github.com/spf13/cobra"
)

// Global flags for persistent flags
var clientDo *client.Do

var BurrowClientCmd = &cobra.Command{
	Use:   "burrow-client",
	Short: "burrow-client interacts with a running burrow chain.",
	Long: `burrow-client interacts with a running burrow chain.

Made with <3 by Monax Industries.

Complete documentation is available at https://monax.io/docs

VERSION: ` + project.FullVersion(),
	Run: func(cmd *cobra.Command, args []string) { cmd.Help() },
}

func Execute() {
	InitBurrowClientInit()
	AddGlobalFlags()
	AddClientCommands()
	BurrowClientCmd.Execute()
}

func InitBurrowClientInit() {
	// initialise an empty Do struct for command execution
	clientDo = client.NewClientDo()
}

func AddGlobalFlags() {
	BurrowClientCmd.PersistentFlags().BoolVarP(&clientDo.Verbose, "verbose", "v", defaultVerbose(), "verbose output; more output than no output flags; less output than debug level; default respects $BURROW_CLIENT_VERBOSE")
	BurrowClientCmd.PersistentFlags().BoolVarP(&clientDo.Debug, "debug", "d", defaultDebug(), "debug level output; the most output available for burrow-client; if it is too chatty use verbose flag; default respects $BURROW_CLIENT_DEBUG")
}

func AddClientCommands() {
	BurrowClientCmd.AddCommand(buildTransactionCommand(), buildStatusCommand())
	BurrowClientCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print full version",
		Run:   func(cmd *cobra.Command, args []string) { fmt.Println(project.FullVersion()) },
	})
}

//------------------------------------------------------------------------------
// Defaults

// defaultVerbose is set to false unless the BURROW_CLIENT_VERBOSE environment
// variable is set to a parsable boolean.
func defaultVerbose() bool {
	return setDefaultBool("BURROW_CLIENT_VERBOSE", false)
}

// defaultDebug is set to false unless the BURROW_CLIENT_DEBUG environment
// variable is set to a parsable boolean.
func defaultDebug() bool {
	return setDefaultBool("BURROW_CLIENT_DEBUG", false)
}

// setDefaultBool returns the provided default value if the environment variable
// is not set or not parsable as a bool.
func setDefaultBool(environmentVariable string, defaultValue bool) bool {
	value := os.Getenv(environmentVariable)
	if value != "" {
		if parsedValue, err := strconv.ParseBool(value); err == nil {
			return parsedValue
		}
	}
	return defaultValue
}

func setDefaultString(envVar, def string) string {
	env := os.Getenv(envVar)
	if env != "" {
		return env
	}
	return def
}
