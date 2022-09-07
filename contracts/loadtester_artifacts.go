package contracts

import (
	"encoding/hex"
	"fmt"

	"github.com/umbracle/ethgo/abi"
)

var abiLoadTester *abi.ABI

// LoadTesterAbi returns the abi of the LoadTester contract
func LoadTesterAbi() *abi.ABI {
	return abiLoadTester
}

var binLoadTester []byte

// LoadTesterBin returns the bin of the LoadTester contract
func LoadTesterBin() []byte {
	return binLoadTester
}

func init() {
	var err error
	abiLoadTester, err = abi.NewABI(abiLoadTesterStr)
	if err != nil {
		panic(fmt.Errorf("cannot parse LoadTester abi: %v", err))
	}
	if len(binLoadTesterStr) != 0 {
		binLoadTester, err = hex.DecodeString(binLoadTesterStr[2:])
		if err != nil {
			panic(fmt.Errorf("cannot parse LoadTester bin: %v", err))
		}
	}
}

var binLoadTesterStr = "608060405234801561001057600080fd5b5061204a806100206000396000f3fe608060405234801561001057600080fd5b50600436106103c55760003560e01c80637c191d20116101ff578063c360aba61161011a578063dd9bef60116100ad578063f279ca811161007c578063f279ca8114610eb6578063f4d1fc6114610ee6578063f58fc36a14610f16578063fde7721c14610f46576103c5565b8063dd9bef6014610df6578063de97a36314610e26578063e9f9b3f214610e56578063ea5141e614610e86576103c5565b8063d117320b116100e9578063d117320b14610d36578063d51e7b5b14610d66578063d53ff3fd14610d96578063d93cd55814610dc6576103c5565b8063c360aba614610c76578063c420eb6114610ca6578063c4bd65d514610cd6578063ce3cf4ef14610d06576103c5565b8063a60a108711610192578063b7b8620711610161578063b7b8620714610bb6578063b81c148414610be6578063bdc875fc14610c16578063bf529ca114610c46576103c5565b8063a60a108714610af6578063a645c9c214610b26578063acaebdf614610b56578063b3d847f214610b86576103c5565b8063918a5fcd116101ce578063918a5fcd14610a3657806391e7b27714610a6657806398456f3e14610a965780639a2b7c8114610ac6576103c5565b80637c191d20146109765780637de8c6f8146109a657806380947f80146109d6578063880eff3914610a06576103c5565b80632871ef85116102ef5780634a61af1f116102825780636e7f1fe7116102515780636e7f1fe7146108b65780636f099c8d146108e657806371d91d28146109165780637b6e0b0e14610946576103c5565b80634a61af1f146107f65780634d2c74b3146108265780635590c2d91461085657806360e13cde14610886576103c5565b80633a411f12116102be5780633a411f12146107365780633a425dfc1461076657806340fe26621461079657806344cf3bc7146107c6576103c5565b80632871ef85146106885780632b21ef44146106b85780632d34e798146106e8578063371303c014610718576103c5565b806316582150116103675780631de2f343116103365780631de2f343146105c85780632007332e146105f8578063219cddeb146106285780632294fc7f14610658576103c5565b8063165821501461050857806318093b461461053857806319b621d6146105685780631aba07ea14610598576103c5565b80630ba8a73b116103a35780630ba8a73b1461045a5780631287a68c1461048a578063135d52f7146104a85780631581cf19146104d8576103c5565b8063034aef71146103ca578063050082f8146103fa578063087b4e841461042a575b600080fd5b6103e460048036038101906103df9190611f38565b610f76565b6040516103f19190611f74565b60405180910390f35b610414600480360381019061040f9190611f38565b610fb1565b6040516104219190611f74565b60405180910390f35b610444600480360381019061043f9190611f38565b610fec565b6040516104519190611f74565b60405180910390f35b610474600480360381019061046f9190611f38565b611026565b6040516104819190611f74565b60405180910390f35b610492611062565b60405161049f9190611f74565b60405180910390f35b6104c260048036038101906104bd9190611f38565b61106b565b6040516104cf9190611f74565b60405180910390f35b6104f260048036038101906104ed9190611f38565b6110a7565b6040516104ff9190611f74565b60405180910390f35b610522600480360381019061051d9190611f38565b6110e2565b60405161052f9190611f74565b60405180910390f35b610552600480360381019061054d9190611f38565b61113d565b60405161055f9190611f74565b60405180910390f35b610582600480360381019061057d9190611f38565b61117b565b60405161058f9190611f74565b60405180910390f35b6105b260048036038101906105ad9190611f38565b61120a565b6040516105bf9190611f74565b60405180910390f35b6105e260048036038101906105dd9190611f38565b611250565b6040516105ef9190611f74565b60405180910390f35b610612600480360381019061060d9190611f38565b61128e565b60405161061f9190611f74565b60405180910390f35b610642600480360381019061063d9190611f38565b6112ca565b60405161064f9190611f74565b60405180910390f35b610672600480360381019061066d9190611f38565b611305565b60405161067f9190611f74565b60405180910390f35b6106a2600480360381019061069d9190611f38565b611344565b6040516106af9190611f74565b60405180910390f35b6106d260048036038101906106cd9190611f38565b61137f565b6040516106df9190611f74565b60405180910390f35b61070260048036038101906106fd9190611f38565b6113ba565b60405161070f9190611f74565b60405180910390f35b6107206113f5565b60405161072d9190611f74565b60405180910390f35b610750600480360381019061074b9190611f38565b611414565b60405161075d9190611f74565b60405180910390f35b610780600480360381019061077b9190611f38565b611450565b60405161078d9190611f74565b60405180910390f35b6107b060048036038101906107ab9190611f38565b61148c565b6040516107bd9190611f74565b60405180910390f35b6107e060048036038101906107db9190611f38565b6114cb565b6040516107ed9190611f74565b60405180910390f35b610810600480360381019061080b9190611f38565b611506565b60405161081d9190611f74565b60405180910390f35b610840600480360381019061083b9190611f38565b611544565b60405161084d9190611f74565b60405180910390f35b610870600480360381019061086b9190611f38565b61157f565b60405161087d9190611f74565b60405180910390f35b6108a0600480360381019061089b9190611f38565b6115c4565b6040516108ad9190611f74565b60405180910390f35b6108d060048036038101906108cb9190611f38565b611600565b6040516108dd9190611f74565b60405180910390f35b61090060048036038101906108fb9190611f38565b61163e565b60405161090d9190611f74565b60405180910390f35b610930600480360381019061092b9190611f38565b611679565b60405161093d9190611f74565b60405180910390f35b610960600480360381019061095b9190611f38565b6116b7565b60405161096d9190611f74565b60405180910390f35b610990600480360381019061098b9190611f38565b6116f3565b60405161099d9190611f74565b60405180910390f35b6109c060048036038101906109bb9190611f38565b61172e565b6040516109cd9190611f74565b60405180910390f35b6109f060048036038101906109eb9190611f38565b61176a565b6040516109fd9190611f74565b60405180910390f35b610a206004803603810190610a1b9190611f38565b6117c7565b604051610a2d9190611f74565b60405180910390f35b610a506004803603810190610a4b9190611f38565b611806565b604051610a5d9190611f74565b60405180910390f35b610a806004803603810190610a7b9190611f38565b611841565b604051610a8d9190611f74565b60405180910390f35b610ab06004803603810190610aab9190611f38565b61188d565b604051610abd9190611f74565b60405180910390f35b610ae06004803603810190610adb9190611f38565b6118cd565b604051610aed9190611f74565b60405180910390f35b610b106004803603810190610b0b9190611f38565b611908565b604051610b1d9190611f74565b60405180910390f35b610b406004803603810190610b3b9190611f38565b611943565b604051610b4d9190611f74565b60405180910390f35b610b706004803603810190610b6b9190611f38565b61197f565b604051610b7d9190611f74565b60405180910390f35b610ba06004803603810190610b9b9190611f38565b6119bb565b604051610bad9190611f74565b60405180910390f35b610bd06004803603810190610bcb9190611f38565b6119f6565b604051610bdd9190611f74565b60405180910390f35b610c006004803603810190610bfb9190611f38565b611a31565b604051610c0d9190611f74565b60405180910390f35b610c306004803603810190610c2b9190611f38565b611a6c565b604051610c3d9190611f74565b60405180910390f35b610c606004803603810190610c5b9190611f38565b611aa7565b604051610c6d9190611f74565b60405180910390f35b610c906004803603810190610c8b9190611f38565b611aeb565b604051610c9d9190611f74565b60405180910390f35b610cc06004803603810190610cbb9190611f38565b611b27565b604051610ccd9190611f74565b60405180910390f35b610cf06004803603810190610ceb9190611f38565b611b62565b604051610cfd9190611f74565b60405180910390f35b610d206004803603810190610d1b9190611f38565b611ba0565b604051610d2d9190611f74565b60405180910390f35b610d506004803603810190610d4b9190611f38565b611bdd565b604051610d5d9190611f74565b60405180910390f35b610d806004803603810190610d7b9190611f38565b611c17565b604051610d8d9190611f74565b60405180910390f35b610db06004803603810190610dab9190611f38565b611c53565b604051610dbd9190611f74565b60405180910390f35b610de06004803603810190610ddb9190611f38565b611c8f565b604051610ded9190611f74565b60405180910390f35b610e106004803603810190610e0b9190611f38565b611cea565b604051610e1d9190611f74565b60405180910390f35b610e406004803603810190610e3b9190611f38565b611d2c565b604051610e4d9190611f74565b60405180910390f35b610e706004803603810190610e6b9190611f38565b611d68565b604051610e7d9190611f74565b60405180910390f35b610ea06004803603810190610e9b9190611f38565b611da5565b604051610ead9190611f74565b60405180910390f35b610ed06004803603810190610ecb9190611f38565b611de7565b604051610edd9190611f74565b60405180910390f35b610f006004803603810190610efb9190611f38565b611e23565b604051610f0d9190611f74565b60405180910390f35b610f306004803603810190610f2b9190611f38565b611e61565b604051610f3d9190611f74565b60405180910390f35b610f606004803603810190610f5b9190611f38565b611ea0565b604051610f6d9190611f74565b60405180910390f35b6000610f806113f5565b50600065deadbeef003690506000805b84811015610fa657369150600181019050610f90565b505080915050919050565b6000610fbb6113f5565b50600065deadbeef003290506000805b84811015610fe157329150600181019050610fcb565b505080915050919050565b6000610ff66113f5565b50600065deadbeef0052905060005b8381101561101c5781600052600181019050611005565b5080915050919050565b60006110306113f5565b50600065deadbeef0001905060005b838110156110585760008201915060018101905061103f565b5080915050919050565b60008054905090565b60006110756113f5565b50600065deadbeef0017905060005b8381101561109d57600082179150600181019050611084565b5080915050919050565b60006110b16113f5565b50600065deadbeef003490506000805b848110156110d7573491506001810190506110c1565b505080915050919050565b60006110ec6113f5565b50600065deadbeef0006905060005b83811015611133577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820691506001810190506110fb565b5080915050919050565b60006111476113f5565b50600065deadbeef001390506000805b8481101561117057600183139150600181019050611157565b505080915050919050565b60006111856113f5565b50600065deadbeef002090507fffffffff000000000000000000000000000000000000000000000000000000006000526000805b848110156111d357600460002091506001810190506111b9565b507f29045a592007d0c246ef02c2223570da9522d0cf0f73282c79a1bc8f0bb2c238811461120057600091505b5080915050919050565b60006112146113f5565b50600065deadbeef00a490508060105260005b83811015611246576004600360028360066010a4600181019050611227565b5080915050919050565b600061125a6113f5565b50600065deadbeef001a90506000805b84811015611283578260001a915060018101905061126a565b505080915050919050565b60006112986113f5565b50600065deadbeef001b905060005b838110156112c0578160001b91506001810190506112a7565b5080915050919050565b60006112d46113f5565b50600065deadbeef004290506000805b848110156112fa574291506001810190506112e4565b505080915050919050565b600061130f6113f5565b50600065deadbeef0031905060003060005b858110156113385781319250600181019050611321565b50505080915050919050565b600061134e6113f5565b50600065deadbeef004890506000805b848110156113745748915060018101905061135e565b505080915050919050565b60006113896113f5565b50600065deadbeef003d90506000805b848110156113af573d9150600181019050611399565b505080915050919050565b60006113c46113f5565b50600065deadbeef004390506000805b848110156113ea574391506001810190506113d4565b505080915050919050565b600060016000546114069190611fbe565b600081905550600054905090565b600061141e6113f5565b50600065deadbeef0004905060005b838110156114465760018204915060018101905061142d565b5080915050919050565b600061145a6113f5565b50600065deadbeef0037905060005b8381101561148257602060008037600181019050611469565b5080915050919050565b60006114966113f5565b50600065deadbeef00a090508060105260005b838110156114c15760066010a06001810190506114a9565b5080915050919050565b60006114d56113f5565b50600065deadbeef003390506000805b848110156114fb573391506001810190506114e5565b505080915050919050565b60006115106113f5565b50600065deadbeef0053905060005b8381101561153a5763deadbeef60005260018101905061151f565b5080915050919050565b600061154e6113f5565b50600065deadbeef003a90506000805b84811015611574573a915060018101905061155e565b505080915050919050565b60006115896113f5565b50600065deadbeef0051905060008160005260005b848110156115b657600051915060018101905061159e565b508091505080915050919050565b60006115ce6113f5565b50600065deadbeef001d905060005b838110156115f6578160001d91506001810190506115dd565b5080915050919050565b600061160a6113f5565b50600065deadbeef001090506000805b848110156116335782600110915060018101905061161a565b505080915050919050565b60006116486113f5565b50600065deadbeef004490506000805b8481101561166e57449150600181019050611658565b505080915050919050565b60006116836113f5565b50600065deadbeef001190506000805b848110156116ac57600183119150600181019050611693565b505080915050919050565b60006116c16113f5565b50600065deadbeef003e905060005b838110156116e95760206000803e6001810190506116d0565b5080915050919050565b60006116fd6113f5565b50600065deadbeef004590506000805b848110156117235745915060018101905061170d565b505080915050919050565b60006117386113f5565b50600065deadbeef0002905060005b8381101561176057600182029150600181019050611747565b5080915050919050565b60006117746113f5565b50600065deadbeef0008905060005b838110156117bd577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600083089150600181019050611783565b5080915050919050565b60006117d16113f5565b50600065deadbeef005490508060005560005b838110156117fc5760005491506001810190506117e4565b5080915050919050565b60006118106113f5565b50600065deadbeef005a90506000805b84811015611836575a9150600181019050611820565b505080915050919050565b600061184b6113f5565b50600065deadbeef0019905060005b83811015611871578119915060018101905061185a565b5065deadbeef0019811461188457801990505b80915050919050565b60006118976113f5565b50600065deadbeef00a190508060105260005b838110156118c3578060066010a16001810190506118aa565b5080915050919050565b60006118d76113f5565b50600065deadbeef0016905060005b838110156118fe5781821691506001810190506118e6565b5080915050919050565b60006119126113f5565b50600065deadbeef004690506000805b8481101561193857469150600181019050611922565b505080915050919050565b600061194d6113f5565b50600065deadbeef0005905060005b838110156119755760018205915060018101905061195c565b5080915050919050565b60006119896113f5565b50600065deadbeef0039905060005b838110156119b157602060008039600181019050611998565b5080915050919050565b60006119c56113f5565b50600065deadbeef005990506000805b848110156119eb575991506001810190506119d5565b505080915050919050565b6000611a006113f5565b50600065deadbeef003890506000805b84811015611a2657389150600181019050611a10565b505080915050919050565b6000611a3b6113f5565b50600065deadbeef004190506000805b84811015611a6157419150600181019050611a4b565b505080915050919050565b6000611a766113f5565b50600065deadbeef003090506000805b84811015611a9c57309150600181019050611a86565b505080915050919050565b6000611ab16113f5565b50600065deadbeef00a390508060105260005b83811015611ae157600360028260066010a3600181019050611ac4565b5080915050919050565b6000611af56113f5565b50600065deadbeef000b905060005b83811015611b1d578160200b9150600181019050611b04565b5080915050919050565b6000611b316113f5565b50600065deadbeef004790506000805b84811015611b5757479150600181019050611b41565b505080915050919050565b6000611b6c6113f5565b50600065deadbeef001c90506000805b84811015611b95578260001c9250600181019050611b7c565b505080915050919050565b6000611baa6113f5565b50600065deadbeef003590506000805b84811015611bd2576000359150600181019050611bba565b505080915050919050565b6000611be76113f5565b50600065deadbeef0055905060005b83811015611c0d5781600055600181019050611bf6565b5080915050919050565b6000611c216113f5565b50600065deadbeef0018905060005b83811015611c4957600082189150600181019050611c30565b5080915050919050565b6000611c5d6113f5565b50600065deadbeef0003905060005b83811015611c8557600082039150600181019050611c6c565b5080915050919050565b6000611c996113f5565b50600065deadbeef0007905060005b83811015611ce0577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82079150600181019050611ca8565b5080915050919050565b6000611cf46113f5565b50600065deadbeef00a290508060105260005b83811015611d225760028160066010a2600181019050611d07565b5080915050919050565b6000611d366113f5565b50600065deadbeef000a905060005b83811015611d5e576001820a9150600181019050611d45565b5080915050919050565b6000611d726113f5565b50600065deadbeef001490506000805b84811015611d9a578283149150600181019050611d82565b505080915050919050565b6000611daf6113f5565b50600065deadbeef0040905060006001430360005b85811015611ddb5781409250600181019050611dc4565b50505080915050919050565b6000611df16113f5565b50600065deadbeef001590506000805b84811015611e185782159150600181019050611e01565b505080915050919050565b6000611e2d6113f5565b50600065deadbeef001290506000805b84811015611e5657826001129150600181019050611e3d565b505080915050919050565b6000611e6b6113f5565b50600065deadbeef003b905060003060005b85811015611e9457813b9250600181019050611e7d565b50505080915050919050565b6000611eaa6113f5565b50600065deadbeef0009905060005b83811015611ef3577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600183099150600181019050611eb9565b5080915050919050565b600080fd5b6000819050919050565b611f1581611f02565b8114611f2057600080fd5b50565b600081359050611f3281611f0c565b92915050565b600060208284031215611f4e57611f4d611efd565b5b6000611f5c84828501611f23565b91505092915050565b611f6e81611f02565b82525050565b6000602082019050611f896000830184611f65565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611fc982611f02565b9150611fd483611f02565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561200957612008611f8f565b5b82820190509291505056fea2646970667358221220d8cc57ca5ddb766a9c3715937d4109a2381dca9d4ee1024c67fabd7fde18383364736f6c634300080f0033"

