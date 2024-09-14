package config

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
)

const (
	read_buf_size = 1024
)

// auto import config
func ImportConfig[CONFIG_TYPE any](config *CONFIG_TYPE, config_path string) error {
	if config == nil {
		return errors.New("config: There is no corresponding instance in config")
	}

	file, err := os.OpenFile(config_path, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	if reader == nil {
		return errors.New("config: Config reader create failed")
	}

	var json_data []byte
	for {
		read_buf := make([]byte, read_buf_size)
		if pan := recover(); pan != nil {
			return errors.New("config: Failed to create the memory allocation required for the read buffer")
		}
		read_len, err := reader.Read(read_buf)
		if err == nil {
			if read_len < len(read_buf) {
				read_buf = read_buf[:read_len]
				json_data = append(json_data, read_buf...)
				if pan := recover(); pan != nil {
					return errors.New("config: Failed to create the memory allocation required for the json data buffer")
				}
				break
			} else {
				json_data = append(json_data, read_buf...)
				if pan := recover(); pan != nil {
					return errors.New("config: Failed to create the memory allocation required for the json data buffer")
				}
			}
		} else if err == io.EOF {
			break
		} else {
			return err
		}
	}

	err = json.Unmarshal(json_data, config)
	if err != nil {
		return err
	}

	return nil
}

// auto export config
func ExportConfig[CONFIG_TYPE any](config *CONFIG_TYPE, config_path string) error {
	if config == nil {
		return errors.New("config: There is no corresponding instance in config")
	}

	file, err := os.OpenFile(config_path, os.O_CREATE|os.O_WRONLY, 0)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	if writer == nil {
		return errors.New("config: Config writer create failed")
	}

	json_data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	_, err = writer.Write(json_data)
	if err != nil {
		return err
	}
	writer.Flush()

	return nil
}
