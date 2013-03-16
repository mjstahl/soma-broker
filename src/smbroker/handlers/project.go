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
	"fmt"
	"io/ioutil"
	"net/http"
)

const projLen = len("/p/")

func HandleProjectRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// POST /p/<project name> (body will contain JSON text)
		// 409 Conflict - if the project name already exists (only 1 project by name right now)
		// 201 Created  - if the project name and body was accepted by the broker
		//

		fmt.Printf("Received POST request from %s\n", r.RemoteAddr)
		fmt.Printf("Request Body:\n")

		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		fmt.Printf("%s\n", body)

		w.WriteHeader(http.StatusCreated)
		//w.WriteHeader(http.StatusConflict)
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func init() {
	http.HandleFunc("/p/", HandleProjectRequest)
}
