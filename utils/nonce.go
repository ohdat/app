package utils

import (
	"fmt"
	"strings"
)

const (
	LfgNonceLayout = "lfg:%d"
)

func GetLfgNonce() string {
	return fmt.Sprintf(LfgNonceLayout, NextSnowflakeID())
}

func LfgNonce() []string {
	return strings.Split(GetLfgNonce(), ":")
}
