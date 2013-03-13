// Copyright (C) 2013 Mark Stahl

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	_ "smbroker/handlers"
)

var usageText = `Usage:
    smbroker [command] [arguments]

The commands are:

    port       start broker on specified port
`

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		printBrokerUsage()		
	}

	switch args[0] {
	case "port":
		if len(args[1:]) > 0 {
			StartBroker(args[1])
		} else {
			StartBroker("10813")
		}
	}
}

func StartBroker(port string) {
	portStr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(portStr, nil)
}

func printBrokerUsage() {
	fmt.Println(usageText)
	os.Exit(0)
}
