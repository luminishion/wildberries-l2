package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var patterns map[string][]string

func init() {
	patterns = map[string][]string{
		"html": []string{
			`src\s*=\s*"(.+?)"`,
			`href\s*=\s*"(.+?)"`,
		},
	}
}

func getPathExt(addr string) (string, string, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return "", "", err
	}

	path := u.Path

	pos := strings.LastIndex(path, "/")
	if pos != -1 {
		pos := strings.LastIndex(path[pos+1:len(path)], ".")
		if pos != -1 {
			return path, path[pos+1 : len(path)], nil
		}
	}

	ret, err := u.Parse("index.html")
	if err != nil {
		return "", "", err
	}

	return ret.Path, "html", nil
}

func download(addr string) ([]byte, error) {
	resp, err := http.Get(addr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func NewDownloader() *Downloader {
	return &Downloader{
		done: make(map[string]struct{}),
	}
}

type Downloader struct {
	queue   []string
	done    map[string]struct{}
	handler func(string, []byte)
}

func (d *Downloader) addAddr(addr string) {
	if _, ok := d.done[addr]; ok {
		return
	}
	d.done[addr] = struct{}{}

	d.queue = append(d.queue, addr)
}

func (d *Downloader) parseAddr(baseAddr string, addr string) {
	u, err := url.Parse(addr)
	if err != nil {
		return
	}

	if u.IsAbs() {
		return
	}

	base, err := url.Parse(baseAddr)
	if err != nil {
		return
	}
	if addr[0] == '/' {
		base.Path = ""
	}

	ret, err := base.Parse(addr)
	if err != nil {
		return
	}
	ret.RawQuery = ""
	ret.Fragment = ""

	d.addAddr(ret.String())
}

func (d *Downloader) parseData(baseAddr string, ext string, data []byte) {
	regs, ok := patterns[ext]
	if !ok {
		return
	}

	for _, reg := range regs {
		r := regexp.MustCompile(reg)
		matches := r.FindAllStringSubmatch(string(data), -1)

		for _, v := range matches {
			d.parseAddr(baseAddr, v[1])
		}
	}
}

func (d *Downloader) downloadAddr(addr string) {
	data, err := download(addr)
	if err != nil {
		return
	}

	path, ext, err := getPathExt(addr)
	if err != nil {
		return
	}

	d.handler(path, data)

	d.parseData(addr, ext, data)
}

func (d *Downloader) Run(addr string) {
	d.queue = append(d.queue, addr)

	for len(d.queue) > 0 {
		i := len(d.queue) - 1
		addr := d.queue[i]
		d.queue = d.queue[:i]

		d.downloadAddr(addr)
	}
}

func (d *Downloader) SetHandler(fn func(string, []byte)) {
	d.handler = fn
}

func main() {
	d := NewDownloader()
	d.SetHandler(func(k string, v []byte) {
		fmt.Println(k, len(v))
	})
	d.Run("https://yourbasic.org/golang/regexp-cheat-sheet/")
}
