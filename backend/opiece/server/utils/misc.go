package utils

import (
	"archive/zip"
	"bytes"
	"errors"
	"github.com/google/uuid"
	"opiece/server/middleware"
	"opiece/server/model"
)

type Target string

func (t Target) In(x []string) bool {
	for _, val := range x {
		if t == Target(val) {
			return true
		}
	}
	return false
}

func GetBool(s string) bool {
	if s == "true" {
		return true
	} else if s == "false" {
		return false
	} else if s == "" {
		return false
	}
	return false
}

func GetUUIDFromToken(token string) (uuid.UUID, error) {
	jwt := middleware.NewJWT()
	claims, err := jwt.ParseToken(token)
	if err != nil {
		return uuid.UUID{}, err
	}
	return claims.UUID, err
}

func CompressFileToZip(files []model.BinaryFile) *bytes.Buffer{
	buf := &bytes.Buffer{}
	zipWriter := zip.NewWriter(buf)
	defer func() {
		err := zipWriter.Close()
		if err != nil {
			panic(errors.New("关闭文件句柄错误: "+ err.Error()))
		}
	}()
	for _, file := range files {
		zipFile, err := zipWriter.Create(file.FileName)
		if err != nil {
			panic(err)
		}
		_, err = zipFile.Write(file.FileData)
		if err != nil {
			// TODO: handle errors
			panic(err)
		}
	}
	return buf
}