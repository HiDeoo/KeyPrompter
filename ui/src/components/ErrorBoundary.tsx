import { motion } from 'framer-motion'
import { ErrorBoundary as Boundary, ErrorBoundaryPropsWithComponent, FallbackProps } from 'react-error-boundary'

import { useWebSocket } from '../contexts/WebSocketContext'

import './ErrorBoundary.css'

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

const Fallback: React.FC<FallbackProps> = () => {
  return (
    <motion.div
      role="alert"
      className="error"
      animate={{ rotate: [6, -12, 12, -6, 0] }}
      transition={{ repeat: Infinity, repeatDelay: 5 }}
    >
      🤕
    </motion.div>
  )
}
