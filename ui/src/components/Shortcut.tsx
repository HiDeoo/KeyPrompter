import { motion } from 'framer-motion'

import { SHORTCUT_DURATION } from '../constants/shortcut'
import { ShortcutEvent } from '../utils/shortcutEvent'

import './Shortcut.css'

const AnimationVariants = {
  initial: { opacity: 0.5, x: '100%' },
  visible: { opacity: 1, x: 0, transition: { duration: SHORTCUT_DURATION * 0.3, type: 'spring' } },
  exit: { opacity: 0, x: '100%', transition: { duration: SHORTCUT_DURATION * 0.2 } },
}

const Shortcut: React.FC<Props> = ({ event }) => {
  return (
    <motion.li className="shortcut" initial="initial" animate="visible" exit="exit" variants={AnimationVariants}>
      {event.id}
    </motion.li>
  )
}

export default Shortcut

interface Props {
  event: ShortcutEvent
}
