/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package asciinema

import (
	"encoding/json"
	"errors"
	"io"
	"time"
)

const (
	maxBuffSize  = 1024
	maxFileSize  = 100 * 1000 * 1000 // 100M
	version      = 2
	defaultShell = "/bin/bash"
	defaultTerm  = "xterm"
)

var (
	newLine = []byte{'\n'}
)

func NewWriter(w io.Writer, opts ...Option) *Writer {
	conf := Config{
		Width:    80,
		Height:   40,
		EnvShell: defaultShell,
		EnvTerm:  defaultTerm,
	}
	for _, setter := range opts {
		setter(&conf)
	}
	return &Writer{
		Config:        conf,
		TimestampNano: conf.Timestamp.UnixNano(),
		writer:        w,
		limit:         maxFileSize,
		WriteBuff:     make([]byte, 0, maxBuffSize),
	}
}

type Writer struct {
	Config
	TimestampNano int64
	writer        io.Writer
	limit         int
	written       int
	WriteBuff     []byte
}

func (w *Writer) WriteHeader() error {
	header := Header{
		Version:   version,
		Width:     w.Width,
		Height:    w.Height,
		Timestamp: w.Timestamp.Unix(),
		Title:     w.Title,
		Env: Env{
			Shell: w.EnvShell,
			Term:  w.EnvTerm,
		},
	}
	raw, err := json.Marshal(header)
	if err != nil {
		return err
	}
	_, err = w.Write(raw)
	if err != nil {
		return err
	}
	_, err = w.Write(newLine)
	return err
}

func (w *Writer) WriteRow(p []byte) error {
	now := time.Now().UnixNano()
	ts := float64(now-w.TimestampNano) / 1000 / 1000 / 1000
	return w.WriteStdout(ts, p)
}

func (w *Writer) WriteStdout(ts float64, data []byte) error {
	row := []interface{}{ts, "o", string(data)}
	raw, err := json.Marshal(row)
	if err != nil {
		return err
	}

	if len(raw) > maxBuffSize { //读取的数据大于最大缓存
		//先缓存写入再将读取的数据写入文件
		_, err := w.Write(w.WriteBuff)
		if err != nil {
			return err
		}
		w.WriteBuff = w.WriteBuff[:0]
		_, err = w.Write(raw)
		if err != nil {
			return err
		}
		_, err = w.Write(newLine)
		return err
	} else if len(raw)+len(w.WriteBuff) > maxBuffSize { //检查是否达到最大容量
		//buff写入文件,清空buff,raw写入buff
		_, err := w.Write(w.WriteBuff)
		if err != nil {
			return err
		}
		w.WriteBuff = w.WriteBuff[:0]
		raw := append(raw, newLine...)
		w.WriteBuff = append(w.WriteBuff, raw...)
	} else {
		raw := append(raw, newLine...)
		w.WriteBuff = append(w.WriteBuff, raw...)
	}
	//w.WriteBuff = append(w.WriteBuff, raw...)
	//w.WriteBuff = append(w.WriteBuff, newLine...)
	//
	//_, err = w.Write(raw)
	//if err != nil {
	//	return err
	//}
	//_, err = w.Write(newLine)
	//return err
	return nil
}

type Header struct {
	Version   int    `json:"version"`
	Width     uint16 `json:"width"`
	Height    uint16 `json:"height"`
	Timestamp int64  `json:"timestamp"`
	Title     string `json:"title"`
	Env       Env    `json:"env"`
}

type Env struct {
	Shell string `json:"SHELL"`
	Term  string `json:"TERM"`
}

func (w *Writer) Write(p []byte) (n int, err error) {
	remainingSpace := w.limit - w.written
	if remainingSpace <= 0 {
		return 0, errors.New("Exceeds the file size")
	}

	if len(p) > remainingSpace {
		p = p[:remainingSpace]
	}

	n, err = w.writer.Write(p)
	w.written += n
	return n, err
}
