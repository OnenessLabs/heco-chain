package systemcontract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"math"
	"math/big"
)

var (
	//TODO: set these to formal addresses
	admin        = common.HexToAddress("0x000000000000000000000000000000000000f004")
	adminTestnet = common.HexToAddress("0x000000000000000000000000000000000000f004")
)

const (
	code = "0x608060405234801561001057600080fd5b50600436106101215760003560e01c8063a29351fb116100ad578063e08b1d3811610071578063e08b1d38146103cf578063efd8d8e2146103f0578063f851a440146103f8578063fb48270c14610400578063fbb847e11461040857610121565b8063a29351fb146102bb578063be6456921461034b578063c4d66de814610365578063c967f90f1461038b578063db78dd28146103aa57610121565b806326782247116100f457806326782247146102605780633656de21146102685780633a061bd3146102855780634fb9e9b71461028d5780636233be5d146102b357610121565b806305b8481014610126578063158ef93e146102015780631b5e358c1461021d578063232e5ffc14610241575b600080fd5b6101496004803603602081101561013c57600080fd5b503563ffffffff16610410565b60405180868152602001856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101c25781810151838201526020016101aa565b50505050905090810190601f1680156101ef5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b610209610597565b604080519115158252519081900360200190f35b6102256105a0565b604080516001600160a01b039092168252519081900360200190f35b61025e6004803603602081101561025757600080fd5b50356105a6565b005b61022561076e565b6101496004803603602081101561027e57600080fd5b503561077d565b6102256107e7565b61025e600480360360208110156102a357600080fd5b50356001600160a01b03166107ed565b610225610888565b61025e600480360360808110156102d157600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561030c57600080fd5b82018360208201111561031e57600080fd5b8035906020019184600183028401116401000000008311171561034057600080fd5b50909250905061088e565b610353610bc5565b60408051918252519081900360200190f35b61025e6004803603602081101561037b57600080fd5b50356001600160a01b0316610bd2565b610393610c4f565b6040805161ffff9092168252519081900360200190f35b6103b2610c54565b6040805167ffffffffffffffff9092168252519081900360200190f35b6103d7610c5b565b6040805163ffffffff9092168252519081900360200190f35b6103b2610c62565b610225610c68565b61025e610c7c565b610353610d36565b60008060008060606003805490508663ffffffff161061046c576040805162461bcd60e51b8152602060048201526012602482015271496e646578206f7574206f662072616e676560701b604482015290519081900360640190fd5b610474610d3c565b60038763ffffffff168154811061048757fe5b60009182526020918290206040805160a081018252600593909302909101805483526001808201546001600160a01b0390811685870152600280840154909116858501526003830154606086015260048301805485516101009482161594909402600019011691909104601f810187900487028301870190945283825293949193608086019391929091908301828280156105635780601f1061053857610100808354040283529160200191610563565b820191906000526020600020905b81548152906001019060200180831161054657829003601f168201915b5050509190925250508151602083015160408401516060850151608090950151929c919b5099509297509550909350505050565b60005460ff1681565b61f00181565b3341146105e7576040805162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b604482015290519081900360640190fd5b60005b60035481101561076a57816003828154811061060257fe5b9060005260206000209060050201600001541415610762576003546000190181146106d65760038054600019810190811061063957fe5b90600052602060002090600502016003828154811061065457fe5b60009182526020909120825460059092020190815560018083015481830180546001600160a01b03199081166001600160a01b03938416179091556002808601548186018054909316931692909217905560038085015490840155600480850180546106d29492860193919281161561010002600019011604610d7d565b5050505b60038054806106e157fe5b6000828152602081206005600019909301928302018181556001810180546001600160a01b03199081169091556002820180549091169055600381018290559061072e6004830182610e02565b5050905560405182907fc2946e69de813a7cede502a3b315aa221abf9fcca5c7134b0ae6b2c3857cf63d90600090a261076a565b6001016105ea565b5050565b6001546001600160a01b031681565b600080600080606060028054905086106107d2576040805162461bcd60e51b8152602060048201526011602482015270125908191bd95cc81b9bdd08195e1a5cdd607a1b604482015290519081900360640190fd5b6107da610d3c565b6002878154811061048757fe5b61f00081565b60005461010090046001600160a01b0316331461083e576040805162461bcd60e51b815260206004820152600a60248201526941646d696e206f6e6c7960b01b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040517faefcaa6215f99fe8c2f605dd268ee4d23a5b596bbca026e25ce8446187f4f1ba90600090a250565b61f00281565b60005461010090046001600160a01b031633146108df576040805162461bcd60e51b815260206004820152600a60248201526941646d696e206f6e6c7960b01b604482015290519081900360640190fd5b6002546108ea610d3c565b6040518060a00160405280838152602001886001600160a01b03168152602001876001600160a01b0316815260200186815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250939094525050600280546001810182559152825160059091027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace81019182556020808501517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf830180546001600160a01b039283166001600160a01b03199182161790915560408701517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad08501805491909316911617905560608501517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad1830155608085015180519596508695939450610a72937f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad290930192910190610e49565b505060038054600181018255600091909152825160059091027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b81019182556020808501517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c830180546001600160a01b039283166001600160a01b03199182161790915560408701517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85d8501805491909316911617905560608501517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85e83015560808501518051869550610b8e937fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85f01929190910190610e49565b50506040518391507f14ca27cd9911371c77ed1cf3cee0a4320613b07478668958754713b3c880cd0f90600090a250505050505050565b6801bc16d674ec80000081565b60005460ff1615610c20576040805162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b604482015290519081900360640190fd5b6000805460ff196001600160a01b0390931661010002610100600160a81b031990911617919091166001179055565b601581565b6201518081565b6003545b90565b61708081565b60005461010090046001600160a01b031681565b6001546001600160a01b03163314610ccc576040805162461bcd60e51b815260206004820152600e60248201526d4e65772061646d696e206f6e6c7960901b604482015290519081900360640190fd5b60018054600080546001600160a01b03808416610100908102610100600160a81b0319909316929092178084556001600160a01b03199094169094556040519204909216917f7ce7ec0b50378fb6c0186ffb5f48325f6593fcb4ca4386f21861af3129188f5c91a2565b60025490565b6040518060a001604052806000815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160008152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610db65780548555610df2565b82800160010185558215610df257600052602060002091601f016020900482015b82811115610df2578254825591600101919060010190610dd7565b50610dfe929150610eb7565b5090565b50805460018160011615610100020316600290046000825580601f10610e285750610e46565b601f016020900490600052602060002090810190610e469190610eb7565b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e8a57805160ff1916838001178555610df2565b82800160010185558215610df2579182015b82811115610df2578251825591602001919060010190610e9c565b610c5f91905b80821115610dfe5760008155600101610ebd56fea26469706673582212201992573d2c4df9dc7ba9f6222b00c2ba3a03f70415ce46192c392f5ef91b968c64736f6c63430006010033"
)

