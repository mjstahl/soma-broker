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
	"io/ioutil"
	"log"
	"net/http"
	"smbroker/data"
)

func HandleProjectRequest(w http.ResponseWriter, r *http.Request) {
	status := 0
	switch r.Method {
	case "POST":
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()

		proj := data.DecodeProjectFrom(r.RemoteAddr, body)
		found := data.BrokerHasProject(proj)
		if found {
			status = http.StatusConflict
			w.WriteHeader(status)
		} else {
			data.StoreProject(proj)

			status = http.StatusCreated
			w.WriteHeader(status)
		}
	default:
		status = http.StatusMethodNotAllowed
		http.Error(w, "Method Not Allowed", 405)
	}

	log.Printf("%s %s => %d %s", r.Method, r.URL, status, http.StatusText(status))
}

func init() {
	http.HandleFunc("/p/", HandleProjectRequest)
}
