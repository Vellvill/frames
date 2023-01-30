package s3

import (
	"fmt"
	"strconv"
	"strings"
)

type awsKey string

// key is parentID + sequence
const key awsKey = "%d_%d"

type AWSObject struct {
	ParentID, Sequence int64
	Obj                []byte
}

func (a awsKey) getParentIDAndSequenceFromKey() (int64, int64, error) {
	split := strings.Split(string(a), "_")

	if len(split) < 2 {
		return 0, 0, fmt.Errorf("wrong key type")
	}

	parentID, err := strconv.ParseInt(split[0], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("error while pirsing parentID, err: %+v", err)
	}

	sequence, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("error while pirsing sequence, err: %+v", err)
	}

	return parentID, sequence, nil
}