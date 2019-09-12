// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dpos

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/33cn/chain33/util"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/33cn/chain33/blockchain"
	"github.com/33cn/chain33/common/address"
	"github.com/33cn/chain33/common/crypto"
	"github.com/33cn/chain33/common/limits"
	"github.com/33cn/chain33/common/log"
	"github.com/33cn/chain33/executor"
	"github.com/33cn/chain33/mempool"
	"github.com/33cn/chain33/p2p"
	"github.com/33cn/chain33/queue"
	"github.com/33cn/chain33/rpc"
	"github.com/33cn/chain33/store"
	_ "github.com/33cn/chain33/system"
	cty "github.com/33cn/chain33/system/dapp/coins/types"
	"github.com/33cn/chain33/types"
	ttypes "github.com/33cn/plugin/plugin/consensus/dpos/types"
	dty "github.com/33cn/plugin/plugin/dapp/dposvote/types"
	_ "github.com/33cn/plugin/plugin/dapp/init"
	pty "github.com/33cn/plugin/plugin/dapp/norm/types"
	_ "github.com/33cn/plugin/plugin/store/init"
	"google.golang.org/grpc"
)

var (
	random    *rand.Rand
	loopCount = 10
	conn      *grpc.ClientConn
	c         types.Chain33Client
	strPubkey = "03EF0E1D3112CF571743A3318125EDE2E52A4EB904BCBAA4B1F75020C2846A7EB4"
	pubkey    []byte

	genesisKey = "CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944"
	validatorKey = "5A6A14DA6F5A42835E529D75D87CC8904544F59EEE5387A37D87EEAD194D7EB2"

	genesisAddr = "14KEKbYtKKQm4wMthSK9J4La4nAiidGozt"
	validatorAddr = "15LsTP6tkYGZcN7tc1Xo2iYifQfowxot3b"
)

const fee = 1e6

func init() {
	err := limits.SetLimits()
	if err != nil {
		panic(err)
	}
	random = rand.New(rand.NewSource(types.Now().UnixNano()))
	log.SetLogLevel("info")
	pubkey, _ = hex.DecodeString(strPubkey)
}
func TestDposPerf(t *testing.T) {
	DposPerf()
	fmt.Println("=======start clear test data!=======")
	clearTestData()
}

