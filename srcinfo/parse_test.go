package srcinfo

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseSrcInfo(t *testing.T) {
	srcinfoData := `pkgbase = grpc
        pkgdesc = Package description
        pkgver = 0.11.1
        pkgrel = 3
        url = http://www.grpc.io/
        arch = i686
        arch = x86_64
        license = BSD
        makedepends = re2c
        makedepends = openssl
        makedepends = protobuf3
        makedepends = php
        source = https://github.com/grpc/grpc/archive/release-0_11_1.tar.gz
        md5sums = fb9b58c1f30deab63bd3ff2d046771a7

pkgname = grpc
        depends = openssl
        depends = protobuf3

pkgname = php-grpc
        depends = grpc=0.11.1-3
        depends = php
`
	srcinfo, err := ParseSrcInfo(strings.NewReader(srcinfoData))
	assert.NoError(t, err)
	assert.Equal(t, &SrcInfo{
		Global: Global{
			PkgBase: "grpc",
			PkgVer:  "0.11.1",
			PkgRel:  "3",
			CommonData: CommonData{
				PkgDesc: "Package description",
				URL:     "http://www.grpc.io/",
				Arch:    []string{"i686", "x86_64"},
				License: []string{"BSD"},
			},
			MakeDepends: []string{"re2c", "openssl", "protobuf3", "php"},
			Source: []string{
				"https://github.com/grpc/grpc/archive/release-0_11_1.tar.gz",
			},
			Checksums: Checksums{
				Md5Sums: []string{"fb9b58c1f30deab63bd3ff2d046771a7"},
			},
		},
		Packages: []Package{
			Package{
				PkgName: "grpc",
				CommonData: CommonData{
					Depends: []string{"openssl", "protobuf3"},
				},
			}, Package{
				PkgName: "php-grpc",
				CommonData: CommonData{
					Depends: []string{"grpc=0.11.1-3", "php"},
				},
			},
		},
	}, srcinfo)
}
