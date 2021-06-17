import { AnimatePresence } from 'framer-motion'
import { nanoid } from 'nanoid'
import { useCallback, useEffect, useState } from 'react'
import { useErrorHandler } from 'react-error-boundary'

import Shortcut from './Shortcut'
import { useConfig } from '../contexts/ConfigContext'
import { useWebSocket } from '../contexts/WebSocketContext'
import { isClientConfig, updateClientConfig } from '../utils/config'
import {
  addShortcutEvent,
  isShortcutEventData,
  removeShortcutEvent,
  ShortcutEvent,
  ShortcutEventData,
} from '../utils/shortcutEvent'

import './Shortcuts.css'

const Shortcuts: React.FC = () => {
  const ws = useWebSocket()
  const handleError = useErrorHandler()
  const { configuration, updateConfiguration } = useConfig()

  const [shortcutEvents, setShortcutEvents] = useState<ShortcutEvent[]>([])

  ws.onclose = () => {
    handleError(new Error('Connection to websocket closed.'))
  }

  const onMessage = useCallback(
    (event: MessageEvent) => {
      let shortcutEventData: ShortcutEventData | undefined

      try {
        const data = JSON.parse(event.data)

        if (isShortcutEventData(data)) {
          shortcutEventData = data
        } else if (isClientConfig(data)) {
          updateClientConfig(updateConfiguration, data)
        } else {
          throw new Error('Invalid data received from the server.')
        }
      } catch (error) {
        handleError(error instanceof SyntaxError ? new Error('Unable to parse server data.') : error)
      }

      if (shortcutEventData) {
        const id = nanoid()

        setShortcutEvents((prevShortcutEvents) =>
          addShortcutEvent(prevShortcutEvents, { ...shortcutEventData!, id }, configuration.count)
        )

        setTimeout(() => {
          setShortcutEvents((prevShortcutEvents) => removeShortcutEvent(prevShortcutEvents, id))
        }, configuration.duration * 1000)
      }
    },
    [configuration, handleError, updateConfiguration]
  )

  useEffect(() => {
    ws.addEventListener('message', onMessage)

    return () => {
      ws.removeEventListener('message', onMessage)
    }
  }, [ws, onMessage])

  return (
    <ul className="shortcuts">
      <AnimatePresence>
        {shortcutEvents.map((shortcutEvent) => (
          <Shortcut key={shortcutEvent.id} event={shortcutEvent} />
        ))}
      </AnimatePresence>
    </ul>
  )
}

export default Shortcuts