type hardForkSysGov struct {
}

func (s *hardForkSysGov) GetName() string {
	return "SysGov"
}

func (s *hardForkSysGov) Update(config *params.ChainConfig, height *big.Int, state *state.StateDB) (err error) {
	contractCode := common.FromHex(code)

	//write code to sys contract
	state.SetCode(SysGovContractAddr, contractCode)
	log.Debug("Write code to system contract account", "addr", SysGovContractAddr.String(), "code", code)

	return
}

func (s *hardForkSysGov) getAdminByChainId(chainId *big.Int) common.Address {
	if chainId.Cmp(params.MainnetChainConfig.ChainID) == 0 {
		return admin
	}

	return adminTestnet
}

func (s *hardForkSysGov) Execute(state *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (err error) {

	method := "initialize"
	data, err := GetInteractiveABI()[SysGovContractName].Pack(method, s.getAdminByChainId(config.ChainID))
	if err != nil {
		log.Error("Can't pack data for initialize", "error", err)
		return err
	}

	msg := types.NewMessage(header.Coinbase, &SysGovContractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, false)

	context := core.NewEVMContext(msg, header, chainContext, nil)
	evm := vm.NewEVM(context, state, config, vm.Config{})

	_, _, err = evm.Call(vm.AccountRef(msg.From()), *msg.To(), msg.Data(), msg.Gas(), msg.Value())

	return
}
