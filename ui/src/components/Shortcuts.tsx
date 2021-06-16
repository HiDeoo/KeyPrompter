import { AnimatePresence } from 'framer-motion'
import { nanoid } from 'nanoid'
import { useCallback, useEffect, useState } from 'react'
import { useErrorHandler } from 'react-error-boundary'

import Shortcut from './Shortcut'
import { SHORTCUT_DURATION } from '../constants/shortcut'
import { useWebSocket } from '../contexts/WebSocketContext'
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

  const [shortcutEvents, setShortcutEvents] = useState<ShortcutEvent[]>([])

  ws.onclose = () => {
    handleError(new Error('Connection to websocket closed.'))
  }

  const onMessage = useCallback(
    (event: MessageEvent) => {
      let shortcutEventData: ShortcutEventData

      try {
        const data = JSON.parse(event.data)

        if (isShortcutEventData(data)) {
          shortcutEventData = data
        } else {
          throw new Error('Invalid shortcut event data.')
        }
      } catch (error) {
        handleError(new Error('Unable to parse shortcut event data.'))
      }

      const id = nanoid()

      setShortcutEvents((prevShortcutEvents) => addShortcutEvent(prevShortcutEvents, { ...shortcutEventData, id }))

      setTimeout(() => {
        setShortcutEvents((prevShortcutEvents) => removeShortcutEvent(prevShortcutEvents, id))
      }, SHORTCUT_DURATION * 1000)
    },
    [handleError]
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