var abiLoadTesterStr = `[{"inputs":[],"name":"getCallCounter","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"inc","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testADD","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testADDMOD","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testADDRESS","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testAND","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testBALANCE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testBASEFEE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testBLOCKHASH","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testBYTE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testCALLDATACOPY","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testCALLDATALOAD","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testCALLDATASIZE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testCALLER","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testCALLVALUE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testCHAINID","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testCODECOPY","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testCODESIZE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testCOINBASE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testDIFFICULTY","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testDIV","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testEQ","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testEXP","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testEXTCODESIZE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testGAS","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testGASLIMIT","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testGASPRICE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testGT","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testISZERO","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testLOG0","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testLOG1","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testLOG2","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testLOG3","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testLOG4","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testLT","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testMLOAD","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testMOD","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testMSIZE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testMSTORE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testMSTORE8","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testMUL","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testMULMOD","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testNOT","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testNUMBER","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testOR","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testORIGIN","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testRETURNDATACOPY","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testRETURNDATASIZE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testSAR","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testSDIV","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testSELFBALANCE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testSGT","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testSHA3","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testSHL","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testSHR","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testSIGNEXTEND","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testSLOAD","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testSLT","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testSMOD","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testSSTORE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testSUB","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testTIMESTAMP","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"x","type":"uint256"}],"name":"testXOR","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"}]`
