/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under,
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package common

import (
	"time"
)

// LocalTimeFormat is the format of local time
const LocalTimeFormat = "2006-01-02 15:04:05"

// Format3399ToLocalTime format time to local time
// "2006-01-02T15:04:05Z07:00" to "2006-01-02 15:04:05"
func Format3399ToLocalTime(t string) (string, error) {
	tt, err := time.Parse(time.RFC3339, t)
	if err != nil {
		return t, err
	}
	return tt.Format(LocalTimeFormat), nil
}
