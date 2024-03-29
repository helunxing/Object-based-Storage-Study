package temp

import (
	"lib/utils"
	"net/url"
	"os"
	"strconv"
	"strings"

	"../locate"
)

// 获取对象hash值
func (t *tempInfo) hash() string {
	s := strings.Split(t.Name, ".")
	return s[0]
}

// 获取对象分片号
func (t *tempInfo) id() int {
	s := strings.Split(t.Name, ".")
	id, _ := strconv.Atoi(s[1])
	return id
}

// 对象转正并加入缓存
func commitTempObject(datFile string, tempinfo *tempInfo) {
	f, _ := os.Open(datFile)
	d := url.PathEscape(utils.CalculateHash(f))
	f.Close()
	os.Rename(datFile, os.Getenv("STORAGE_ROOT")+"/objects/"+tempinfo.Name+"."+d)
	locate.Add(tempinfo.hash(), tempinfo.id())
}
