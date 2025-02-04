/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package prometheus

import (
	"dubbo.apache.org/dubbo-go/v3/common/constant"
)

const (
	reporterName       = "prometheus"
	applicationNameKey = constant.ApplicationNameKey
	groupKey           = constant.GroupKey
	hostnameKey        = constant.HostnameKey
	interfaceKey       = constant.InterfaceKey
	ipKey              = constant.IpKey
	methodKey          = constant.MethodKey
	versionKey         = constant.VersionKey
)

const (
	providerField = "provider"
	consumerField = "consumer"

	qpsField      = "qps"
	requestsField = "requests"
	rtField       = "rt"

	milliSecondsField = "milliseconds"

	minField  = "min"
	maxField  = "max"
	sumField  = "sum"
	avgField  = "avg"
	lastField = "last"

	totalField      = "total"
	aggregateField  = "aggregate"
	processingField = "processing"
	succeedField    = "succeed"
)

var (
	quantiles = []float64{0.5, 0.9, 0.95, 0.99}
)
