package main

type Client interface {
	Read(path, env string) (string, error)
	ReadBulk(path string) (map[string]string, error)
	Write(path, value string) error
	WriteBulk(path string, data map[string]string) error
}
