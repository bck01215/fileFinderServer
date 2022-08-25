package fileFinder

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"sort"
	"sync"

	"github.com/google/fscrypt/filesystem"
	"github.com/sirupsen/logrus"
)

var device string

type fileDisplay struct {
	sync.RWMutex
	Files []FileDisplay
}

var files fileDisplay

type FileDisplay struct {
	Size int64  `json:"size"`
	Path string `json:"path"`
}
type bySize []FileDisplay

func (a bySize) Len() int           { return len(a) }
func (a bySize) Less(i, j int) bool { return a[i].Size < a[j].Size }
func (a bySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func GetTopXFiles(mountpoint string, limit int) ([]FileDisplay, error) {
	log.SetOutput(io.Discard)
	if mountpoint == "" {
		return nil, fmt.Errorf("Path cannot be empty")
	}
	if limit < 1 {
		return nil, fmt.Errorf("Limit must be 1 or greater")

	}
	if mountpoint[len(mountpoint)-1:] != "/" {
		mountpoint = mountpoint + "/"
	}
	mount, err := filesystem.FindMount(mountpoint)
	if err != nil {
		return nil, err
	}
	device = mount.Device

	entries, err := os.ReadDir(mountpoint)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	getFiles(mountpoint, entries, &wg)
	wg.Wait()
	sort.Sort(bySize(files.Files))
	var shortFiles []FileDisplay
	if len(files.Files) > limit {
		shortFiles = files.Files[len(files.Files)-limit:]
	} else {
		shortFiles = files.Files
	}
	files = fileDisplay{}
	return shortFiles, nil

}

func getFiles(start string, entries []fs.DirEntry, wg *sync.WaitGroup) {
	for _, entry := range entries {
		wg.Add(1)
		go handleEntry(start, entry, wg)
	}

}

func handleEntry(start string, entry fs.DirEntry, wg *sync.WaitGroup) {
	defer wg.Done()
	var file FileDisplay
	mount, err := filesystem.FindMount(start + entry.Name())
	if err != nil {
		logrus.Errorln(err, start+entry.Name())
		return
	}
	if mount.Device == device {
		if entry.Type().IsRegular() {
			fileInfo, err := os.Stat(start + entry.Name())
			if err != nil {
				logrus.Errorln(err, start+entry.Name())
				return
			}
			file.Path = start + entry.Name()
			file.Size = fileInfo.Size()
			files.Append(file)
		} else if entry.IsDir() {
			entries, err := os.ReadDir(start + entry.Name())
			if err != nil {
				logrus.Errorln(err, start+entry.Name())
				return
			}
			logrus.Info("Searching ", start+entry.Name())
			getFiles(start+entry.Name()+"/", entries, wg)
		}
	}

}

func (f *FileDisplay) DisplaySizeIEC() string {
	const unit = 1024
	b := f.Size
	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f%ciB",
		float64(b)/float64(div), "KMGTPE"[exp])
}

func (f *FileDisplay) Delete() error {
	err := os.Remove(f.Path)
	if err != nil {
		return err
	}
	return nil
}

func (fd *fileDisplay) Append(item FileDisplay) {
	fd.Lock()
	defer fd.Unlock()

	fd.Files = append(fd.Files, item)
}
