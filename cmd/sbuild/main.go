package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Build struct {
	GoRun         bool
	Watch         bool
	TscProjectDir string
	CleanGo       bool
	CleanTsc      bool
	Tsc           bool
	GoGenerate    bool
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

func deleteFilesInFolder(folderPath string) error {
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			err := os.Remove(path)
			if err != nil {
				return err
			}
			fmt.Printf("Deleted file: %s\n", path)
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("All files deleted successfully.")
	return nil
}

func (build Build) compileTS() error {
	if build.Tsc {
		if build.CleanTsc {
			err := deleteFilesInFolder(filepath.Join(build.TscProjectDir, "dist"))
			if err != nil {
				return err
			}
		}
		tscCmd := exec.Command("tsc", "--project", build.TscProjectDir, "--watch", "false")
		tscCmd.Stdout = os.Stdout
		tscCmd.Stderr = os.Stderr

		err := tscCmd.Run()
		if err != nil {
			return err
		}
	}

	return nil
}

func (build Build) generate() error {

	if build.GoGenerate {
		if build.CleanGo {
			err := deleteFilesInFolder(filepath.Join("web", "ts-gen"))
			if err != nil {
				return err
			}
		}
		generateCmd := exec.Command("go", "generate", "./...")
		generateCmd.Stdout = os.Stdout
		generateCmd.Stderr = os.Stderr

		err := generateCmd.Run()
		if err != nil {
			return err
		}

	}
	return nil
}

func main() {

	build := Build{
		TscProjectDir: "./web",
		Watch:         false,
	}
	flag.BoolVar(&build.GoRun, "run", false, "`go run` instead of `go build`")

	flag.BoolVar(&build.CleanTsc, "clean", false, "delete the typescript files")

	flag.BoolVar(&build.CleanGo, "clean-go", true, "delete the typescript files")

	flag.BoolVar(&build.Tsc, "tsc", true, "delete the typescript files")

	flag.BoolVar(&build.GoGenerate, "generate", true, "delete the typescript files")

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
