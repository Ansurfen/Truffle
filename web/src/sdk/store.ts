import { legacy_createStore as createStore } from "redux"

type storeType = "MOUNT" | "UNMOUNT" | "SUBSCRIBE"
export type Store = any
export type opcode = number | storeType
export type operand = string | {}
export interface Instrument {
    op?: opcode
    arg0?: operand
    arg1?: operand
    arg2?: operand
}
export const op: Record<storeType, opcode> = {
    MOUNT: 0,
    UNMOUNT: 1,
    SUBSCRIBE: 2
};
type TFStoreCallback<T> = (state: T) => void
type TFStoreUnsubscribe = () => void;
export interface ITFStore<T> {
    state: T
    subscribe: (state: TFStoreCallback<T>) => TFStoreUnsubscribe;
}

class TFStore<T> implements ITFStore<T> {
    public state!: T;

    constructor(state: T) {
        this.state = state
    }

    public subscribe(state: TFStoreCallback<T>): TFStoreUnsubscribe {
        return tfvm.subscribe(() => state(this.state))
    };
}

function vmOpt(state = {
    stores: <Map<string, Store>>{}
}, action: { type: Instrument; }) {
    switch (action.type.op) {
        case op.MOUNT:
            state.stores.set("a", {})
            break
        case op.UNMOUNT:

        default:

    }
    return state
}

// truffle virtual machine
let tfvm = createStore(vmOpt)

export const defineStore = <Id extends string, T>(id: Id, state: T): ITFStore<T> => {
    tfvm.dispatch({ type: { op: op.MOUNT, arg0: state as operand } })
    return new TFStore<T>(tfvm.getState().stores.get(id))
}

export const clearStore = <Id extends string>(id: Id): void => {
    tfvm.dispatch({ type: { op: op.UNMOUNT, arg0: id } })
}

function test() {
    const store = defineStore("test", {
        a: 0,
        b: "",
        c: false,
    })
    console.log(store.state.a)
    let TD = store.subscribe((state) => {
        console.log(state.a)
    })
    TD()
    clearStore("test")
}