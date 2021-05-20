package checkpoints

import (
	// 	"context"
	"fmt"
	// 	"strings"

	// 	"github.com/ethereum/go-ethereum/accounts/abi"
	// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// 	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/internal/ethapi"
	// 	"github.com/ethereum/go-ethereum/eth"
	// 	"github.com/ethereum/go-ethereum/internal/ethapi"
)
const CHECPOINTER_CONTRACT_ADDRESS = "0xE890f129a8C21c19938DFa9EBCDB80f0b5033921"
// type Casdhekcpointer struct {
// 	ethapi ethapi.PublicBlockChainAPI
// 	abi abi.ABI
// }
// // VerifyForkHashes verifies that blocks conforming to network hard-forks do have
// // the correct hashes, to avoid clients going off on different chains. This is an
// // optional feature.

const ABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"
const BYTECODE = "608060405234801561001057600080fd5b506101b9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80636ed28ed01461003b5780638f88708b14610057575b600080fd5b610055600480360381019061005091906100fc565b610087565b005b610071600480360381019061006c91906100d3565b6100a2565b60405161007e9190610147565b60405180910390f35b80600080848152602001908152602001600020819055505050565b6000806000838152602001908152602001600020549050919050565b6000813590506100cd8161016c565b92915050565b6000602082840312156100e557600080fd5b60006100f3848285016100be565b91505092915050565b6000806040838503121561010f57600080fd5b600061011d858286016100be565b925050602061012e858286016100be565b9150509250929050565b61014181610162565b82525050565b600060208201905061015c6000830184610138565b92915050565b6000819050919050565b61017581610162565b811461018057600080fd5b5056fea26469706673582212204653b43f30f0487d35476d7b81653cfec47c6134580aee5387f772e40677976864736f6c63430007060033"

var backend *ethapi.Backend = nil

func SetBackend(b *ethapi.Backend) {
	backend = b
	fmt.Printf("\n\nbackend: %p\n\n", backend)
}
// func Deploy() (error){
// 	parsed, err := abi.JSON(strings.NewReader(ABI))

// 	if err != nil {
// 		return err
// 	}

// 	addr, tx, bound, err := bind.DeployContract(
		
// 	)
// }

func x() (error) {
	// parsed, err := abi.JSON(strings.NewReader(ABI))

	// if err != nil {
	// 	return err
	// }
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// // ethapi.DoCall(
	// // 	ctx,
		
	// // )

	return nil
}
// func GetCheckpoint(blockhash common.Hash) ([]common.Address, error) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()
// 	ethapi.DoCall(
// 		ctx,

// 	)
// 	return nil, fmt.Errorf("not implemented")
// }
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
