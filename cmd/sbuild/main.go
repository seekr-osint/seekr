package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
	//"syscall"
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

		for build.Watch {
			time.Sleep(10 * time.Millisecond)
			goCmd := exec.Command("go", "run", "main.go")
			goCmd.Stdout = os.Stdout
			goCmd.Stderr = os.Stderr
			fmt.Printf("go run main.go\n")
			err := goCmd.Start()
			if err != nil {
				fmt.Printf("Failed to start command: %v\n", err)
				return err
			}
			signalChannel := make(chan os.Signal, 1)
			signal.Notify(signalChannel, syscall.SIGINT)

			go func(cmd *exec.Cmd) {
				sig := <-signalChannel
				fmt.Printf("Received signal: %v\n", sig)

				cmd.Process.Signal(sig)
			}(goCmd)

			err = goCmd.Wait()
			if err != nil {
				fmt.Printf("Command failed: %v\n", err)
				//return err
			}

			fmt.Println("Program finished successfully")
		}
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

		tscCmd := exec.Command("tsc", "--project", build.TscProjectDir, "--watch", fmt.Sprintf("%v", build.Watch))
		tscCmd.Stderr = os.Stderr
		var err error
		if build.Watch {
			err = tscCmd.Start()
		} else {
			tscCmd.Stdout = os.Stdout
			err = tscCmd.Run()
		}
		if err != nil {
			fmt.Printf("error: %v\n", err)
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

	flag.BoolVar(&build.Watch, "watch", false, "delete the typescript files")

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
	fmt.Printf("buildign Go...\n")
	err = build.buildGo()
	if err != nil {
		fmt.Printf("Failed to build go: %v\n", err)
		return
	}

}
