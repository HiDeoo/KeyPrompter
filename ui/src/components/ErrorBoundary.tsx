import { ErrorBoundary as Boundary, ErrorBoundaryPropsWithComponent, FallbackProps } from 'react-error-boundary'

import { useWebSocket } from '../contexts/WebSocketContext'

const ErrorBoundary: React.FC = ({ children }) => {
  const ws = useWebSocket()

  const onError: ErrorBoundaryPropsWithComponent['onError'] = (error) => {
    if (ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify({ type: 'error', message: error.message }))
    }
  }

  return (
    <Boundary FallbackComponent={Fallback} onError={onError}>
      {children}
    </Boundary>
  )
}

export default ErrorBoundary

const Fallback: React.FC<FallbackProps> = ({ error }) => {
  // TODO(HiDeoo)
  return (
    <div role="alert">
      <p>Something went wrong:</p>
      <pre>{error.message}</pre>
    </div>
  )
}
