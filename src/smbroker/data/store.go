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
	"sync"
)

type Project struct {
	Name string
	Peers []Peer
	Objects []Object
}

type Peer struct {
	Addr string
	Port int
	ID uint64
}

var Store *store

type store struct {
	sync.Mutex

	Projects map[string]Project
}

func BrokerHasProject(proj Project) bool {
	Store.Lock()
	defer Store.Unlock()

	if _, ok := Store.Projects[proj.Name]; ok {
		return true
	}

	return false
}

func StoreProject(proj Project) {
	Store.Lock()
	defer Store.Unlock()

	Store.Projects[proj.Name] = proj
}

func init() {
	Store = &store{Projects: map[string]Project{}}
}