import { loadMicroApp, registerMicroApps, start } from "qiankun"

export interface IPlugin {
    name: string
    entry: string
    activeRule: string
    container: string
    props?: any
}

export const createPlugin = (name: string, entry: string, rule: string, func?: Function): IPlugin => {
    return {
        name: name,
        entry: entry,
        activeRule: rule,
        container: "#sandBox",
        props: {
            utils: {
                func: func
            }
        }
    }
}

export const installPlugins = (plugins: IPlugin[]) => {
    registerMicroApps(plugins)
    start()
}

export const installPlugin = (plugin: IPlugin) => {
    loadMicroApp(plugin)
}