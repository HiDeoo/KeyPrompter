import { ConfigProvider } from '../contexts/ConfigContext'
import ErrorBoundary from './ErrorBoundary'
import Shortcuts from './Shortcuts'

function App() {
  return (
    <ErrorBoundary>
      <ConfigProvider>
        <Shortcuts />
      </ConfigProvider>
    </ErrorBoundary>
  )
}

export default App
