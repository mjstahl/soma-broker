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

package data

import (
	"encoding/json"
	"fmt"
	"strings"
)

type JsonProject struct {
	Name    string
	Port    int
	RID     uint64
	Objects []Object
}

type Object struct {
	Name      string
	OID       uint64
	Behaviors map[string]uint64
}

func DecodeProjectFrom(addr string, body []byte) Project {
	var proj JsonProject
	err := json.Unmarshal(body, &proj)
	if err != nil {
		fmt.Println("project: decode error: ", err)
	}

	address := strings.Split(addr, ":")
	peer := Peer{Addr: address[0], Port: proj.Port, ID: proj.RID}
	project := Project{Name: proj.Name, Peers: []Peer{peer}, Objects: proj.Objects}
	return project
}

func EncodeProject(name string) []byte {
	proj := GetProject(name)

	bytes, err := json.Marshal(proj)
	if err != nil {
		fmt.Println("project: encode err: ", err)
	}

	return bytes
}
