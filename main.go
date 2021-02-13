package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	srcBase := os.Args[1]
	dstBase := os.Args[2]
	mvTarget := os.Args[3]

	srcBaseAbs, err := filepath.Abs(srcBase)
	if err != nil {
		log.Fatalln(err)
	}
	mvTargetAbs, err := filepath.Abs(mvTarget)
	if err != nil {
		log.Fatalln(err)
	}

	if strings.Index(mvTargetAbs, srcBaseAbs) != 0 {
		return
	}

	targetRelPath := strings.TrimPrefix(mvTargetAbs, srcBaseAbs)

	targetPath := filepath.Join(dstBase, targetRelPath)
	targetDir := filepath.Join(dstBase, filepath.Dir(targetRelPath))

	os.MkdirAll(targetDir, os.ModePerm)

	err = os.Rename(mvTargetAbs, targetPath)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(targetDir)
	fmt.Println(targetPath)
}
