package mon

import (
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
)

type monNode struct {
	id       int
	peers    []string
	proposeC chan string
	node     raft.Node
}

func newMonNode(id int, peers []string) *monNode {
	m := &monNode{
		id:       id,
		peers:    peers,
		proposeC: make(chan string),
	}
	return m
}
