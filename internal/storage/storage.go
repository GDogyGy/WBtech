package storage

import "errors"

var (
	JsonNotValid      = errors.New("json not valid")
	TrackNumbNotExist = errors.New("track number not exist")
)
