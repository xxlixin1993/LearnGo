package main

import (
	"bytes"
	"fmt"
	"os"
	"bufio"
	"io"
	"sync"
)

var (
	DEFAULT_COMMENT = []byte{'#'}
	DEFAULT_COMMENT_SEM = []byte{';'}
)

type Config struct {
	// map is not safe.
	sync.RWMutex
	// Section:key=value
	data map[string]map[string]string
}

func (c *Config) parse(fname string) (err error) {
	c.Lock()
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer c.Unlock()
	defer f.Close()

	buf := bufio.NewReader(f)

	var section string
	var lineNum int

	for {
		lineNum++
		line, _, err := buf.ReadLine()
		if err == io.EOF {
			break
		} else if bytes.Equal(line, []byte{}) {
			continue
		} else if err != nil {
			return err
		}

		line = bytes.TrimSpace(line)
		switch {
		case bytes.HasPrefix(line, DEFAULT_COMMENT):
			continue
		case bytes.HasPrefix(line,DEFAULT_COMMENT_SEM):
			continue
		case bytes.HasPrefix(line, []byte{'['}) && bytes.HasSuffix(line, []byte{']'}):
			section = string(line[1 : len(line)-1])
		default:
			optionVal := bytes.SplitN(line, []byte{'='}, 2)
			if len(optionVal) != 2 {
				return fmt.Errorf("parse %s the content error : line %d , %s = ? ", fname, lineNum, optionVal[0])
			}
			option := bytes.TrimSpace(optionVal[0])
			value := bytes.TrimSpace(optionVal[1])
			fmt.Printf("config %s %s %s", section, string(option), string(value))
			//c.AddConfig(section, string(option), string(value))
		}
	}

	return nil
}

func main() {
	(&Config{}).parse("/Users/lixin/go/src/LearnGo/demo/config.ini")
}