import Web3 from "web3";

const EnableWeb3 = (enable: boolean) => {
    return function (target: any) {
        if (!enable) {
            return class extends target {
                public async establish(): Promise<void> { }
                public async getAccounts(): Promise<string[] | undefined> { return }
                public async getBalance(address: string): Promise<string | undefined> { return }
            }
        }
    }
}

// truffle web3
@EnableWeb3(true)
export class TWeb3 {
    private static web3?: Web3

    public async establish(): Promise<void> {
        if (typeof window['ethereum'] !== 'undefined') {
            TWeb3.web3 = new Web3(window['ethereum']);
            // @ts-ignore
            const enable = await ethereum.enable();
        } else {
            throw new Error("未能初始化Web3")
        }
    }

    public async getAccounts(): Promise<string[] | undefined> {
        let ret: string[] = []
        if (typeof TWeb3.web3 !== 'undefined') {
            return await TWeb3.web3.eth.getAccounts();
        }
        return ret
    }

    public async getBalance(address: string): Promise<string | undefined> {
        if (typeof TWeb3.web3 !== 'undefined') {
            return await TWeb3.web3.eth.getBalance(address);
        }
        return
    }
}
