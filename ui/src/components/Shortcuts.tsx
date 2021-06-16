import { nanoid } from 'nanoid'
import { useCallback, useEffect, useState } from 'react'
import { useErrorHandler } from 'react-error-boundary'

import { SHORTCUT_DURATION } from '../constants/shortcut'
import { useWebSocket } from '../contexts/WebSocketContext'
import {
  addShortcutEvent,
  isShortcutEventData,
  removeShortcutEvent,
  ShortcutEvent,
  ShortcutEventData,
} from '../utils/shortcutEvent'

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
      }, SHORTCUT_DURATION)
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
    <div>
      {shortcutEvents.map((shortcutEvent) => (
        <div key={shortcutEvent.id}>{JSON.stringify(shortcutEvent)}</div>
      ))}
    </div>
  )
}

export default Shortcuts
