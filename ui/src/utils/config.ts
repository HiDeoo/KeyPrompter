export function isClientConfig(data: unknown): data is ClientConfig {
  return typeof data === 'object' && typeof (data as ClientConfig)['font-size'] === 'number'
}

export function updateCLientConfig(clientConfig: ClientConfig) {
  if (clientConfig['font-size']) {
    document.documentElement.style.setProperty('--font-size', `${clientConfig['font-size']}px`)
  }
}

export interface ClientConfig {
  'font-size': number
}
