import { createContext, useContext } from 'react'

// TODO(HiDeoo)
const WSServerUrl = 'ws://localhost:3333/ws' // `ws://${window.location.host}/ws`

const WebSocketContext = createContext(new WebSocket(WSServerUrl))

export const useWebSocket = (): WebSocket => useContext(WebSocketContext)
