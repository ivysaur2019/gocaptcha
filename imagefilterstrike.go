// Copyright 2013 hanguofeng. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package gocaptcha

import ()

//ImageFilter is the interface of image filter
type ImageFilterStrike struct {
	StrikeLineNum int
}

func (filter *ImageFilterStrike) proc(cimage *CImage) {
	cimage.strikeThrough(filter.StrikeLineNum)
}
