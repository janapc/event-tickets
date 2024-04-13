'use client'
import { useState, type FormEvent, useContext } from 'react'
import { FiArrowRight } from 'react-icons/fi'
import axios from 'axios'
import { FormContext } from '@/app/context/form.context'
import styles from './style.module.css'

export default function RegisterLead (): React.ReactNode {
  const [error, setError] = useState('')
  const { email, setEmail, nextPage, setNextPage } = useContext(FormContext)

  async function handleSubmit (e: FormEvent<HTMLFormElement>): Promise<void> {
    e.preventDefault()
    setError('')
    if (email === '') {
      setError('the email is mandatory')
      return
    }
    try {
      await axios.post(process.env.NEXT_PUBLIC_API_LEAD, {
        email,
        language: navigator.language,
        converted: false
      })
      setNextPage(nextPage + 1)
    } catch (e) {
      setError('sorry, we had a problem, please try again later')
    }
    setError('')
  }

  return (
    <div className={styles.marketing}>
      <form onSubmit={handleSubmit}>
        <input
          className={styles.inputEmail}
          type="email"
          name="email"
          id="email"
          placeholder="E-mail"
          value={email}
          onChange={(e) => {
            setEmail(e.target.value)
          }}
        />
        <button className={styles.btnSubmit} type="submit">
          Next <FiArrowRight />
        </button>
      </form>
      <span className={styles.errorMessage}>{error}</span>
    </div>
  )
}
