import { motion } from 'framer-motion'

import { useConfig } from '../contexts/ConfigContext'
import { getShortcutRepresentation, ShortcutEvent } from '../utils/shortcutEvent'

import './Shortcut.css'

const Shortcut: React.FC<Props> = ({ event }) => {
  const { configuration } = useConfig()

  return (
    <motion.li
      className="shortcut"
      initial="initial"
      animate="visible"
      exit="exit"
      variants={{
        initial: { opacity: 0.5, y: '100%' },
        visible: { opacity: 1, x: 0, y: 0, transition: { duration: configuration.duration * 0.3, type: 'spring' } },
        exit: { opacity: 0, x: '-100%', transition: { duration: configuration.duration * 0.2 } },
      }}
    >
      <span className="description">
        {getShortcutRepresentation(event)}
        {event.count > 1 && <span className="count">x{event.count}</span>}
      </span>
    </motion.li>
  )
}

export default Shortcut

interface Props {
  event: ShortcutEvent
}
