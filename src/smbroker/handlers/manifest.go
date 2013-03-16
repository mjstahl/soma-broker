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
)

func HandleManifestRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// GET /m/<project name> (body will contain JSON text)
		// 200 OK - available, return json body for project
		// 404 Not Found  - if no peers are hosting this project
		//
	case "HEAD":
		// HEAD /m/<project name>
		// 404 Not Found - if the project doesn't exist
	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func init() {
	http.HandleFunc("/m/", HandleManifestRequest)
}
