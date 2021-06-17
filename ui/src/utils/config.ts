import { ConfigContextDefinition } from '../contexts/ConfigContext'

export function isClientConfig(data: unknown): data is ClientConfig {
  return (
    typeof data === 'object' &&
    typeof (data as ClientConfig)['bg-color'] === 'string' &&
    typeof (data as ClientConfig)['font-color'] === 'string' &&
    typeof (data as ClientConfig)['font-size'] === 'number'
  )
}

export function updateClientConfig(
  runtimeUpdater: ConfigContextDefinition['updateConfiguration'],
  clientConfig: ClientConfig
) {
  runtimeUpdater((prevConfig) => ({
    count: clientConfig.count > 0 ? clientConfig.count : prevConfig.count,
    duration: clientConfig.duration > 0 ? clientConfig.duration : prevConfig.duration,
  }))

  if (clientConfig['bg-color'].length > 0) {
    document.documentElement.style.setProperty('--bg-color', clientConfig['bg-color'])
  }

  if (clientConfig['font-color'].length > 0) {
    document.documentElement.style.setProperty('--font-color', clientConfig['font-color'])
  }

  if (clientConfig['font-size']) {
    document.documentElement.style.setProperty('--font-size', `${clientConfig['font-size']}px`)
  }
}

export type ClientConfig = UIClientConfig & RuntimeClientConfig

export interface UIClientConfig {
  'bg-color': string
  'font-color': string
  'font-size': number
}

export interface RuntimeClientConfig {
  count: number
  duration: number
}
