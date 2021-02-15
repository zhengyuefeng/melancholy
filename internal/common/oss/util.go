/******
** @date : 2/15/2021 8:26 PM
** @author : zrx
** @description:
******/
package oss

import (
	"github.com/HaHadaxigua/melancholy/internal/global/consts"
	"regexp"
)

func VerifyBucketName(name string) bool {
	if ok, _ := regexp.MatchString(consts.OssBucketPattern, name); !ok {
		return false
	}
	return true
}
