package api

import (
	vfs "github.com/blang/vfs"
)

func CreateFs() vfs.Filesystem { // not used yet kinda complicated
	//var osfs vfs.Filesystem = vfs.OS()
	var osfs vfs.Filesystem = vfs.OS()
	osfs.Mkdir("/tmp", 0777)
	return osfs
}
