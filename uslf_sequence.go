// Copyright 2016 Albert Nigmatzianov. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package id3v2

// uslfSequence stores several USLT frames.
// Key for uslfSequence is language and content descriptor,
// so there is only one USLT frame with the same language and
// content descriptor.
//
// It's only needed for internal operations. Users of library id3v2 should not
// use any sequence in no case.
type uslfSequence struct {
	sequence map[string]UnsynchronisedLyricsFrame
}

func newUSLFSequence() sequencer {
	return &uslfSequence{
		sequence: make(map[string]UnsynchronisedLyricsFrame),
	}
}

func (us uslfSequence) Frames() []Framer {
	frames := make([]Framer, 0, len(us.sequence))
	for _, f := range us.sequence {
		frames = append(frames, f)
	}
	return frames
}

func (us *uslfSequence) AddFrame(f Framer) {
	uslf := f.(UnsynchronisedLyricsFrame)
	id := uslf.Language + uslf.ContentDescriptor
	us.sequence[id] = uslf
}