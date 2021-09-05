package gutil

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type FsStorage struct {
	storageDir string
	nsMap      sync.Map
}

func NewFsStorage(storageDir string) *FsStorage {
	_ = os.Mkdir(storageDir, os.FileMode(0755))
	f := &FsStorage{}
	f.storageDir = storageDir
	return f
}

func (f *FsStorage) AddNamespace(namespace string) {
	nsPath := filepath.FromSlash(f.storageDir + "/" + namespace)
	_ = os.Mkdir(nsPath, os.FileMode(0755))
	f.nsMap.Store(namespace, nsPath)
}

func (f *FsStorage) Add(ns, key string, v interface{}) error {
	keyFile, err := f.getKeyFile(ns, key)
	if err != nil {
		return err
	}
	file, err := os.OpenFile(keyFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = file.Write(bytes)
	return err
}

func (f *FsStorage) Del(ns, key string) error {
	keyFile, err := f.getKeyFile(ns, key)
	if err != nil {
		return err
	}
	return os.Remove(keyFile)
}

func (f *FsStorage) Get(ns, key string) ([]byte, error) {
	keyFile, err := f.getKeyFile(ns, key)
	if err != nil {
		return nil, err
	}
	return os.ReadFile(keyFile)
}

func (f *FsStorage) GetKeyList(ns string) ([]string, error) {
	nsPath, err := f.getNsPath(ns)
	if err != nil {
		return nil, err
	}
	res := make([]string, 0)
	// iterate
	err = filepath.Walk(nsPath, func(path string, info fs.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".json" {
			res = append(res, strings.TrimSuffix(info.Name(), ".json"))
			return nil
		}
		return nil
	})
	return res, err
}

func (f *FsStorage) getKeyFile(ns string, key string) (string, error) {
	nsPath, err := f.getNsPath(ns)
	if err != nil {
		return "", err
	}
	return filepath.FromSlash(nsPath + "/" + key + ".json"), nil
}

func (f *FsStorage) getNsPath(ns string) (string, error) {
	nsPath, ok := f.nsMap.Load(ns)
	if !ok {
		return "", errors.New("no such namespace")
	}
	nsPathStr, ok := nsPath.(string)
	if !ok {
		return "", errors.New("can not reach here")
	}
	return nsPathStr, nil
}