func DposPerf() {
	q, chain, s, mem, exec, cs, p2p := initEnvDpos()
	defer chain.Close()
	defer mem.Close()
	defer exec.Close()
	defer s.Close()
	defer q.Close()
	defer cs.Close()
	defer p2p.Close()
	err := createConn()
	for err != nil {
		err = createConn()
	}
	time.Sleep(10 * time.Second)
	for i := 0; i < loopCount; i++ {
		NormPut()
		time.Sleep(time.Second)
	}
	//从创世地址向测试地址转入代币
	sendTransferTx(genesisKey, validatorAddr, 2000000000000)
	time.Sleep(3 * time.Second)
	in := &types.ReqBalance{}
	in.Addresses = append(in.Addresses, validatorAddr)
	acct, err := c.GetBalance(context.Background(), in)
	if err != nil || len(acct.Acc) == 0 {
		fmt.Println("no balance for ", validatorAddr)
	} else {
		fmt.Println(validatorAddr, " balance:", acct.Acc[0].Balance, "frozen:", acct.Acc[0].Frozen)
	}
	//从测试地址向dos合约转入代币
	sendTransferToExecTx(validatorKey, "dpos", 1600000000000)

	time.Sleep(3 * time.Second)

	in2 := &types.ReqBalance{}
	in2.Addresses = append(in2.Addresses, validatorAddr)
	acct, err = c.GetBalance(context.Background(), in2)
	if err != nil || len(acct.Acc) == 0 {
		fmt.Println("no balance for ", validatorAddr)
	} else {
		fmt.Println(validatorAddr, " balance:", acct.Acc[0].Balance, "frozen:", acct.Acc[0].Frozen)
	}

	sendRegistCandidatorTx(strPubkey, validatorAddr, "127.0.0.1", validatorKey)

	time.Sleep(3 * time.Second)

	now := time.Now().Unix()
	task := DecideTaskByTime(now)

	dposClient := cs.(*Client)
	dposClient.csState.QueryCycleBoundaryInfo(task.Cycle)
	dposClient.csState.GetCBInfoByCircle(task.Cycle)
	dposClient.csState.QueryVrf(pubkey, task.Cycle)
	dposClient.csState.QueryVrfs(dposClient.csState.validatorMgr.Validators, task.Cycle)
	dposClient.csState.GetVrfInfoByCircle(task.Cycle, VrfQueryTypeM)
	dposClient.csState.GetVrfInfoByCircle(task.Cycle, VrfQueryTypeRP)
	dposClient.csState.GetVrfInfosByCircle(task.Cycle)
	input := []byte("data1")
	hash, proof := dposClient.csState.VrfEvaluate(input)
	if dposClient.csState.VrfProof(pubkey, input, hash, proof) {
		fmt.Println("VrfProof ok")
	}
	dposClient.QueryTopNCandidators(1)

	time.Sleep(1 * time.Second)
	info := &dty.DposCBInfo{
		Cycle:      task.Cycle,
		StopHeight: dposClient.GetCurrentHeight(),
		StopHash:   hex.EncodeToString(dposClient.GetCurrentBlock().Hash()),
		Pubkey:     strPubkey,
	}
	dposClient.csState.SendCBTx(info)
	sendCBTx(dposClient.csState, info)
	time.Sleep(2 * time.Second)
	/*
		info2 := dposClient.csState.GetCBInfoByCircle(task.Cycle)
		if info2 != nil && info2.StopHeight == info.StopHeight {
			fmt.Println("GetCBInfoByCircle ok")
		} else {
			fmt.Println("GetCBInfoByCircle failed")
		}
		time.Sleep(1 * time.Second)
	*/
	for {
		now = time.Now().Unix()
		task = DecideTaskByTime(now)
		if now < task.CycleStart+(task.CycleStop-task.CycleStart)/2 {
			break
		}
		time.Sleep(1 * time.Second)
	}

	vrfM := &dty.DposVrfMRegist{
		Pubkey: strPubkey,
		Cycle:  task.Cycle,
		M:      "data1",
	}
	if dposClient.csState.SendRegistVrfMTx(vrfM) {
		fmt.Println("SendRegistVrfMTx ok")
	} else {
		fmt.Println("SendRegistVrfMTx failed")
	}
	sendRegistVrfMTx(dposClient.csState, vrfM)

	time.Sleep(2 * time.Second)

	vrfInfo, err := dposClient.csState.QueryVrf(pubkey, task.Cycle)
	if err != nil || vrfInfo == nil {
		fmt.Println("QueryVrf failed")
	} else {
		fmt.Println("QueryVrf ok,", vrfInfo.Cycle, "|", len(vrfInfo.M))
	}

	for {
		now = time.Now().Unix()
		task = DecideTaskByTime(now)
		if now > task.CycleStart+(task.CycleStop-task.CycleStart)/2 {
			break
		}
		time.Sleep(1 * time.Second)
	}
	vrfRP := &dty.DposVrfRPRegist{
		Pubkey: strPubkey,
		Cycle:  task.Cycle,
		R:      "22a58fbbe8002939b7818184e663e6c57447f4354adba31ad3c7f556e153353c",
		P:      "5ed22d8c1cc0ad131c1c9f82daec7b99ff25ae5e717624b4a8cf60e0f3dca2c97096680cd8df0d9ed8662ce6513edf5d1676ad8d72b7e4f0e0de687bd38623f404eb085d28f5631207cf97a02c55f835bd3733241c7e068b80cf75e2afd12fd4c4cb8e6f630afa2b7b2918dff3d279e50acab59da1b25b3ff920b69c443da67320",
	}
	if dposClient.csState.SendRegistVrfRPTx(vrfRP) {
		fmt.Println("SendRegistVrfRPTx ok")
	} else {
		fmt.Println("SendRegistVrfRPTx failed")
	}

	sendRegistVrfRPTx(dposClient.csState, vrfRP)

	time.Sleep(2 * time.Second)

	vrfInfo, err = dposClient.csState.QueryVrf(pubkey, task.Cycle)
	if err != nil || vrfInfo == nil {
		fmt.Println("QueryVrf failed")
	} else {
		fmt.Println("QueryVrf ok,", vrfInfo.Cycle, "|", len(vrfInfo.M), "|", len(vrfInfo.R), "|", len(vrfInfo.P))
	}
	time.Sleep(2 * time.Second)

	var cands []*dty.Candidator
	cand := &dty.Candidator{
		Pubkey:  pubkey,
		Address: hex.EncodeToString(dposClient.csState.privValidator.GetAddress()),
		IP:      "127.0.0.1",
		Votes:   100,
		Status:  0,
	}
	cands = append(cands, cand)
	topNCand := &dty.TopNCandidator{
		Cands:        cands,
		Hash:         []byte("abafasfda"),
		Height:       dposClient.GetCurrentHeight(),
		SignerPubkey: pubkey,
	}
	reg := &dty.TopNCandidatorRegist{
		Cand: topNCand,
	}

	if dposClient.csState.SendTopNRegistTx(reg) {
		fmt.Println("SendTopNRegistTx ok")
	} else {
		fmt.Println("SendTopNRegistTx failed")
	}
	//sendTopNRegistTx(dposClient.csState, reg)

	time.Sleep(2 * time.Second)

	dposClient.QueryTopNCandidators(dposClient.GetCurrentHeight() / blockNumToUpdateDelegate)

	time.Sleep(2 * time.Second)
}

