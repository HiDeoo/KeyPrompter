import { createContext, useContext, useState } from 'react'
import { RuntimeClientConfig } from '../utils/config'

const InitialConfiguration: RuntimeClientConfig = {
  /**
   * Maximum number of shortcuts visible at the same time.
   */
  count: 5,
  /**
   * Duration in seconds during when a shortcut is visible on screen.
   */
  duration: 5,
}

const ConfigContext = createContext<ConfigContextDefinition | undefined>(undefined)

export const useConfig = (): ConfigContextDefinition => {
  const config = useContext(ConfigContext)

  if (!config) {
    throw new Error('useConfig must be used within a ConfigProvider.')
  }

  return config
}

export const ConfigProvider: React.FC = ({ children }) => {
  const [configuration, setConfiguration] = useState<RuntimeClientConfig>(InitialConfiguration)

  return (
    <ConfigContext.Provider value={{ configuration, updateConfiguration: setConfiguration }}>
      {children}
    </ConfigContext.Provider>
  )
}

export interface ConfigContextDefinition {
  configuration: RuntimeClientConfig
  updateConfiguration: React.Dispatch<React.SetStateAction<RuntimeClientConfig>>
}
