package misc

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
)
const CHECPOINTER_CONTRACT_ADDRESS = "0xE890f129a8C21c19938DFa9EBCDB80f0b5033921"
// type Casdhekcpointer struct {
// 	ethapi ethapi.PublicBlockChainAPI
// 	abi abi.ABI
// }
// // VerifyForkHashes verifies that blocks conforming to network hard-forks do have
// // the correct hashes, to avoid clients going off on different chains. This is an
// // optional feature.
// func (c *Casdhekcpointer) getCheckpoint(blockhash common.Hash) ([]common.Address, error) {
	
// 	method := "attest"
// 	ctx, cancel :=	context.WithCancel(context.Background())
// 	defer cancel() //
// 	data, err := c.abi.Pack(method)
// 	if err != nil {
// 		return nil, err
// 	}

// 	msgData := (hexutil.Bytes)(data)
// 	toAddress :=common.HexToAddress(CHECPOINTER_CONTRACT_ADDRESS)
// 	gas := (hexutil.Uint64)(uint64(math.MaxUint64 / 2))
// 	// res, err := c.ethapi.Call(ctx, ethapi.CallArgs{
// 	// 	Gas: &gas,
// 	// 	To: &toAddress,
// 	// 	Data: &msgData,
// 	// }, rpc.BlockNumberOrHashWithHash(blockhash, false), nil)

// 	if err != nil {
// 		return nil, err
// 	}

// 	out := new([]common.Address)

// 	err = c.abi.UnpackIntoInterface(out, method, res)
// 	if err != nil {
// 		return nil, err
// 	}




// 	if err != nil {
// 		return nil, err
// 	}

// 	return nil, fmt.Errorf("not implemented")
// }
func Check(header *types.Header) error {
	
	
	fmt.Println("Pizda")
	// All ok, return
	return nil
}
