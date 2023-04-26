package main

import (
	"context"
	_ "embed"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/cristalhq/acmd"
	"github.com/mholt/archiver/v4"
)

type Bundle struct {
	Filename string
}

var VERSION = "0.2.9"

var (
	ErrFailedToCompile = errors.New("failed to compile plugin")
	ErrNoPluginSo      = errors.New("can't find plugin.so")
)

func main() {
	cmds := []acmd.Command{
		{
			Name:        "bundle",
			Description: "bundels a plugin",
			ExecFunc:    bundle,
		},
	}

	r := acmd.RunnerOf(cmds, acmd.Config{
		AppName:        "seekr-cli",
		AppDescription: "seekr cli tooling",
		Version:        VERSION,
		// Context - if nil `signal.Notify` will be used
		// Args - if nil `os.Args[1:]` will be used
		// Usage - if nil default print will be used
	})

	if err := r.Run(); err != nil {
		r.Exit(err)
	}
}

func (bundle Bundle) createArchive() error {
	if bundle.Filename == "" {
		return ErrNoPluginSo // FIXME wrong error message
	}
	files, err := archiver.FilesFromDisk(nil, map[string]string{
		"./plugin.so": "api/plugin.so",
	})
	if err != nil {
		return err
	}

	out, err := os.Create(bundle.Filename)
	if err != nil {
		return err
	}
	defer out.Close()

	format := archiver.CompressedArchive{
		Compression: archiver.Gz{},
		Archival:    archiver.Tar{},
	}

	err = format.Archive(context.Background(), out, files)
	if err != nil {
		return err
	}
	return err
}
func bundle(ctx context.Context, args []string) error {
	fmt.Printf("Bundeling plugin")
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", "plugin.so")
	if err := cmd.Run(); err != nil {
		return ErrFailedToCompile
	}
	bundle := Bundle{
		Filename: "plugin.bundle",
	}
	err := bundle.createArchive()
	if err != nil {
		return err
	}
	return nil
}
