// Copyright (c) 2023 Jingdong Technology Information Technology Co., Ltd.
// SBOM-TOOL is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package composer

import (
	"testing"

	"gitee.com/JD-opensource/sbom-tool/pkg/model"
	"gitee.com/JD-opensource/sbom-tool/pkg/util"
)

type testComposerLockItem struct {
	title    string
	filePath string
	expected []model.Package
}

var composerLockTestdata = []testComposerLockItem{
	{
		title:    "testComposerJson test",
		filePath: "test_material/composer.lock",
		expected: []model.Package{
			{Name: "adoy/fastcgi-client", Version: "1.0.2", Type: model.PkgTypeComposer, LicenseDeclared: []string{"MIT"}},
			{Name: "alcaeus/mongo-php-adapter", Version: "1.1.11", Type: model.PkgTypeComposer, LicenseDeclared: []string{"MIT"}},
		},
	},
}

func TestParseComposerLockFile(t *testing.T) {
	for _, item := range composerLockTestdata {
		parse := NewComposerLockFileParser()
		pkgs, err := parse.Parse(item.filePath)
		if err != nil {
			t.Errorf("test error[%v]: %e", item.title, err)
		}

		if !util.SliceEqual(pkgs, item.expected, func(p1 model.Package, p2 model.Package) bool {
			return model.PackageEqual(&p1, &p2)
		}) {
			t.Errorf("test failed[%v]: expected = %v got %v", item.title, item.expected, pkgs)
		}
	}
}

func BenchmarkComposerLockParser(b *testing.B) {
	parse := NewComposerLockFileParser()
	for i := 0; i < b.N; i++ {
		_, _ = parse.Parse("test_material/composer.lock")
	}
}
