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

package conf

import (
	"time"

	"github.com/go-ini/ini"
)

// Conf 配置
var Conf Data

// Data 配置
type Data struct {
	Mode        string `ini:"mode"`
	BackendAddr string `ini:"backendAddr"`
	Log         `ini:"log"`
}

// Log 日志配置
type Log struct {
	LogDir   string        `ini:"logDir"`
	Module   string        `ini:"module"`
	MaxAge   time.Duration `ini:"maxAge"`
	Rotation time.Duration `ini:"rotation"`
	Debug    bool          `ini:"debug"`
}

// InitialConf 初始化配置
func InitialConf(file string) {
	if len(file) <= 0 {
		file = "config.ini"
	}
	data, err := ini.Load(file)
	if err != nil {
		panic(err)
	}
	err = data.StrictMapTo(&Conf)
	if err != nil {
		panic(err)
	}
}
