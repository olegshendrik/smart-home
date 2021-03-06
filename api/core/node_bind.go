package core

import (
	"github.com/e154/smart-home/api/models"
	"github.com/e154/smart-home/api/common"
)

// Javascript Binding
//
// node
//	 .send()
//	 .name()
//	 .description()
//	 .ip()
//	 .status()
//
type NodeBind struct {
	node *models.Node
}

func (n *NodeBind) Send(protocol string, device *DeviceBind, return_result bool, command []byte) common.Result {
	return n.node.Send(protocol, device.model, return_result, command)
}

func (n *NodeBind) Name() string {
	return n.node.Name
}

func (n *NodeBind) Ip() string {
	return n.node.Ip
}

func (n *NodeBind) Status() string {
	return n.node.Status
}

func (n *NodeBind) Description() string {
	return n.node.Description
}