/**
* Copyright 2018 Comcast Cable Communications Management, LLC
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
* http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package prometheus

import (
	"github.com/Comcast/trickster/internal/config"
	"github.com/Comcast/trickster/internal/routing"
	"github.com/Comcast/trickster/internal/util/log"
)

// RegisterRoutes registers the routes for the Client into the proxy's HTTP multiplexer
func (c *Client) RegisterRoutes(originName string, o *config.OriginConfig) {

	// Host Header-based routing
	log.Debug("Registering Origin Handlers", log.Pairs{"originType": o.Type, "originName": originName})
	routing.Router.HandleFunc("/"+mnHealth, c.HealthHandler).Methods("GET").Host(originName)
	routing.Router.HandleFunc(APIPath+mnQueryRange, c.QueryRangeHandler).Methods("GET", "POST").Host(originName)
	routing.Router.HandleFunc(APIPath+mnQuery, c.QueryHandler).Methods("GET", "POST").Host(originName)
	routing.Router.HandleFunc(APIPath+mnSeries, c.SeriesHandler).Methods("GET", "POST").Host(originName)
	routing.Router.HandleFunc(APIPath+mnLabels, c.ObjectProxyCacheHandler).Methods("GET", "POST").Host(originName)
	routing.Router.HandleFunc(APIPath+mnLabel, c.ObjectProxyCacheHandler).Methods("GET").Host(originName)
	routing.Router.HandleFunc(APIPath+mnTargets, c.ObjectProxyCacheHandler).Methods("GET").Host(originName)
	routing.Router.HandleFunc(APIPath+mnRules, c.ObjectProxyCacheHandler).Methods("GET").Host(originName)
	routing.Router.HandleFunc(APIPath+mnAlerts, c.ObjectProxyCacheHandler).Methods("GET").Host(originName)
	routing.Router.HandleFunc(APIPath+mnAlertManagers, c.ObjectProxyCacheHandler).Methods("GET").Host(originName)
	routing.Router.HandleFunc(APIPath+mnStatus, c.ObjectProxyCacheHandler).Methods("GET").Host(originName)
	routing.Router.PathPrefix(APIPath).HandlerFunc(c.ProxyHandler).Methods("GET", "POST").Host(originName)
	routing.Router.PathPrefix("/").HandlerFunc(c.ProxyHandler).Methods("GET", "POST").Host(originName)

	// Path based routing
	routing.Router.HandleFunc("/"+originName+"/"+mnHealth, c.HealthHandler).Methods("GET")
	routing.Router.HandleFunc("/"+originName+APIPath+mnQueryRange, c.QueryRangeHandler).Methods("GET", "POST")
	routing.Router.HandleFunc("/"+originName+APIPath+mnQuery, c.QueryHandler).Methods("GET", "POST")
	routing.Router.HandleFunc("/"+originName+APIPath+mnSeries, c.SeriesHandler).Methods("GET", "POST")
	routing.Router.HandleFunc("/"+originName+APIPath+mnLabels, c.ObjectProxyCacheHandler).Methods("GET", "POST")
	routing.Router.HandleFunc("/"+originName+APIPath+mnLabel, c.ObjectProxyCacheHandler).Methods("GET")
	routing.Router.HandleFunc("/"+originName+APIPath+mnTargets, c.ObjectProxyCacheHandler).Methods("GET")
	routing.Router.HandleFunc("/"+originName+APIPath+mnRules, c.ObjectProxyCacheHandler).Methods("GET")
	routing.Router.HandleFunc("/"+originName+APIPath+mnAlerts, c.ObjectProxyCacheHandler).Methods("GET")
	routing.Router.HandleFunc("/"+originName+APIPath+mnAlertManagers, c.ObjectProxyCacheHandler).Methods("GET")
	routing.Router.HandleFunc("/"+originName+APIPath+mnStatus, c.ObjectProxyCacheHandler).Methods("GET")
	routing.Router.PathPrefix("/"+originName+APIPath).HandlerFunc(c.ProxyHandler).Methods("GET", "POST")
	routing.Router.PathPrefix("/"+originName+"/").HandlerFunc(c.ProxyHandler).Methods("GET", "POST")

	// If default origin, set those routes too
	if o.IsDefault {
		log.Debug("Registering Default Origin Handlers", log.Pairs{"originType": o.Type})
		routing.Router.HandleFunc("/"+mnHealth, c.HealthHandler).Methods("GET")
		routing.Router.HandleFunc(APIPath+mnQueryRange, c.QueryRangeHandler).Methods("GET", "POST")
		routing.Router.HandleFunc(APIPath+mnQuery, c.QueryHandler).Methods("GET", "POST")
		routing.Router.HandleFunc(APIPath+mnSeries, c.SeriesHandler).Methods("GET", "POST")
		routing.Router.HandleFunc(APIPath+mnLabels, c.ObjectProxyCacheHandler).Methods("GET", "POST")
		routing.Router.HandleFunc(APIPath+mnLabel, c.ObjectProxyCacheHandler).Methods("GET")
		routing.Router.HandleFunc(APIPath+mnTargets, c.ObjectProxyCacheHandler).Methods("GET")
		routing.Router.HandleFunc(APIPath+mnRules, c.ObjectProxyCacheHandler).Methods("GET")
		routing.Router.HandleFunc(APIPath+mnAlerts, c.ObjectProxyCacheHandler).Methods("GET")
		routing.Router.HandleFunc(APIPath+mnAlertManagers, c.ObjectProxyCacheHandler).Methods("GET")
		routing.Router.HandleFunc(APIPath+mnStatus, c.ObjectProxyCacheHandler).Methods("GET")
		routing.Router.PathPrefix(APIPath).HandlerFunc(c.ProxyHandler).Methods("GET", "POST")
		routing.Router.PathPrefix("/").HandlerFunc(c.ProxyHandler).Methods("GET", "POST")
	}

}