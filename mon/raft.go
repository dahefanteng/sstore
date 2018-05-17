package mon

import (
	"github.com/coreos/etcd/raft"
	//"github.com/coreos/etcd/raft/raftpb"
	"context"
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

func StartRaft() {
	//receive propose from state machine
	go func() {
		for rc.proposeC != nil {
			select {
			case prop, ok := <-rc.proposeC:
				if !ok {
					rc.proposeC = nil
				} else {
					rc.node.Propose(context.TODO(), []byte(prop))
				}
			}
		}
	}()

	//store raftlog and publish to state machine
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				rc.node.Tick()
			case rd := <-rc.node.Ready():
				fmt.Println(rd)
			case err := <-rc.transport.ErrorC:
				fmt.Println(err)
			}
		}
	}()
}
