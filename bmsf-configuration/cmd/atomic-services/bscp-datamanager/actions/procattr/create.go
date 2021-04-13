/*
Tencent is pleased to support the open source community by making Blueking Container Service available.
Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except
in compliance with the License. You may obtain a copy of the License at
http://opensource.org/licenses/MIT
Unless required by applicable law or agreed to in writing, software distributed under
the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
either express or implied. See the License for the specific language governing permissions and
limitations under the License.
*/

package procattr

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/spf13/viper"

	"bk-bscp/internal/database"
	"bk-bscp/internal/dbsharding"
	pbcommon "bk-bscp/internal/protocol/common"
	pb "bk-bscp/internal/protocol/datamanager"
	"bk-bscp/internal/strategy"
	"bk-bscp/pkg/common"
)

// CreateAction creates a procattr object.
type CreateAction struct {
	ctx   context.Context
	viper *viper.Viper
	smgr  *dbsharding.ShardingManager

	req  *pb.CreateProcAttrReq
	resp *pb.CreateProcAttrResp

	sd *dbsharding.ShardingDB
}

// NewCreateAction creates new CreateAction.
func NewCreateAction(ctx context.Context, viper *viper.Viper, smgr *dbsharding.ShardingManager,
	req *pb.CreateProcAttrReq, resp *pb.CreateProcAttrResp) *CreateAction {
	action := &CreateAction{ctx: ctx, viper: viper, smgr: smgr, req: req, resp: resp}

	action.resp.Seq = req.Seq
	action.resp.Code = pbcommon.ErrCode_E_OK
	action.resp.Message = "OK"

	return action
}

// Err setup error code message in response and return the error.
func (act *CreateAction) Err(errCode pbcommon.ErrCode, errMsg string) error {
	act.resp.Code = errCode
	act.resp.Message = errMsg
	return errors.New(errMsg)
}

// Input handles the input messages.
func (act *CreateAction) Input() error {
	if err := act.verify(); err != nil {
		return act.Err(pbcommon.ErrCode_E_DM_PARAMS_INVALID, err.Error())
	}
	return nil
}

// Output handles the output messages.
func (act *CreateAction) Output() error {
	// do nothing.
	return nil
}

func (act *CreateAction) verify() error {
	var err error

	if err = common.ValidateString("cloud_id", act.req.CloudId,
		database.BSCPNOTEMPTY, database.BSCPNAMELENLIMIT); err != nil {
		return err
	}
	if err = common.ValidateString("ip", act.req.Ip,
		database.BSCPNOTEMPTY, database.BSCPNORMALSTRLENLIMIT); err != nil {
		return err
	}
	if err = common.ValidateString("biz_id", act.req.BizId,
		database.BSCPNOTEMPTY, database.BSCPIDLENLIMIT); err != nil {
		return err
	}
	if err = common.ValidateString("app_id", act.req.AppId,
		database.BSCPNOTEMPTY, database.BSCPIDLENLIMIT); err != nil {
		return err
	}
	if err = common.ValidateString("labels", act.req.Labels,
		database.BSCPEMPTY, database.BSCPLABELSSIZELIMIT); err != nil {
		return err
	}
	if len(act.req.Labels) == 0 {
		act.req.Labels = strategy.EmptySidecarLabels
	}

	act.req.Path = filepath.Clean(act.req.Path)
	if err = common.ValidateString("path", act.req.Path,
		database.BSCPNOTEMPTY, database.BSCPCFGFPATHLENLIMIT); err != nil {
		return err
	}

	if err = common.ValidateString("creator", act.req.Creator,
		database.BSCPNOTEMPTY, database.BSCPNAMELENLIMIT); err != nil {
		return err
	}
	if err = common.ValidateString("memo", act.req.Memo,
		database.BSCPEMPTY, database.BSCPLONGSTRLENLIMIT); err != nil {
		return err
	}
	return nil
}

func (act *CreateAction) createProcAttrOverride() (pbcommon.ErrCode, string) {
	st := database.ProcAttr{
		CloudID:      act.req.CloudId,
		IP:           act.req.Ip,
		BizID:        act.req.BizId,
		AppID:        act.req.AppId,
		Path:         act.req.Path,
		Labels:       act.req.Labels,
		Creator:      act.req.Creator,
		LastModifyBy: act.req.Creator,
		Memo:         act.req.Memo,
	}

	err := act.sd.DB().
		Where(database.ProcAttr{
			CloudID: act.req.CloudId,
			IP:      act.req.Ip,
			BizID:   act.req.BizId,
			AppID:   act.req.AppId,
			Path:    act.req.Path,
		}).
		Assign(st).
		FirstOrCreate(&st).Error

	if err != nil {
		return pbcommon.ErrCode_E_DM_DB_EXEC_ERR, err.Error()
	}
	return pbcommon.ErrCode_E_OK, ""
}

func (act *CreateAction) createNewProcAttr() (pbcommon.ErrCode, string) {
	st := database.ProcAttr{
		CloudID:      act.req.CloudId,
		IP:           act.req.Ip,
		BizID:        act.req.BizId,
		AppID:        act.req.AppId,
		Path:         act.req.Path,
		Labels:       act.req.Labels,
		Creator:      act.req.Creator,
		LastModifyBy: act.req.Creator,
		Memo:         act.req.Memo,
	}

	err := act.sd.DB().
		Where(database.ProcAttr{
			CloudID: act.req.CloudId,
			IP:      act.req.Ip,
			BizID:   act.req.BizId,
			AppID:   act.req.AppId,
			Path:    act.req.Path,
		}).
		Attrs(st).
		FirstOrCreate(&st).Error

	if err != nil {
		return pbcommon.ErrCode_E_DM_DB_EXEC_ERR, err.Error()
	}

	if st.Labels != act.req.Labels || st.Creator != act.req.Creator ||
		st.LastModifyBy != act.req.Creator || st.Memo != act.req.Memo {
		return pbcommon.ErrCode_E_DM_ALREADY_EXISTS,
			"the procattr with target cloudid-ip-bizid-appid-path already exist."
	}
	return pbcommon.ErrCode_E_OK, ""
}

func (act *CreateAction) createProcAttr() (pbcommon.ErrCode, string) {
	if act.req.Override {
		return act.createProcAttrOverride()
	}
	return act.createNewProcAttr()
}

// Do makes the workflows of this action base on input messages.
func (act *CreateAction) Do() error {
	// BSCP sharding db.
	sd, err := act.smgr.ShardingDB(dbsharding.BSCPDBKEY)
	if err != nil {
		return act.Err(pbcommon.ErrCode_E_DM_ERR_DBSHARDING, err.Error())
	}
	act.sd = sd

	// create procattr.
	if errCode, errMsg := act.createProcAttr(); errCode != pbcommon.ErrCode_E_OK {
		return act.Err(errCode, errMsg)
	}
	return nil
}
