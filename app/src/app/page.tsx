'use client'
import { useContext } from 'react'
import { FiTag } from 'react-icons/fi'
import RegisterMarketing from './form/registerLead/register-lead'
import ListEvents from './form/listEvents/list-events'
import Payment from './form/payment/payment'
import PaymentProgress from './form/payment-progress/payment-progress'
import { FormContext } from './context/form.context'
import styles from './page.module.css'

export default function Home (): React.ReactNode {
  const { nextPage } = useContext(FormContext)

  function pageSelected (pageNumber: number): React.ReactNode {
    switch (pageNumber) {
      case 0:
        return <RegisterMarketing />

      case 1:
        return <ListEvents />

      case 2:
        return <Payment />

      case 3:
        return <PaymentProgress />

      default:
        return <span>OI</span>
    }
  }

  return (
    <main className={styles.main}>
      <header>
        <FiTag />
        <h3>Buy Ticket</h3>
      </header>
      {pageSelected(nextPage)}
    </main>
  )
}