func initEnvDpos() (queue.Queue, *blockchain.BlockChain, queue.Module, queue.Module, *executor.Executor, queue.Module, queue.Module) {
	var q = queue.New("channel")
	flag.Parse()
	cfg, sub := types.InitCfg("chain33.test.toml")
	types.Init(cfg.Title, cfg)
	chain := blockchain.New(cfg.BlockChain)
	chain.SetQueueClient(q.Client())

	exec := executor.New(cfg.Exec, sub.Exec)
	exec.SetQueueClient(q.Client())
	types.SetMinFee(0)
	s := store.New(cfg.Store, sub.Store)
	s.SetQueueClient(q.Client())

	cs := New(cfg.Consensus, sub.Consensus["dpos"])
	cs.SetQueueClient(q.Client())

	mem := mempool.New(cfg.Mempool, nil)
	mem.SetQueueClient(q.Client())
	network := p2p.New(cfg.P2P)

	network.SetQueueClient(q.Client())

	rpc.InitCfg(cfg.RPC)
	gapi := rpc.NewGRpcServer(q.Client(), nil)
	go gapi.Listen()
	return q, chain, s, mem, exec, cs, network
}

func createConn() error {
	var err error
	url := "127.0.0.1:8802"
	fmt.Println("grpc url:", url)
	conn, err = grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	c = types.NewChain33Client(conn)
	//r = rand.New(rand.NewSource(types.Now().UnixNano()))
	return nil
}

func generateKey(i, valI int) string {
	key := make([]byte, valI)
	binary.PutUvarint(key[:10], uint64(valI))
	binary.PutUvarint(key[12:24], uint64(i))
	if _, err := rand.Read(key[24:]); err != nil {
		os.Exit(1)
	}
	return string(key)
}

func generateValue(i, valI int) string {
	value := make([]byte, valI)
	binary.PutUvarint(value[:16], uint64(i))
	binary.PutUvarint(value[32:128], uint64(i))
	if _, err := rand.Read(value[128:]); err != nil {
		os.Exit(1)
	}
	return string(value)
}

func getprivkey(key string) crypto.PrivKey {
	bkey, err := hex.DecodeString(key)
	if err != nil {
		panic(err)
	}
	priv, err := ttypes.ConsensusCrypto.PrivKeyFromBytes(bkey)
	if err != nil {
		panic(err)
	}
	return priv
}

func prepareTxList() *types.Transaction {
	var key string
	var value string
	var i int

	key = generateKey(i, 32)
	value = generateValue(i, 180)

	nput := &pty.NormAction_Nput{Nput: &pty.NormPut{Key: []byte(key), Value: []byte(value)}}
	action := &pty.NormAction{Value: nput, Ty: pty.NormActionPut}
	tx := &types.Transaction{Execer: []byte("norm"), Payload: types.Encode(action), Fee: fee}
	tx.To = address.ExecAddress("norm")
	tx.Nonce = random.Int63()
	tx.Sign(types.SECP256K1, getprivkey("CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944"))
	return tx
}

