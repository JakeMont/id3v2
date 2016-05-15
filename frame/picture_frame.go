package frame

import (
	"bytes"
	"errors"
	"github.com/bogem/id3v2/util"
	"os"
)

var (
	PictureTypes = [...]string{
		"Other",
		"File icon",
		"Other file icon",
		"Cover (front)",
		"Cover (back)",
		"Leaflet page",
		"Media",
		"Lead artist",
		"Artist",
		"Conductor",
		"Band/Orchestra",
		"Composer",
		"Lyricist/text writer",
		"Recording Location",
		"During recording",
		"During performance",
		"Movie/video screen capture",
		"A bright coloured fish",
		"Illustration",
		"Band/artist logotype",
		"Publisher/Studio logotype",
	}
)

type PictureFramer interface {
	Framer

	Description() string
	SetDescription(string)

	MimeType() string
	SetMimeType(string)

	Picture() []byte
	SetPicture([]byte) error

	PictureType() string
	SetPictureType(string) error
}

type PictureFrame struct {
	description string
	id          string
	mimeType    string
	picture     bytes.Buffer
	pictureType byte
}

func NewPictureFrame() *PictureFrame {
	pf := new(PictureFrame)
	pf.SetID("APIC")
	return pf
}

func (pf PictureFrame) Form() ([]byte, error) {
	var b bytes.Buffer
	if header, err := FormFrameHeader(&pf); err != nil {
		return nil, err
	} else {
		b.Write(header)
	}
	b.WriteByte(util.NativeEncoding)
	b.WriteString(pf.mimeType)
	b.WriteByte(0)
	b.WriteByte(pf.pictureType)
	b.WriteString(pf.description)
	b.WriteByte(0)
	b.Write(pf.Picture())
	return b.Bytes(), nil
}

func (pf PictureFrame) ID() string {
	return pf.id
}

func (pf *PictureFrame) SetID(id string) {
	pf.id = id
}

func (pf PictureFrame) Size() uint32 {
	encodingSize, pictureTypeSize := 1, 1
	size := uint32(encodingSize + len(pf.mimeType) + 1 + pictureTypeSize + len(pf.description) + 1 + pf.picture.Len())
	return size
}

func (pf PictureFrame) Description() string {
	return pf.description
}

func (pf *PictureFrame) SetDescription(desc string) {
	pf.description = desc
}

func (pf PictureFrame) MimeType() string {
	return pf.mimeType
}

func (pf *PictureFrame) SetMimeType(mt string) {
	pf.mimeType = mt
}

func (pf PictureFrame) Picture() []byte {
	return pf.picture.Bytes()
}

func (pf *PictureFrame) SetPicture(b []byte) error {
	pf.picture.Reset()
	if _, err := pf.picture.Write(b); err != nil {
		return err
	}
	return nil
}

func (pf *PictureFrame) SetPictureFromFile(file *os.File) error {
	pf.picture.Reset()
	if _, err := pf.picture.ReadFrom(file); err != nil {
		return err
	}
	return nil
}

func (pf PictureFrame) PictureType() string {
	return PictureTypes[pf.pictureType]
}

func (pf *PictureFrame) SetPictureType(pt string) error {
	for k, v := range PictureTypes {
		if v == pt {
			pf.pictureType = byte(k)
			return nil
		}
	}
	return errors.New("Unsupported picture type")
}