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
	"log"
	"net/http"
	"path/filepath"
	"smbroker/data"
)

func HandleManifestRequest(w http.ResponseWriter, r *http.Request) {
	var status int

	projName := filepath.Base(r.URL.String())
	switch r.Method {
	case "HEAD":
		if data.BrokerHasProject(projName) {
			status = http.StatusOK
			w.WriteHeader(status)
		}
	case "GET":
		if data.BrokerHasProject(projName) {
			bytes := data.EncodeProject(projName)
			status = http.StatusOK
			w.Write(bytes)
		} else {
			status = http.StatusNotFound
			w.WriteHeader(status)
		}
	default:
		status = http.StatusMethodNotAllowed
	}

	log.Printf("%s %s => %d %s", r.Method, r.URL, status, http.StatusText(status))
}

func init() {
	http.HandleFunc("/m/", HandleManifestRequest)
}
