package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/hanwen/go-fuse/v2/fs"
	"github.com/hanwen/go-fuse/v2/fuse"
)

func main() {
	root := &HogeRoot{}
	mnt := "/tmp/x"
	_ = os.Mkdir(mnt, 0755)
	server, err := fs.Mount(mnt, root, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("to unmount: fusermount -u %s\n", mnt)
	server.Wait()
}

type HogeFile struct {
	fs.Inode
	content []byte
}

type HogeRoot struct {
	fs.Inode
}

var _ = (fs.NodeOpener)((*HogeFile)(nil))

var _ = (fs.NodeGetattrer)((*HogeFile)(nil))

func (f *HogeFile) Open(ctx context.Context, flags uint32) (fs.FileHandle, uint32, syscall.Errno) {
	return nil, fuse.FOPEN_KEEP_CACHE, fs.OK
}

func (f *HogeFile) Getattr(ctx context.Context, fh fs.FileHandle, out *fuse.AttrOut) syscall.Errno {
	out.Size = uint64(len(f.content))
	return 0
}

func (f *HogeFile) Read(ctx context.Context, fh fs.FileHandle, dest []byte, off int64) (fuse.ReadResult, syscall.Errno) {
	return fuse.ReadResultData(f.content), fs.OK
}

// The root populates the tree in its OnAdd method
var _ = (fs.NodeOnAdder)((*HogeRoot)(nil))

func (r *HogeRoot) OnAdd(ctx context.Context) {
	ch := r.Inode.NewPersistentInode(ctx, &HogeFile{
		content: []byte("hogehoge"),
	}, fs.StableAttr{})
	r.Inode.AddChild("hoge", ch, true)
	ch = r.Inode.NewPersistentInode(ctx, &HogeFile{
		content: []byte("fugafuga"),
	}, fs.StableAttr{})
	r.Inode.AddChild("fuga", ch, true)
}
