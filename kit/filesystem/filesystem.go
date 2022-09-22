package filesystem

import (
	"kwd/kernel/app"
	"kwd/kit/filesystem/local"
	"kwd/kit/filesystem/qiniu"
	"kwd/kit/interface"
)

func New() _interface.FilesystemInterface {
	return Disk(app.Cfg.File.Driver)
}

func Disk(disk string) _interface.FilesystemInterface {

	var storage _interface.FilesystemInterface

	switch disk {
	case "qiniu":
		storage = qiniu.New()
	default:
		storage = local.New()
	}

	return storage
}
