'use client'
import { FiCheckCircle } from 'react-icons/fi'
import styles from './style.module.css'
import { useContext } from 'react'
import { FormContext } from '@/app/context/form.context'

export default function PaymentProgress (): React.ReactNode {
  const { email } = useContext(FormContext)
  return (
    <div className={styles.paymentProgress}>
      <FiCheckCircle size={42} />
      <p>Payment in progress...</p> <br />
      <p>You will receive an email with more information.</p>
      <p>It will arrive in this email {email}</p>
    </div>
  )
}
