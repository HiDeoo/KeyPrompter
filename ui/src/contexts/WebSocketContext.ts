import { createContext, useContext } from 'react'

const WSServerUrl = `ws://${window.location.host}/ws`

const WebSocketContext = createContext(new WebSocket(WSServerUrl))

export const useWebSocket = (): WebSocket => useContext(WebSocketContext)
