package gitlab

import (
	"net/url"
	"path/filepath"
	"strings"
)

// URLParse todo
func URLParse(gitlabURL string) *Path {
	u, _ := url.Parse(gitlabURL)

	urlPath := u.Path
	parts := strings.SplitN(urlPath, "/", 6)
	partsLen := len(parts)
	// fmt.Println(parts, partsLen)
	sysPath := ""
	entrypoint := ""

	if partsLen >= 5 {
		group := parts[1]
		project := parts[2]
		sha := parts[4]
		sysPath = "/mnt/dav/" + u.Host + "/" + group + "+" + project + "/" + sha + "/"

		if partsLen == 6 && len(parts[5]) != 0 {
			srcPath := parts[5]
			filename := filepath.Base(srcPath)
			sysPath += strings.TrimSuffix(srcPath, filename)
			entrypoint += strings.TrimSuffix(filename, ".py") + ".main"
		}

		return &URLInfo{
			Domain:     u.Host,
			Group:      group,
			Project:    project,
			Sha:        sha,
			SysPath:    sysPath,
			EntryPoint: entrypoint,
		}
	}

	return &URLInfo{
		Domain:     "err_domain",
		Group:      "err_group",
		Project:    "err_project",
		Sha:        "err_sha",
		SysPath:    "err_syspath",
		EntryPoint: "err_entrypoint",
	}
}

// URLInfo todo
type URLInfo struct {
	Domain     string
	Group      string
	Project    string
	Sha        string
	SysPath    string
	EntryPoint string
}

// func pathParse(bs []byte) [][]byte {
// 	result := [][]byte{}
// 	pos := 0
// 	buffer := new(bytes.Buffer)
// 	for _, b := range bs {
// 		if (pos == 0 || pos == 3) && b == byte('/') {
// 			buffer.Reset()
// 		}
// 		if (pos == 1 || pos == 2 || pos == 4) && b == byte('/') {
// 			item := buffer.Bytes()
// 			buffer.Reset()
// 			result = append(result, item)

// 			pos++
// 		} else {
// 			buffer.WriteByte(b)
// 		}
// 	}
// }