func clearTestData() {
	err := os.RemoveAll("datadir")
	if err != nil {
		fmt.Println("delete datadir have a err:", err.Error())
	}
	fmt.Println("test data clear successfully!")
}

func NormPut() {
	tx := prepareTxList()

	reply, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if !reply.IsOk {
		fmt.Fprintln(os.Stderr, errors.New(string(reply.GetMsg())))
		return
	}
}

// SendCBTx method
func sendCBTx(cs *ConsensusState, info *dty.DposCBInfo) bool {
	//info.Pubkey = strings.ToUpper(hex.EncodeToString(cs.privValidator.GetPubKey().Bytes()))
	canonical := dty.CanonicalOnceCBInfo{
		Cycle:      info.Cycle,
		StopHeight: info.StopHeight,
		StopHash:   info.StopHash,
		Pubkey:     info.Pubkey,
	}

	byteCB, err := json.Marshal(&canonical)
	if err != nil {
		dposlog.Error("marshal CanonicalOnceCBInfo failed", "err", err)
	}

	sig, err := cs.privValidator.SignMsg(byteCB)
	if err != nil {
		dposlog.Error("SignCBInfo failed.", "err", err)
		return false
	}

	info.Signature = hex.EncodeToString(sig.Bytes())
	tx, err := cs.client.CreateRecordCBTx(info)
	if err != nil {
		dposlog.Error("CreateRecordCBTx failed.", "err", err)
		return false
	}

	tx.Fee = fee
	cs.privValidator.SignTx(tx)
	dposlog.Info("Sign RecordCBTx ok.")

	reply, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	if !reply.IsOk {
		fmt.Fprintln(os.Stderr, errors.New(string(reply.GetMsg())))
		return false
	}

	return true
}


func sendRegistVrfMTx(cs *ConsensusState, info *dty.DposVrfMRegist) bool {
	tx, err := cs.client.CreateRegVrfMTx(info)
	if err != nil {
		dposlog.Error("CreateRegVrfMTx failed.", "err", err)
		return false
	}
	tx.Fee = fee

	cs.privValidator.SignTx(tx)
	dposlog.Info("Sign RegistVrfMTx ok.")
	reply, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	if !reply.IsOk {
		fmt.Fprintln(os.Stderr, errors.New(string(reply.GetMsg())))
		return false
	}

	return true
}

// SendRegistVrfRPTx method
func sendRegistVrfRPTx(cs *ConsensusState, info *dty.DposVrfRPRegist) bool {
	tx, err := cs.client.CreateRegVrfRPTx(info)
	if err != nil {
		dposlog.Error("CreateRegVrfRPTx failed.", "err", err)
		return false
	}

	tx.Fee = fee

	cs.privValidator.SignTx(tx)
	dposlog.Info("Sign RegVrfRPTx ok.")
	reply, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	if !reply.IsOk {
		fmt.Fprintln(os.Stderr, errors.New(string(reply.GetMsg())))
		return false
	}

	return true
}

func sendTopNRegistTx(cs *ConsensusState, reg *dty.TopNCandidatorRegist) bool {
	//info.Pubkey = strings.ToUpper(hex.EncodeToString(cs.privValidator.GetPubKey().Bytes()))
	obj := dty.CanonicalTopNCandidator(reg.Cand)
	reg.Cand.Hash = obj.ID()
	reg.Cand.SignerPubkey = cs.privValidator.GetPubKey().Bytes()

	byteCB, err := json.Marshal(reg.Cand)
	if err != nil {
		dposlog.Error("marshal TopNCandidator failed", "err", err)
	}

	sig, err := cs.privValidator.SignMsg(byteCB)
	if err != nil {
		dposlog.Error("TopNCandidator failed.", "err", err)
		return false
	}

	reg.Cand.Signature = sig.Bytes()
	tx, err := cs.client.CreateTopNRegistTx(reg)
	if err != nil {
		dposlog.Error("CreateTopNRegistTx failed.", "err", err)
		return false
	}

	tx.Fee = fee

	cs.privValidator.SignTx(tx)
	dposlog.Info("Sign TopNRegistTx ok.")
	reply, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	}
	if !reply.IsOk {
		fmt.Fprintln(os.Stderr, errors.New(string(reply.GetMsg())))
		return false
	}

	return true
}

