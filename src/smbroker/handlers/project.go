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

package handlers

import (
	"net/http"
	"fmt"
)

const projLen = len("/p/")

func HandleProjectRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// POST /p/<project name> (body will contain JSON text)
		// 409 Conflict - if the project name already exists (only 1 project by name right now)
		// 201 Created  - if the project name and body was accepted by the broker
		//
		fmt.Fprintf(w, "Received POST request. This should return a 409 on failure, or a 201 on success.")
	default:
        http.Error(w, "Method Not Allowed", 405)
	}
}

func init() {
	http.HandleFunc("/p/", HandleProjectRequest)
}
