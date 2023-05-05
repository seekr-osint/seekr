package main

import (
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"

	"github.com/cristalhq/acmd"
	"github.com/mholt/archiver/v4"
)

type Bundle struct {
	Filename string
	SrcFiles []string
}

var VERSION = "0.2.9"

var (
	ErrFailedToCompile       = errors.New("failed to compile plugin")
	ErrFailedToLoadManifesto = errors.New("failed to load manifesto")
	ErrNoPluginSo            = errors.New("can't find plugin.so")
	ErrNoMainPackage         = errors.New("err not a main package")
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
	afiles := map[string]string{
		"./plugin.so": "api/plugin.so",
	}
	for _, v := range bundle.SrcFiles {
		afiles[v] = fmt.Sprintf("api/src/%s", v)
	}
	files, err := archiver.FilesFromDisk(nil, afiles)
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

type Manifesto struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

func (m *Manifesto) validateName() error {
	// Regular expression that matches alphanumeric characters, underscores, and hyphens
	validName := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	if !validName.MatchString(m.Name) {
		return errors.New("name contains invalid characters")
	}
	return nil
}

func loadManifesto() (*Manifesto, error) {
	file, err := ioutil.ReadFile("manifesto.json")
	if err != nil {
		return nil, errors.New("failed to read manifesto.json")
	}

	var manifesto Manifesto
	err = json.Unmarshal(file, &manifesto)
	if err != nil {
		return nil, errors.New("failed to unmarshal manifesto.json")
	}

	if err := manifesto.validateName(); err != nil {
		return nil, err
	}

	return &manifesto, nil
}

func BuildModule(path string) ([]string, error) {
	pkg, err := build.ImportDir(path, 0)
	if err != nil {
		return []string{}, err
	}
	if pkg.Name != "main" {
		return []string{}, ErrNoMainPackage
	}
	files := []string{}
	for _, file := range pkg.GoFiles {
		files = append(files, fmt.Sprintf("./%s", file))
	}
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", "plugin.so", pkg.ImportPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return files, cmd.Run()
}

func bundle(ctx context.Context, args []string) error {
	fmt.Printf("Loading Manifesto\n")
	manifesto, err := loadManifesto()
	if err != nil {
		return ErrFailedToLoadManifesto
	}
	fmt.Printf("Building plugin\n")
	//cmd := exec.Command("go", "build", "-buildmode=plugin", "-o", "plugin.so")
	files, err := BuildModule(".")
	if err != nil {
		return ErrFailedToCompile
	}
	bundle := Bundle{
		Filename: fmt.Sprintf("%s.bundle", manifesto.Name),
		SrcFiles: files,
	}
	err = bundle.createArchive()
	if err != nil {
		return err
	}
	return nil
}
