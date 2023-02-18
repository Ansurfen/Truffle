type TFWebSocketHandle = () => void
type TFWebSocketCallback = (this: WebSocket, e: Event) => void
type TFWebSocketMessageCallback = (this: WebSocket, e: any) => void
type StringMap<T> = Map<string, T>
type TFWebSocketPayload = string | ArrayBufferLike | Blob | ArrayBufferView

class TFWebSocket {
    private conns: StringMap<TFWebSocketClient>
    private strict: boolean // automatically create websocket connection when calling on or emit function
    private count = 0

    constructor(strict: boolean = false) {
        this.conns = new Map<string, TFWebSocketClient>()
        this.strict = strict
    }

    public open(key: string, opt: TFWebSocketOpt, handle: TFWebSocketCallback): void {
        opt.openHandle = new Map()
        opt.messageHandle = new Map()
        opt.openHandle.set(this.count.toString(), handle)
        let cli = new TFWebSocketClient(opt)
        this.conns.set(key, cli)
    }

    public on(key: string, handle: TFWebSocketMessageCallback): void {
        if (!this.strict) {
            let conn = this.conns.get(key)
            if (typeof conn !== undefined) {
                if (typeof conn?.opt.messageHandle !== undefined) {
                    conn!.opt.messageHandle?.set(this.count.toString(), handle)
                    this.count++
                }
            } else {
                conn = new TFWebSocketClient(<TFWebSocketOpt>{
                    url: "",
                    heartTime: 3000,
                    reconnTime: 3000,
                    reconnCount: 3,
                })
                this.conns.set(key, conn)
            }
        }
    }

    public emit(key: string, data: TFWebSocketPayload, retry: boolean = false): void {
        let conn = this.conns.get(key)
        if (typeof conn !== undefined)
            conn!.send(data, retry)
        else
            console.log(key, " is invalid")
    }

    public close(key: string) {
        let conn = this.conns.get(key)
        if (typeof conn !== undefined) {
            // release conn
            if (conn!.unref()) conn!.free()
        }
    }
}

interface TFWebSocketOpt {
    url: string | null
    heartTime: number
    reconnTime?: number
    reconnCount?: number
    openRequest?: string
    openHandle?: StringMap<TFWebSocketCallback>
    messageHandle?: StringMap<TFWebSocketMessageCallback>
    closeHandle?: StringMap<TFWebSocketCallback>
    errorHandle?: StringMap<TFWebSocketCallback>
}

class Heart {
    private heartTimeout!: number
    private heartTimer!: number
    private timeout: number

    constructor(timeout: number = 5000) {
        this.timeout = timeout
    }

    public reset(): void {
        clearTimeout(this.heartTimeout)
        clearTimeout(this.heartTimer)
    }

    public setTimeout(timeout: number): void {
        this.timeout = timeout
    }

    public beat(func: Function): void {
        this.heartTimeout = setTimeout((e: Event) => {
            func(e)
            this.heartTimer = setTimeout((e: Event) => {
                func(e)
                this.reset()
                this.beat(func)
            }, this.timeout)
        }, this.timeout)
    }
}

// heartTime = -1 close heartbeat
const defaultTFWebSocketOpt = (url: string): TFWebSocketOpt => {
    return {
        url: url,
        heartTime: 1000,
        openHandle: new Map(),
        closeHandle: new Map(),
        errorHandle: new Map(),
        messageHandle: new Map
    }
}


class TFWebSocketClient {
    private refCount: number
    protected ws!: WebSocket
    public opt!: TFWebSocketOpt
    protected heart!: Heart
    protected reconnTimer = 0
    protected reconnCount = 10

    constructor(opt: TFWebSocketOpt) {
        this.refCount = 0
        this.opt = opt
        this.malloc()
    }

    public ref(): void {
        this.refCount++
    }

    public unref(): boolean {
        this.refCount--
        return this.refCount <= 0
    }

    public count(): number {
        return this.refCount
    }

    public malloc(): void {
        if (!window['WebSocket']) {
            throw new Error('The platform no support current browser')
        }
        if (!this.opt.url) {
            throw new Error('invalid url')
        }
        this.heart = new Heart()
        this.ws = new WebSocket(this.opt.url)
        this.onOpen(this.opt.openRequest as string, this.opt.openHandle as StringMap<TFWebSocketCallback>)
        this.onError(this.opt.openHandle as StringMap<TFWebSocketCallback>)
        this.onMessage(this.opt.messageHandle as StringMap<TFWebSocketMessageCallback>)
        this.onClose(this.opt.closeHandle as StringMap<TFWebSocketCallback>)
    }

    public send(data: TFWebSocketPayload, retry: boolean = false): void {
        if (this.ws.readyState === WebSocket.OPEN)
            this.ws.send(data)
        else if (retry && this.ws.readyState === WebSocket.CONNECTING)
            setTimeout(() => {
                this.ws.send(data)
            }, 3000)
    }

    public free(): void {
        this.ws.close()
    }

    private onOpen(request: string, callbacks: StringMap<TFWebSocketCallback>): void {
        this.ws.onopen = (e: Event) => {
            clearTimeout(this.reconnTimer)
            this.opt.reconnCount = this.reconnCount
            this.heart.reset()
            this.heart.beat(() => {
                this.ws.send(request)
            })
            if (typeof callbacks === 'object') {
                callbacks.forEach(callback => {
                    callback.bind(this.ws)(e)
                })
            } else {
                typeof this.opt.openHandle === 'object' && this.opt.openHandle.forEach(handle => { handle.bind(this.ws)(e) })
            }
        }
    }

    private onError(callbacks: StringMap<TFWebSocketCallback>): void {
        this.ws.onerror = (e: Event) => {
            if (typeof callbacks === 'object') {
                callbacks.forEach(callback => {
                    callback.bind(this.ws)(e)
                })
            } else {
                typeof this.opt.errorHandle === 'object' && this.opt.errorHandle.forEach(handle => { handle.bind(this.ws)(e) })
            }
        }
    }

    private onMessage(callbacks: StringMap<TFWebSocketMessageCallback>): void {
        this.ws.onmessage = (e: MessageEvent<string>) => {
            // prehandle
            if (typeof callbacks === 'object') {
                console.log(callbacks)
                callbacks.forEach(callback => {
                    callback.bind(this.ws)(e.data)
                })
            } else {
                typeof this.opt.messageHandle === 'object' && this.opt.messageHandle.forEach(handle => { handle.bind(this.ws)(e.data) })
            }
        }
    }

    private onClose(callbacks: StringMap<TFWebSocketCallback>): void {
        this.ws.onclose = (e: Event) => {
            // prehandle
            if (typeof callbacks === 'object') {
                callbacks.forEach(callback => {
                    callback.bind(this.ws)(e)
                })
            } else {
                typeof this.opt.closeHandle === 'object' && this.opt.closeHandle.forEach(handle => { handle.bind(this.ws)(e) })
            }
        }
    }
}

export {
    TFWebSocketClient,
    TFWebSocket,
    type TFWebSocketOpt,
    type TFWebSocketCallback,
}