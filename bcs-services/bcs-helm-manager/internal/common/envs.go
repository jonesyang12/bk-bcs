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
 *
 */

package common

import (
	"path/filepath"

	"github.com/Tencent/bk-bcs/bcs-services/bcs-helm-manager/internal/utils/envx"
	"github.com/Tencent/bk-bcs/bcs-services/bcs-helm-manager/internal/utils/path"
)

var (
	// LocalizeFilePath 国际化配置文件
	LocalizeFilePath = envx.GetEnv(
		"LOCALIZE_FILE_PATH", filepath.Dir(filepath.Dir(path.GetCurPKGPath()))+"/internal/i18n/locale/lc_msgs.yaml",
	)
)
