import { useEffect } from 'react'
import './App.css'

function App() {
  useEffect(() => {
    const socket = new WebSocket('ws://localhost:8484/ws')

    socket.onopen = () => {
      console.log('onopen')
    }

    socket.onmessage = (msg) => {
      console.log('onmessage', msg)
    }

    socket.onclose = (event) => {
      console.log('onclose', event)
    }

    socket.onerror = (error) => {
      console.log('onerror', error)
    }
  }, [])

  return <div className="hello">Hello</div>
}

export default App
