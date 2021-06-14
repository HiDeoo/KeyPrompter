import { useEffect } from 'react'
import './App.css'

const socket = new WebSocket('ws://localhost:8484/ws')

function App() {
  useEffect(() => {
    socket.onopen = () => {
      console.log('onopen')
    }

    socket.onmessage = (msg) => {
      try {
        const data = JSON.parse(msg.data)
        console.log('data ', data)
      } catch (error) {
        console.log('error ', error)
      }
    }

    socket.onclose = (event) => {
      console.log('onclose', event)
    }

    socket.onerror = (error) => {
      console.log('onerror', error)
    }
  }, [])

  function send() {
    socket.send(JSON.stringify({ type: 'error', message: 'hello from client' }))
  }

  return (
    <div className="hello">
      Hello <button onClick={send}>Send</button>
    </div>
  )
}

export default App
