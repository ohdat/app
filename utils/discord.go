package utils

import (
	"strconv"
)

// DCNonce returns a snowflake string that can be used as a nonce for Discord.
func DCNonce() string {
	return strconv.Itoa(int(NextSnowflakeID()))
}
