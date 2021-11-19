package envinfo

import "regexp"

var versionRegex = regexp.MustCompile(`\d+\.[\d+|.]+`)
