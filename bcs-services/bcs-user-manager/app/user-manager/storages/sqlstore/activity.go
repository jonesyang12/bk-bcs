/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package sqlstore xxx
package sqlstore

import (
	"fmt"
	"time"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-user-manager/app/user-manager/models"
)

// SearchActivities search activities
func SearchActivities(projectCode, resourceType, activityType string, status models.ActivityStatus,
	startTime, endTime time.Time, offset, limit int) ([]*models.Activity, int, error) {
	var activities []*models.Activity
	if projectCode == "" {
		return nil, 0, fmt.Errorf("projectCode can not be empty")
	}

	query := GCoreDB.Model(&models.Activity{}).Where("project_code = ?", projectCode)
	if resourceType != "" {
		query = query.Where("resource_type = ?", resourceType)
	}
	if activityType != "" {
		query = query.Where("activity_type = ?", activityType)
	}
	if status != 0 {
		query = query.Where("status = ?", status)
	}
	if startTime.Unix() != 0 {
		query = query.Where("created_at >= ?", startTime)
	}
	if endTime.Unix() != 0 {
		query = query.Where("created_at <= ?", endTime)
	}
	count := 0
	if err := query.Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := query.Offset(offset).Limit(limit).Order("created_at desc").Find(&activities).Error; err != nil {
		return nil, 0, err
	}
	return activities, count, nil
}

// CreateActivity create activity
func CreateActivity(activity []*models.Activity) error {
	for i := range activity {
		err := GCoreDB.Create(activity[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}
