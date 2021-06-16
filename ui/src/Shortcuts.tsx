import { useCallback, useEffect } from 'react'
import { useErrorHandler } from 'react-error-boundary'

import { useWebSocket } from './WebSocketContext'

const Shortcuts: React.FC = () => {
  const ws = useWebSocket()
  const handleError = useErrorHandler()

  ws.onclose = () => {
    handleError(new Error('Connection to websocket closed.'))
  }

  const onMessage = useCallback(() => {
    console.log('MESSAGE')
  }, [])

  useEffect(() => {
    ws.addEventListener('message', onMessage)

    return () => {
      ws.removeEventListener('message', onMessage)
    }
  }, [ws, onMessage])

  return <div>Shortcuts2</div>
}

export default Shortcuts
