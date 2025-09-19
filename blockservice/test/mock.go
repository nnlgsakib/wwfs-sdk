package bstest

import (
	testinstance "github.com/nnlgsakib/wwfs-sdk/bitswap/testinstance"
	tn "github.com/nnlgsakib/wwfs-sdk/bitswap/testnet"
	"github.com/nnlgsakib/wwfs-sdk/blockservice"
	mockrouting "github.com/nnlgsakib/wwfs-sdk/routing/mock"
	delay "github.com/ipfs/go-ipfs-delay"
)

// Mocks returns |n| connected mock Blockservices
func Mocks(n int, opts ...blockservice.Option) []blockservice.BlockService {
	net := tn.VirtualNetwork(delay.Fixed(0))
	routing := mockrouting.NewServer()
	sg := testinstance.NewTestInstanceGenerator(net, routing, nil, nil)
	instances := sg.Instances(n)

	var servs []blockservice.BlockService
	for _, i := range instances {
		servs = append(servs, blockservice.New(i.Blockstore,
			i.Exchange, opts...))
	}
	return servs
}
