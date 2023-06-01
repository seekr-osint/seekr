package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

type Build struct {
	GoRun         bool
	Watch         bool
	TscProjectDir string
}

func moveFilesToParentDir(dirPath string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		src := filepath.Join(dirPath, file.Name())
		dst := filepath.Join(filepath.Dir(dirPath), file.Name())

		err = os.Rename(src, dst)
		if err != nil {
			return err
		}

		fmt.Printf("Moved file: %s\n", file.Name())
	}

	return nil
}

func (build Build) buildGo() error {
	if !build.GoRun {
		goCmd := exec.Command("go", "build", "main.go")
		goCmd.Stdout = os.Stdout
		goCmd.Stderr = os.Stderr

		err := goCmd.Run()
		if err != nil {
			return err
		}
	} else {

		goCmd := exec.Command("go", "run", "main.go")
		goCmd.Stdout = os.Stdout
		goCmd.Stderr = os.Stderr

		err := goCmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func (build Build) compileTS() error {
	tscCmd := exec.Command("tsc", "--project", build.TscProjectDir, "--watch", "false")
	tscCmd.Stdout = os.Stdout
	tscCmd.Stderr = os.Stderr

	err := tscCmd.Run()
	if err != nil {
		return err
	}
	//moveFilesToParentDir(filepath.Join(build.TscProjectDir, "dist", "ts"))
	if err != nil {
		return err
	}

	return nil
}

func (build Build) generate() error {
	generateCmd := exec.Command("go", "generate", "./...")
	generateCmd.Stdout = os.Stdout
	generateCmd.Stderr = os.Stderr

	err := generateCmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func main() {

	build := Build{
		TscProjectDir: "./web",
		Watch:         false,
	}
	flag.BoolVar(&build.GoRun, "run", false, "`go run` instead of `go build`")

	flag.Parse()

	fmt.Printf("running go generate...\n")
	err := build.generate()
	if err != nil {
		fmt.Printf("Failed to run go generate: %v\n", err)
		return
	}

	fmt.Printf("transpiling TypeScript...\n")
	err = build.compileTS()
	if err != nil {
		fmt.Printf("Failed to compile TypeScript: %v\n", err)
		return
	}

	fmt.Printf("building Go...\n")
	err = build.buildGo()
	if err != nil {
		fmt.Printf("Failed to build/run Go application: %v\n", err)
		return
	}
}
