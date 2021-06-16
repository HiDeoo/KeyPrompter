import { nanoid } from 'nanoid'
import { useCallback, useEffect, useState } from 'react'
import { useErrorHandler } from 'react-error-boundary'

import { useWebSocket } from './WebSocketContext'

const ShortcutDuration = 5000 // In milliseconds.
const ShortcutMaxCount = 5

const Shortcuts: React.FC = () => {
  const ws = useWebSocket()
  const handleError = useErrorHandler()

  const [shortcuts, setShortcuts] = useState<Shortcut[]>([])

  ws.onclose = () => {
    handleError(new Error('Connection to websocket closed.'))
  }

  const onMessage = useCallback(() => {
    const id = nanoid()

    setShortcuts((prevShortcuts) => addShortcut(prevShortcuts, id))

    setTimeout(() => {
      setShortcuts((prevShortcuts) => removeShortcut(prevShortcuts, id))
    }, ShortcutDuration)
  }, [])

  useEffect(() => {
    ws.addEventListener('message', onMessage)

    return () => {
      ws.removeEventListener('message', onMessage)
    }
  }, [ws, onMessage])

  return (
    <div>
      {shortcuts.map((tt) => (
        <div key={tt}>{tt}</div>
      ))}
    </div>
  )
}

export default Shortcuts

function addShortcut(shortcuts: Shortcut[], shortcut: Shortcut): Shortcut[] {
  return [...shortcuts.slice(-1 * ShortcutMaxCount + 1), shortcut]
}

function removeShortcut(shortcuts: Shortcut[], shortcut: Shortcut): Shortcut[] {
  const shortcutIndex = shortcuts.indexOf(shortcut)

  if (shortcutIndex !== -1) {
    return [...shortcuts.slice(0, shortcutIndex), ...shortcuts.slice(shortcutIndex + 1)]
  }

  return shortcuts
}

type Shortcut = string
