/*
 * Copyright (c) 2023 ivfzhou
 * hlk_manager is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *          http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 */

package service

import (
	"context"
	"strconv"

	"gitee.com/CertificateAndSigningManageSystem/backend/protocol"
	"gitee.com/CertificateAndSigningManageSystem/common/util"
	"gitee.com/CertificateAndSigningManageSystem/hlk_manager/conf"
)

// 查询任务信息
func queryJobInfo(ctx context.Context, id int) (*protocol.HLK_QueryJobInfoRsp, error) {
	res, err := util.HTTPJsonGet[protocol.HLK_QueryJobInfoRsp](ctx,
		conf.Conf.BackendAddr+"/hlk/queryJobInfo?id="+strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	return res, nil
}
