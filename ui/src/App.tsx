import { useEffect } from 'react'
import './App.css'

const socket = new WebSocket('ws://localhost:8484/ws')

function App() {
  useEffect(() => {
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

  function send() {
    socket.send('hello from client')
  }

  return (
    <div className="hello">
      Hello <button onClick={send}>Send</button>
    </div>
  )
}

export default App
