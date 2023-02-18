const TMap = (window as any).TMap

type TFPosition = {
    lat: string | number,
    lng: string | number,
    height?: number
}

interface TFMapOpt {
    center: any
    zoom: number
    viewMode: string
}

interface TFMarkerOpt {
    styles: any
    geometries: any
}

const createTFPosition = (pos: TFPosition) => {
    return new TMap.LatLng(pos.lat, pos.lng, pos.height)
}

const createTFMarkerStyle = (markerStyle: any) => {
    return new TMap.MarkerStyle(markerStyle)
}

// truffle map
class TFMap {
    private map: any
    private marker: any[]
    private opt: TFMapOpt
    private models: any[]

    constructor(opt: TFMapOpt, pos: TFPosition, selector: string) {
        let box = document.querySelector<HTMLElement>(selector)
        if (box === null) {
            throw new Error('Fail to get selector')
        }
        opt.center = new TMap.LatLng(pos.lat, pos.lng)
        this.opt = opt
        this.map = new TMap.Map(box, opt)
        this.marker = []
        this.models = []
    }

    public makeMarker(opt: TFMarkerOpt): void {
        let marker = new TMap.MultiMarker({
            map: this.map,
            styles: opt.styles,
            geometries: opt.geometries
        })
        this.marker.push(marker)
    }

    public makeInfoWnd(content: string, pos: TFPosition): void {
        let infoWnd = new TMap.InfoWindow({
            content: content,
            position: new TMap.LatLng(pos.lat, pos.lng),
            map: this.map
        })
    }

    public makeGLTFModel(url: string, id: string, pos: TFPosition) {
        let model = new TMap.model.GLTFModel({
            url: url,
            map: this.map,
            id: id,
            position: createTFPosition(pos),
            rotation: [0, -90, 0],
            scale: [20, 20, 30],
        })
        this.models.push(model)
        // model.on('loaded', () => {
        //     console.log('model loaded')
        // })
    }
}

class TFLocation {
    private key: string
    private name: string
    private geolocation: any

    constructor(key: string, name: string) {
        this.key = key
        this.name = name
    }

    public getLocation(func: Function): void {
        if (typeof this.geolocation === 'undefined') {
            // @ts-ignore
            this.geolocation = new qq.maps.Geolocation(this.key, this.name);
        }
        this.geolocation.getLocation((curPos: any) => {
            func(curPos)
        }, () => {
            console.log('get location err')
        })
    }
}

export {
    TFMap,
    type TFMapOpt,
    TFLocation,
    type TFPosition,
    type TFMarkerOpt,
    createTFPosition,
    createTFMarkerStyle
}