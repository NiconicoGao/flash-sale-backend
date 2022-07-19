package utils

import "github.com/bwmarrin/snowflake"

var node *snowflake.Node

func GenerateID() int64 {
	return node.Generate().Int64()
}

func InitSnowflake() (err error) {
	node, err = snowflake.NewNode(1)
	if err != nil {
		return err
	}

	return nil
}
