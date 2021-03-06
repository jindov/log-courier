/*
* Copyright 2014-2015 Jason Woods.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package registrar

import (
	"github.com/driskell/log-courier/lc-lib/core"
)

type DeletedEvent struct {
	stream core.Stream
}

func NewDeletedEvent(stream core.Stream) *DeletedEvent {
	return &DeletedEvent{
		stream: stream,
	}
}

func (e *DeletedEvent) Process(state map[core.Stream]*FileState) {
	if _, ok := state[e.stream]; ok {
		log.Debug("Registrar received a deletion event for %s", *state[e.stream].Source)
	} else {
		log.Warning("Registrar received a deletion event for UNKNOWN (%p)", e.stream)
	}

	// Purge the registrar entry - means the file is deleted so we can't resume
	// This keeps the state clean so it doesn't build up after thousands of log files
	delete(state, e.stream)
}
