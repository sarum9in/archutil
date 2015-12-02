package srcinfo

type SrcInfo struct {
	Global   Global    `srcinfo:"section:pkgbase"`
	Packages []Package `srcinfo:"section:pkgname"`
}

type Global struct {
	CommonData
	PkgBase                    string
	PkgVer, PkgRel, Epoch      string
	MakeDepends                []string
	NoExtract, Options, Backup []string
	Source                     []string
	Checksums
}

type Package struct {
	CommonData
	PkgName         string
	Options, Backup []string
}

type CommonData struct {
	PkgDesc, URL, Install, Changelog  string
	Arch, Groups, License             []string
	CheckDepends, Depends, OptDepends []string
	Provides, Conflicts, Replaces     []string
}

type Checksums struct {
	Md5Sums    []string
	Sha1Sums   []string
	Sha224Sums []string
	Sha256Sums []string
	Sha384Sums []string
	Sha512Sums []string
}
