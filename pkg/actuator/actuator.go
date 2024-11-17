/*
 * Copyright 2021 Layotto Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package actuator

import (
	"mosn.io/layotto/kit/logger"

	"mosn.io/layotto/pkg/filter/stream/common/http"
)

type Actuator struct {
	endpointRegistry map[string]http.Endpoint
	Logger           logger.Logger
}

// New init an Actuator.
func New() *Actuator {
	a := &Actuator{
		endpointRegistry: make(map[string]http.Endpoint),
		Logger:           logger.NewLayottoLogger("actuator"),
	}
	return a
}

func (act *Actuator) OnLogLevelChanged(level logger.LogLevel) {
	act.Logger.SetLogLevel(level)
}

// GetEndpoint get an Endpoint from Actuator with name.
func (act *Actuator) GetEndpoint(name string) (endpoint http.Endpoint, ok bool) {
	e, ok := act.endpointRegistry[name]
	return e, ok
}

// AddEndpoint add an Endpoint to Actuator。
func (act *Actuator) AddEndpoint(name string, ep http.Endpoint) {
	if _, ok := act.endpointRegistry[name]; ok {
		act.Logger.Warnf("Duplicate Endpoint name: %v !", name)
	}
	act.endpointRegistry[name] = ep
}
