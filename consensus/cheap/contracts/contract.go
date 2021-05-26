package contracts

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
)

type Contract struct {
	Address common.Address
	Abi string
	Code string
}
// TODO: make this read soidity files and populate abi and code
var Contracts map[string]Contract = map[string]Contract {
	"Dummy" : {
		Address: common.HexToAddress("0x1337000000000000000000000000000000000000"),
		Abi: `[
			{
				"inputs": [],
				"stateMutability": "nonpayable",
				"type": "constructor"
			},
			{
				"inputs": [],
				"name": "A",
				"outputs": [
					{
						"internalType": "uint64",
						"name": "",
						"type": "uint64"
					}
				],
				"stateMutability": "pure",
				"type": "function"
			}
		]`,
		Code: 
			"6080604052348015600f57600080fd5b506004361060285760003560e01c8063f446c1d014602d575b600080fd5b60336047565b604051603e9190605e565b60405180910390f35b6000611337905090565b6058816077565b82525050565b6000602082019050607160008301846051565b92915050565b600067ffffffffffffffff8216905091905056fea26469706673582212208eb084396988ccea5e49eca052fd74fe9f102653d5e0ddebdcfb2138428db45b64736f6c63430007060033",
		},
		"Checkpointer" : {
			Address: common.HexToAddress("0x1111000000000000000000000000000000000000"),
			Abi: `[
				{
					"inputs": [
						{
							"internalType": "uint256",
							"name": "number",
							"type": "uint256"
						},
						{
							"internalType": "bytes32",
							"name": "blockHash",
							"type": "bytes32"
						}
					],
					"name": "attest",
					"outputs": [],
					"stateMutability": "nonpayable",
					"type": "function"
				},
				{
					"inputs": [
						{
							"internalType": "address",
							"name": "toTrust",
							"type": "address"
						}
					],
					"name": "trust",
					"outputs": [],
					"stateMutability": "nonpayable",
					"type": "function"
				},
				{
					"inputs": [],
					"stateMutability": "nonpayable",
					"type": "constructor"
				},
				{
					"inputs": [
						{
							"internalType": "uint256",
							"name": "blockNumber",
							"type": "uint256"
						},
						{
							"internalType": "address",
							"name": "verifier",
							"type": "address"
						}
					],
					"name": "getBlockByNumber",
					"outputs": [
						{
							"components": [
								{
									"internalType": "bytes32",
									"name": "hash",
									"type": "bytes32"
								},
								{
									"internalType": "uint256",
									"name": "savedBlockNumber",
									"type": "uint256"
								}
							],
							"internalType": "struct Checkpointer.Checkpoint",
							"name": "",
							"type": "tuple"
						}
					],
					"stateMutability": "view",
					"type": "function"
				},
				{
					"inputs": [],
					"name": "getTrusted",
					"outputs": [
						{
							"internalType": "address[]",
							"name": "",
							"type": "address[]"
						}
					],
					"stateMutability": "view",
					"type": "function"
				}
			]`,
			Code: "608060405234801561001057600080fd5b506004361061004c5760003560e01c806302165185146100515780630a28b1081461008157806317e859261461009f5780634637d827146100bb575b600080fd5b61006b6004803603810190610066919061088f565b6100d7565b6040516100789190610b07565b60405180910390f35b610089610155565b6040516100969190610a65565b60405180910390f35b6100b960048036038101906100b491906108cb565b610331565b005b6100d560048036038101906100d09190610866565b6104da565b005b6100df61080a565b60008084815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060405180604001604052908160008201548152602001600182015481525050905092915050565b6060600060028054905067ffffffffffffffff81111561019e577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156101cc5781602001602082028036833780820191505090505b50905060005b60028054905081101561032957600160006002838154811061021d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168282815181106102dc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050808061032190610be6565b9150506101d2565b508091505090565b606461ffff16431015610379576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037090610a87565b60405180910390fd5b606461ffff164361038a9190610b6c565b8211156103cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c390610ae7565b60405180910390fd5b6000824090506000801b811415610418576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040f90610ac7565b60405180910390fd5b81811461045a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045190610aa7565b60405180910390fd5b60405180604001604052808381526020018481525060008085815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000015560208201518160010155905050505050565b80600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606461ffff1660028054905010156105d2576002339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610807565b6000806002600081548110610610577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163190506000600190505b6002805490508110156107805781600282815481106106a3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1631101561076d5780925060028181548110610728577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163191505b808061077890610be6565b91505061065b565b5033600283815481106107bc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505b50565b604051806040016040528060008019168152602001600081525090565b60008135905061083681610d28565b92915050565b60008135905061084b81610d3f565b92915050565b60008135905061086081610d56565b92915050565b60006020828403121561087857600080fd5b600061088684828501610827565b91505092915050565b600080604083850312156108a257600080fd5b60006108b085828601610851565b92505060206108c185828601610827565b9150509250929050565b600080604083850312156108de57600080fd5b60006108ec85828601610851565b92505060206108fd8582860161083c565b9150509250929050565b6000610913838361091f565b60208301905092915050565b61092881610ba0565b82525050565b600061093982610b32565b6109438185610b4a565b935061094e83610b22565b8060005b8381101561097f5781516109668882610907565b975061097183610b3d565b925050600181019050610952565b5085935050505092915050565b61099581610bb2565b82525050565b60006109a8601083610b5b565b91506109b382610c5e565b602082019050919050565b60006109cb602383610b5b565b91506109d682610c87565b604082019050919050565b60006109ee601383610b5b565b91506109f982610cd6565b602082019050919050565b6000610a11601383610b5b565b9150610a1c82610cff565b602082019050919050565b604082016000820151610a3d600085018261098c565b506020820151610a506020850182610a56565b50505050565b610a5f81610bdc565b82525050565b60006020820190508181036000830152610a7f818461092e565b905092915050565b60006020820190508181036000830152610aa08161099b565b9050919050565b60006020820190508181036000830152610ac0816109be565b9050919050565b60006020820190508181036000830152610ae0816109e1565b9050919050565b60006020820190508181036000830152610b0081610a04565b9050919050565b6000604082019050610b1c6000830184610a27565b92915050565b6000819050602082019050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b6000610b7782610bdc565b9150610b8283610bdc565b925082821015610b9557610b94610c2f565b5b828203905092915050565b6000610bab82610bbc565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000610bf182610bdc565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610c2457610c23610c2f565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f636861696e20697320746f6f206e657700000000000000000000000000000000600082015250565b7f626c6f636b6861736820646f65736e2774206d6174636820617474657374617460008201527f696f6e0000000000000000000000000000000000000000000000000000000000602082015250565b7f626c6f636b68617368206e6f7420666f756e6400000000000000000000000000600082015250565b7f626c6f636b20697320746f6f20726563656e7400000000000000000000000000600082015250565b610d3181610ba0565b8114610d3c57600080fd5b50565b610d4881610bb2565b8114610d5357600080fd5b50565b610d5f81610bdc565b8114610d6a57600080fd5b5056fea2646970667358221220a89d56318ed129150ca361d9dc64efb8547f5299e51a49fcef1656c571b8769e64736f6c63430008010033",
		},
}


func Deploy(state *state.StateDB) {
	for n, c := range Contracts {
		code, err := hex.DecodeString(c.Code)
		if err != nil {
			panic(fmt.Sprintf("Bad code for contract %s\n", n))
		}

		state.SetCode(c.Address, code)
		// TODO: find a way to check if code was set correctly
	}
}