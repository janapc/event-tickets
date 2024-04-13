'use client'
import { type FormEvent, useState, useContext } from 'react'
import { FiArrowRight } from 'react-icons/fi'
import styles from './style.module.css'
import { FormContext } from '@/app/context/form.context'
import axios from 'axios'

export default function Payment (): React.ReactNode {
  const { event, setNextPage, nextPage, email } = useContext(FormContext)
  const [name, setName] = useState('')
  const [cardNumber, setCardNumber] = useState('')
  const [securityCode, setSecurityCode] = useState('')
  const [error, setError] = useState<string>('')

  async function handleSubmit (e: FormEvent<HTMLFormElement>): Promise<void> {
    e.preventDefault()
    try {
      await axios.post(process.env.NEXT_PUBLIC_API_PAYMENT, {
        name,
        email,
        event_id: event?.id,
        event_amount: event?.price,
        card_number: cardNumber,
        security_code: securityCode,
        event_name: event?.name,
        event_description: event?.description,
        event_image_url: event?.image_url,
        language: navigator.language.split('-')[0]
      })
      setNextPage(nextPage + 1)
    } catch (error) {
      setError('sorry, we had a problem, please try again later')
    }
  }

  return (
    <div className={styles.payment}>
      <form method="post" onSubmit={handleSubmit}>
        <input
          type="text"
          name="name"
          id="name"
          placeholder="Name"
          value={name}
          onChange={(e) => {
            setName(e.target.value)
          }}
        />
        <input
          type="number"
          name="cardNumber"
          id="cardNumber"
          placeholder="Card Number"
          value={cardNumber}
          onChange={(e) => {
            setCardNumber(e.target.value)
          }}
        />
        <input
          type="number"
          name="securityCode"
          id="securityCode"
          placeholder="Security Code"
          value={securityCode}
          onChange={(e) => {
            setSecurityCode(e.target.value)
          }}
        />
        {event !== null && (
          <div className={styles.card}>
            <img src={event.image_url} />
            <div className={styles.description}>
              <b>{event.name}</b>
              <p>
                {new Date(event.event_date).toLocaleDateString(
                  navigator.language
                )}
              </p>
              <b className={styles.price}>
                {new Intl.NumberFormat(navigator.language, {
                  style: 'currency',
                  currency: event.currency
                }).format(event.price)}
              </b>
            </div>
          </div>
        )}
        <span className={styles.errorMessage}>{error}</span>
        <button className={styles.btnSubmit} type="submit">
          Next <FiArrowRight />
        </button>
      </form>
    </div>
  )
}
