package service

import (
	"github.com/HaHadaxigua/melancholy/internal/file/model"
	"github.com/HaHadaxigua/melancholy/internal/file/msg"
)

/**
这个文件夹中存放一些不是那么重要的方法。一些逻辑无关的方法
*/

// buildFileSearchItem 构建文件搜索的返回体
func buildFileSearchResult(folders []*model.Folder, files []*model.File) *msg.RspFileSearchResult {
	lenFiles, lenFolders := len(files), len(folders)
	var res msg.RspFileSearchResult
	list := make([]*msg.RspFileSearchItem, lenFiles+lenFolders)
	for i := 0; i < lenFiles; i++ {
		list[i] = files[i].ToFileSearchItem()
	}
	for j := lenFiles; j < len(list); j++ {
		list[j] = folders[j].ToFileSearchItem()
	}
	res.List = list
	return &res
}