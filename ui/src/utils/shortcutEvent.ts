import { SHORTCUT_MAX_COUNT } from '../constants/shortcut'

const ModifierMap: Record<string, string> = {
  command: '⌘',
  control: '⌃',
  fn: 'Fn',
  option: '⌥',
  shift: '⇧',
}

const CharacterMap: Record<string, string> = {
  '\t': '⇥',
  ' ': '␣',
}

export function addShortcutEvent(shortcutEvents: ShortcutEvent[], shortcutEvent: ShortcutEvent): ShortcutEvent[] {
  const previousShortcutEvent = shortcutEvents[shortcutEvents.length - 1]

  if (previousShortcutEvent && areShortcutEventsIdentical(previousShortcutEvent, shortcutEvent)) {
    return [...shortcutEvents.slice(0, -1), { ...previousShortcutEvent, count: previousShortcutEvent.count + 1 }]
  }

  shortcutEvent.count = 1

  return [...shortcutEvents.slice(-1 * SHORTCUT_MAX_COUNT + 1), shortcutEvent]
}

export function removeShortcutEvent(shortcutEvents: ShortcutEvent[], id: ShortcutEvent['id']): ShortcutEvent[] {
  const shortcutEventIndex = shortcutEvents.findIndex((shortcutEvent) => shortcutEvent.id === id)

  if (shortcutEventIndex !== -1) {
    return [...shortcutEvents.slice(0, shortcutEventIndex), ...shortcutEvents.slice(shortcutEventIndex + 1)]
  }

  return shortcutEvents
}

export function isShortcutEventData(data: unknown): data is ShortcutEventData {
  return (
    typeof data === 'object' &&
    typeof (data as ShortcutEventData).character === 'string' &&
    typeof (data as ShortcutEventData).code === 'number' &&
    typeof (data as ShortcutEventData).modifiers === 'object' &&
    typeof (data as ShortcutEventData).modifiers.command === 'boolean' &&
    typeof (data as ShortcutEventData).modifiers.control === 'boolean' &&
    typeof (data as ShortcutEventData).modifiers.fn === 'boolean' &&
    typeof (data as ShortcutEventData).modifiers.option === 'boolean' &&
    typeof (data as ShortcutEventData).modifiers.shift === 'boolean'
  )
}

export function getShortcut(shortcutEvent: ShortcutEvent): string {
  return getShortcutModifiers(shortcutEvent).concat(getShortcutCharacter(shortcutEvent).toUpperCase())
}

function getShortcutModifiers({ modifiers }: ShortcutEvent): string {
  let shortcutModifiers = ''

  Object.entries(modifiers).forEach(([modifier, pressed]) => {
    if (pressed && ModifierMap.hasOwnProperty(modifier)) {
      shortcutModifiers = shortcutModifiers.concat(ModifierMap[modifier])
    }
  })

  return shortcutModifiers
}

function getShortcutCharacter({ character }: ShortcutEvent): string {
  if (CharacterMap.hasOwnProperty(character)) {
    return CharacterMap[character]
  }

  return character
}

function areShortcutEventsIdentical(l: ShortcutEvent, r: ShortcutEvent): boolean {
  return (
    l.character === r.character &&
    l.code === r.code &&
    l.modifiers.command === r.modifiers.command &&
    l.modifiers.control === r.modifiers.control &&
    l.modifiers.fn === r.modifiers.fn &&
    l.modifiers.option === r.modifiers.option &&
    l.modifiers.shift === r.modifiers.shift
  )
}

export interface ShortcutEvent {
  id: string
  character: string
  code: number
  count: number
  modifiers: {
    command: boolean
    control: boolean
    fn: boolean
    option: boolean
    shift: boolean
  }
}

export type ShortcutEventData = Omit<ShortcutEvent, 'id'>
