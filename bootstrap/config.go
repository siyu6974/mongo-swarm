package main

import (
	"strings"

	"github.com/pkg/errors"
)

type Config struct {
	DataSet   string
	ConfigSet string
	Mongos    string
	Retry     int
	Wait      int
	Port      int
}

type ShardReplicaSet struct {
	replSetName string
	members     []string
}

// SharedReplicaSet format <replicaSetName1>/data1:27017,data2:27017,data3:27017;<replicaSetName2>/data1:27017,data2:27017,data3:27017...
func ParseShards(definition string) ([]string, error) {
	shards := strings.Split(definition, ";")

	// if len(members) < 3 {
	// 	return nil, errors.New("Invalid ReplicaSet definition, a minimum of 3 members required")
	// }

	return shards, nil
}

// SharedReplicaSet format <replicaSetName1>/data1:27017,data2:27017,data3:27017;<replicaSetName2>/data1:27017,data2:27017,data3:27017...
func ParseShardReplicaSet(definition string) ([]ShardReplicaSet, error) {
	var r []ShardReplicaSet
	shards := strings.Split(definition, ";")
	for _, mongos := range shards {
		parts := strings.Split(mongos, "/")
		replSetName := parts[0]
		members := strings.Split(parts[1], ",")
		// if len(members) < 3 {
		// 	return nil, errors.New("Invalid ReplicaSet definition, a minimum of 3 members required")
		// }
		r = append(r, ShardReplicaSet{replSetName: replSetName, members: members})
	}

	// if len(members) < 3 {
	// 	return nil, errors.New("Invalid ReplicaSet definition, a minimum of 3 members required")
	// }

	return r, nil
}

// ReplicaSet format <replicaSetName>/data1:27017,data2:27017,data3:27017
func ParseReplicaSet(definition string) (string, []string, error) {

	parts := strings.Split(definition, "/")
	if len(parts) != 2 {
		return "", nil, errors.New("Invalid ReplicaSet definition, expected <replicaSetName>/data1:27017,data2:27017,data3:27017")
	}

	replSetName := parts[0]

	members := strings.Split(parts[1], ",")

	// if len(members) < 3 {
	// 	return "", nil, errors.New("Invalid ReplicaSet definition, a minimum of 3 members required")
	// }

	return replSetName, members, nil
}

func ParseMongos(definition string) ([]string, error) {
	list := strings.Split(definition, ",")

	for _, mongos := range list {
		parts := strings.Split(mongos, ":")
		if len(parts) != 2 {
			return nil, errors.Errorf("%v invalid format, expected <HOST>:<PORT>", mongos)
		}
	}

	return list, nil
}
