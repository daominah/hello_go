package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	dirPath := `/media/tungdt/WindowsData/syncthing/Master_Duel_art_full/upscayled_2048_png`
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("error os.ReadDir: %v", err)
	}

	regExpYgopass8, err := regexp.Compile(`_[0-9]{7,8}_`)
	if err != nil {
		log.Fatalf("error regexp.Compile: %v", err)
	}

	for i, f := range entries {
		log.Printf("i %v ________________", i)
		if f.IsDir() {
			continue
		}

		ext := filepath.Ext(f.Name()) // ext includes the dot
		nameWOE := strings.TrimSuffix(f.Name(), ext)
		//log.Printf("file nameWOE: %v, ext: %v", nameWOE, ext)
		if strings.HasSuffix(nameWOE, "_up2048") {
			continue
		}

		nameRemovedPass8 := regExpYgopass8.ReplaceAllString(nameWOE, "_")
		newName := fmt.Sprintf("%v_up2048%v", nameRemovedPass8, ext)

		oldFullPath := filepath.Join(dirPath, f.Name())
		newFullPath := filepath.Join(dirPath, newName)
		log.Printf("renaming: %v to %v", oldFullPath, newFullPath)
		err := os.Rename(oldFullPath, newFullPath)
		if err != nil {
			log.Printf("error os.Rename: %v", err)
		}
	}
}
