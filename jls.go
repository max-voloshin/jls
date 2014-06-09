package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"syscall"
)

type Item struct {
	Name         string `json:"name"`
	Path         string `json:"path"`
	Size         int64  `json:"size"`
	Dir          bool   `json:"dir"`
	File         bool   `json:"file"`
	AppendOnly   bool   `json:"append_only"`
	Exclusive    bool   `json:"exclusive"`
	Temporary    bool   `json:"temporary"`
	Symlink      bool   `json:"symlink"`
	Device       bool   `json:"device"`
	NamedPipe    bool   `json:"named_pipe"`
	Socket       bool   `json:"socket"`
	Setuid       bool   `json:"setuid"`
	Setgid       bool   `json:"setgid"`
	CharDevice   bool   `json:"char_device"`
	Sticky       bool   `json:"sticky"`
	OwnerRead    bool   `json:"owner_read"`
	OwnerWrite   bool   `json:"owner_write"`
	OwnerExecute bool   `json:"owner_execure"`
	GroupRead    bool   `json:"group_read"`
	GroupWrite   bool   `json:"group_write"`
	GroupExecute bool   `json:"group_execute"`
	OtherRead    bool   `json:"other_read"`
	OtherWrite   bool   `json:"other_write"`
	OtherExecute bool   `json:"other_execute"`
	Permissions  string `json:"permissions"`
	Mtime        int64  `json:"mtime"`
	Ctime        int64  `json:"ctime"`
	Atime        int64  `json:"atime"`
	Uid          uint32 `json:"uid"`
	Gid          uint32 `json:"gid"`
}

func main() {
	defaultUsage := flag.Usage
	flag.Usage = func() {
		defaultUsage()
		fmt.Fprintf(os.Stderr, "\n\t%s [file ...]\n\nJSON edition of 'ls' util\n", os.Args[0])
	}
	flag.Parse()

	var names []string
	if 0 == flag.NArg() {
		workingDirectory, error := os.Getwd()
		if error != nil {
			log.Fatal(error)
		}
		names = append(names, workingDirectory)
	} else {
		names = flag.Args()
	}

	for _, name := range names {
		path, error := filepath.Abs(name)
		if error != nil {
			log.Fatal(error)
		}
		info, error := os.Lstat(name)
		if error != nil {
			log.Fatal(error)
		}
		if info.IsDir() {
			dir, error := os.Open(name)
			if error != nil {
				log.Fatal(error)
			}
			list, error := dir.Readdir(0)
			if error != nil {
				log.Fatal(error)
			}
			for _, file := range list {
				handleFile(path+string(os.PathSeparator)+file.Name(), file)
			}
		} else {
			handleFile(path, info)
		}
	}
}

func handleFile(path string, file os.FileInfo) {
	item := new(Item)

	item.Name = file.Name()

	item.Path = path

	item.Size = file.Size()

	item.File = file.Mode().IsRegular()
	item.Dir = file.Mode().IsDir()
	item.AppendOnly = file.Mode()&os.ModeAppend == os.ModeAppend
	item.Exclusive = file.Mode()&os.ModeExclusive == os.ModeExclusive
	item.Temporary = file.Mode()&os.ModeTemporary == os.ModeTemporary
	item.Symlink = file.Mode()&os.ModeSymlink == os.ModeSymlink
	item.Device = file.Mode()&os.ModeDevice == os.ModeDevice
	item.NamedPipe = file.Mode()&os.ModeNamedPipe == os.ModeNamedPipe
	item.Socket = file.Mode()&os.ModeSocket == os.ModeSocket
	item.Setuid = file.Mode()&os.ModeSetuid == os.ModeSetuid
	item.CharDevice = file.Mode()&os.ModeCharDevice == os.ModeCharDevice
	item.Sticky = file.Mode()&os.ModeSticky == os.ModeSticky

	item.OwnerRead = file.Mode().Perm()&0400 == 0400
	item.OwnerWrite = file.Mode().Perm()&0200 == 0200
	item.OwnerExecute = file.Mode().Perm()&0100 == 0100
	item.GroupRead = file.Mode().Perm()&0040 == 0040
	item.GroupWrite = file.Mode().Perm()&0020 == 0020
	item.GroupExecute = file.Mode().Perm()&0010 == 0010
	item.OtherRead = file.Mode().Perm()&0004 == 0004
	item.OtherWrite = file.Mode().Perm()&0002 == 0002
	item.OtherExecute = file.Mode().Perm()&0001 == 0001
	item.Permissions = file.Mode().Perm().String()

	item.Mtime, _ = file.Sys().(*syscall.Stat_t).Mtimespec.Unix()
	item.Atime, _ = file.Sys().(*syscall.Stat_t).Atimespec.Unix()
	item.Ctime, _ = file.Sys().(*syscall.Stat_t).Ctimespec.Unix()

	item.Uid = file.Sys().(*syscall.Stat_t).Uid
	item.Gid = file.Sys().(*syscall.Stat_t).Gid

	json, error := json.Marshal(item)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(string(json))
}
