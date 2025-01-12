// Copyright (c) 2023 Jingdong Technology Information Technology Co., Ltd.
// SBOM-TOOL is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package gem

import (
	"gitee.com/JD-opensource/sbom-tool/pkg/inventory/pckg/collector"
)

type Collector struct {
	collector.BaseCollector
}

var parsers []collector.FileParser

func init() {
	parsers = append(parsers, NewGemSpecParser())
	parsers = append(parsers, NewGemfileParser())
	parsers = append(parsers, NewGemFileLockParser())
}

func NewCollector() *Collector {
	c := Collector{}
	c.Name = Name()
	c.PurlType = PkgType()
	c.Parsers = parsers
	return &c
}
