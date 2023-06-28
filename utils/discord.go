package utils

import (
	"strconv"
)

// var sfnode *discord.Node
// var dcsfOnce sync.Once

// func GetDcSnowflake() *discord.Node {
// 	dcsfOnce.Do(func() {
// 		sfnode, _ = discord.NewNode()
// 	})
// 	return sfnode
// }

// DCNonce returns a snowflake string that can be used as a nonce for Discord.
func DCNonce() string {
	return strconv.Itoa(int(NextSnowflakeID()))
}
