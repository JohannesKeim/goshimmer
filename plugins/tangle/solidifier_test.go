package tangle

import (
	"os"
	"sync"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/iotaledger/goshimmer/packages/gossip"
	pb "github.com/iotaledger/goshimmer/packages/gossip/proto"
	"github.com/iotaledger/goshimmer/packages/model/meta_transaction"
	"github.com/iotaledger/goshimmer/packages/model/value_transaction"
	"github.com/iotaledger/hive.go/events"
	"github.com/iotaledger/hive.go/node"
	"github.com/iotaledger/hive.go/parameter"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	parameter.FetchConfig(false)
	os.Exit(m.Run())
}

func TestSolidifier(t *testing.T) {
	// show all error messages for tests
	// TODO: adjust logger package

	// start a test node
	node.Start(PLUGIN)

	// create transactions and chain them together
	transaction1 := value_transaction.New()
	transaction1.SetValue(1)
	require.NoError(t, transaction1.DoProofOfWork(meta_transaction.MIN_WEIGHT_MAGNITUDE))

	transaction2 := value_transaction.New()
	transaction2.SetValue(2)
	transaction2.SetBranchTransactionHash(transaction1.GetHash())
	require.NoError(t, transaction2.DoProofOfWork(meta_transaction.MIN_WEIGHT_MAGNITUDE))

	transaction3 := value_transaction.New()
	transaction3.SetValue(3)
	transaction3.SetBranchTransactionHash(transaction2.GetHash())
	require.NoError(t, transaction3.DoProofOfWork(meta_transaction.MIN_WEIGHT_MAGNITUDE))

	transaction4 := value_transaction.New()
	transaction4.SetValue(4)
	transaction4.SetBranchTransactionHash(transaction3.GetHash())
	require.NoError(t, transaction4.DoProofOfWork(meta_transaction.MIN_WEIGHT_MAGNITUDE))

	// setup event handlers
	var wg sync.WaitGroup
	Events.TransactionSolid.Attach(events.NewClosure(func(transaction *value_transaction.ValueTransaction) {
		t.Log("Tx solidified", transaction.GetValue())
		wg.Done()
	}))

	gossip.Events.RequestTransaction.Attach(events.NewClosure(func(ev *gossip.RequestTransactionEvent) {
		tx := &pb.Transaction{Body: transaction3.MetaTransaction.GetBytes()}
		b, _ := proto.Marshal(tx)
		gossip.Events.TransactionReceived.Trigger(&gossip.TransactionReceivedEvent{Body: b})
	}))

	// issue transactions
	wg.Add(4)
	tx := &pb.Transaction{Body: transaction1.MetaTransaction.GetBytes()}
	b, _ := proto.Marshal(tx)
	gossip.Events.TransactionReceived.Trigger(&gossip.TransactionReceivedEvent{Body: b})

	tx = &pb.Transaction{Body: transaction2.MetaTransaction.GetBytes()}
	b, _ = proto.Marshal(tx)
	gossip.Events.TransactionReceived.Trigger(&gossip.TransactionReceivedEvent{Body: b})

	// tx = &pb.Transaction{Body: transaction3.MetaTransaction.GetBytes()}
	// b, _ = proto.Marshal(tx)
	// gossip.Events.TransactionReceived.Trigger(&gossip.TransactionReceivedEvent{Body: b})

	tx = &pb.Transaction{Body: transaction4.MetaTransaction.GetBytes()}
	b, _ = proto.Marshal(tx)
	gossip.Events.TransactionReceived.Trigger(&gossip.TransactionReceivedEvent{Body: b})

	// wait until all are solid
	wg.Wait()

	// shutdown test node
	node.Shutdown()

}
