'use client'
import { useState, useEffect, useContext } from 'react'
import { FiArrowRight } from 'react-icons/fi'
import { FormContext } from '@/app/context/form.context'
import styles from './style.module.css'
import axios from 'axios'

interface Event {
  id: string
  name: string
  description: string
  image_url: string
  currency: string
  price: number
  event_date: string
  created_at: string
  updated_at: string
}

export default function ListEvents (): React.ReactNode {
  const { setNextPage, nextPage, setEvent } = useContext(FormContext)
  const [selectEventId, setSelectEventId] = useState('')
  const [events, setEvents] = useState<Event[]>([])
  const [error, setError] = useState<string>('')

  useEffect(() => {
    setError('')
    async function getEvents (): Promise<void> {
      try {
        const responseGetToken: { data: { token: string } } = await axios.post(
          `${process.env.NEXT_PUBLIC_API_USERS}/get-token`,
          {
            email: process.env.NEXT_PUBLIC_EMAIL_USERS,
            password: process.env.NEXT_PUBLIC_PASSWORD_USERS
          }
        )
        const responseGetEvents: { data: Event[] } = await axios.get(
          `${process.env.NEXT_PUBLIC_API_EVENTS}`,
          {
            headers: {
              Authorization: 'Bearer ' + responseGetToken.data.token
            }
          }
        )
        setEvents(responseGetEvents.data)
      } catch (error) {
        setError('sorry, we had a problem, please try again later')
      }
    }
    void getEvents()
  }, [])

  function cardEvent (id: string): React.ReactNode {
    const event = events.find((e) => e.id === id)
    if (event !== undefined) {
      setEvent(event)
      return (
        <div className={styles.card}>
          <img src={event.image_url} className={styles.imageUrl} />
          <div className={styles.info}>
            <p>{event.description}</p>
            <br />
            <p className={styles.price}>
              {new Intl.NumberFormat(navigator.language, {
                style: 'currency',
                currency: event.currency
              }).format(event.price)}
            </p>
            <br />
            <p>
              <b>Date:</b>
              {new Date(event.event_date).toLocaleDateString(
                navigator.language
              )}
            </p>
          </div>
          <button
            className={styles.btnNext}
            onClick={() => {
              setNextPage(nextPage + 1)
            }}
          >
            Next <FiArrowRight />
          </button>
        </div>
      )
    }
  }

  return (
    <div className={styles.events}>
      {events.length > 0 &&
        (
        <select
          value={selectEventId}
          onChange={(e) => {
            setSelectEventId(e.target.value)
          }}
        >
          <option value="0">Select event:</option>
          {events.map((item) => (
            <option key={item.id} value={item.id}>
              {item.name}
            </option>
          ))}
        </select>
        )}
      {selectEventId !== '' && cardEvent(selectEventId)}
      <span className={styles.errorMessage}>{error}</span>
    </div>
  )
}
