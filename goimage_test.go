package goimage

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"hago-base-web-realtimetranscode/transcode/processor/goimage/goimage/heif"
	"image"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const testdataDir = "testdata"
const testoutDir = "testout"

func Test_Image(t *testing.T) {
	os.RemoveAll(testoutDir)
	os.MkdirAll(testoutDir, 0755)

	t.Log("libheif version:", heif.GetVersion())

	tests := []struct {
		name string
	}{
		{"test11.png"},
		{"test12.png"},
		{"test13.png"},
		{"test14.png"},
		{"test15.png"},
		//{""},
		{"test21.jpeg"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run("circle", func(t *testing.T) {
				img := testDecodeImg(t, tt.name)
				img = Circle(img, 0)
				testFormat2heif(t, img, tt.name, "circle")
			})

			t.Run("resize", func(t *testing.T) {
				img := testDecodeImg(t, tt.name)
				img = Resize(img, 720, 0)
				testFormat2heif(t, img, tt.name, "resize")
			})
		})
	}
}

func testDecodeImg(t *testing.T, name string) image.Image {
	r := readFile(filepath.Join(testdataDir, name))
	defer r.Close()
	img, magic, err := image.Decode(r)
	require.NoError(t, err)
	t.Log("decode image:", name, "type=", magic)
	return img
}
func testFormat2heif(t *testing.T, img image.Image, name string, action string) {
	t.Run("format2heif_"+action, func(t *testing.T) {
		name = strings.TrimSuffix(name, filepath.Ext(name))
		name += "_" + action + ".heif"
		outFilename := filepath.Join(testoutDir, name)
		buf := bytes.NewBuffer(nil)

		err := heif.Encode(buf, img)
		require.NoError(t, err)
		err = ioutil.WriteFile(outFilename, buf.Bytes(), 0644)
		require.NoError(t, err)
	})

}

func readFile(filename string) io.ReadCloser {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return file
}