func sendTransferTx(fromKey ,to string, amount int64) bool {
	signer := util.HexToPrivkey(fromKey)
	var tx *types.Transaction
	transfer := &cty.CoinsAction{}
	v := &cty.CoinsAction_Transfer{Transfer: &types.AssetsTransfer{Amount: amount, Note: []byte(""), To: to}}
	transfer.Value = v
	transfer.Ty = cty.CoinsActionTransfer
	execer := []byte("coins")
	tx = &types.Transaction{Execer: execer, Payload: types.Encode(transfer), To: to, Fee: fee}
	tx, err := types.FormatTx(string(execer), tx)
	if err != nil {
		fmt.Println("in sendTransferTx formatTx failed")
		return false
	}

	tx.Sign(types.SECP256K1, signer)
	reply, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Println("in sendTransferTx SendTransaction failed")

		return false
	}
	if !reply.IsOk {
		fmt.Fprintln(os.Stderr, errors.New(string(reply.GetMsg())))
		fmt.Println("in sendTransferTx SendTransaction failed,reply not ok.")

		return false
	}
	fmt.Println("sendTransferTx ok")

	return true
}

func sendTransferToExecTx(fromKey ,execName string, amount int64) bool {
	signer := util.HexToPrivkey(fromKey)
	var tx *types.Transaction
	transfer := &cty.CoinsAction{}
	execAddr := address.ExecAddress(execName)
	v := &cty.CoinsAction_TransferToExec{TransferToExec: &types.AssetsTransferToExec{Amount: amount, Note: []byte(""), ExecName: execName, To: execAddr}}
	transfer.Value = v
	transfer.Ty = cty.CoinsActionTransferToExec
	execer := []byte("coins")
	tx = &types.Transaction{Execer: execer, Payload: types.Encode(transfer), To: address.ExecAddress("dpos"), Fee: fee}
	tx, err := types.FormatTx(string(execer), tx)
	if err != nil {
		fmt.Println("sendTransferToExecTx formatTx failed.")

		return false
	}

	tx.Sign(types.SECP256K1, signer)
	reply, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Println("in sendTransferToExecTx SendTransaction failed")

		return false
	}
	if !reply.IsOk {
		fmt.Fprintln(os.Stderr, errors.New(string(reply.GetMsg())))
		fmt.Println("in sendTransferToExecTx SendTransaction failed,reply not ok.")

		return false
	}

	fmt.Println("sendTransferToExecTx ok")

	return true
}


func sendRegistCandidatorTx(ppubkey, addr, ip, privKey string) bool {
	signer := util.HexToPrivkey(privKey)
	var tx *types.Transaction
	action := &dty.DposVoteAction{}

	v := &dty.DposVoteAction_Regist{
		Regist: &dty.DposCandidatorRegist{
			Pubkey: ppubkey,
			Address: addr,
			IP: ip,
		},
	}

	action.Value = v
	action.Ty = dty.DposVoteActionRegist
	execer := []byte("dpos")
	tx = &types.Transaction{Execer: execer, Payload: types.Encode(action), To: address.ExecAddress(string(execer)), Fee: fee}
	tx, err := types.FormatTx(string(execer), tx)
	if err != nil {
		fmt.Println("sendRegistCandidatorTx formatTx failed.")

		return false
	}

	tx.Sign(types.SECP256K1, signer)
	reply, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Println("in sendRegistCandidatorTx SendTransaction failed")

		return false
	}
	if !reply.IsOk {
		fmt.Fprintln(os.Stderr, errors.New(string(reply.GetMsg())))
		fmt.Println("in sendTransferToExecTx SendTransaction failed,reply not ok.")

		return false
	}

	fmt.Println("sendRegistCandidatorTx ok")

	return true
}
