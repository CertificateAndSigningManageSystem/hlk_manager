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

package main

import (
	"flag"
	"fmt"
	"os"

	"gitee.com/CertificateAndSigningManageSystem/common/ctxs"
	"gitee.com/CertificateAndSigningManageSystem/common/log"

	"gitee.com/CertificateAndSigningManageSystem/hlk_manager/conf"
	"gitee.com/CertificateAndSigningManageSystem/hlk_manager/consts"
	"gitee.com/CertificateAndSigningManageSystem/hlk_manager/service"
)

var mode string

func init() {
	conf.InitialConf("app.ini")
	ctx := ctxs.NewCtx("init")
	flag.StringVar(&mode, "mode", "", fmt.Sprintf("运行模式；测试机%s，控制器%s，宿主机%s",
		consts.TestSystemMode, consts.TestServerMode, consts.HostMachineMode))
	flag.Parse()
	if len(mode) <= 0 {
		// 读取配置文件
		mode = conf.Conf.Mode
		if len(mode) <= 0 {
			// 读取环境变量
			mode = os.Getenv(consts.ModeEnvKey)
			if len(mode) <= 0 {
				log.Fatal(ctx, "no mode available")
			}
		}
	}
}

func main() {
	ctx := ctxs.NewCtx("main")
	// 根据不同 mode 运行不同逻辑
	switch mode {
	case consts.TestSystemMode:
		service.StartTestSystemHandler(ctx)
	case consts.TestServerMode:
	case consts.HostMachineMode:
	default:
		log.Fatal(ctx, "未知的模式", mode)
	}
}
