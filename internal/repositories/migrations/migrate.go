package migrations

import "os"

// Migrate 数据库变更同步
func Migrate() {
	if len(os.Args) > 1 && os.Args[1][:5] == "-test" {
		return
	}

}
