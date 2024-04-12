'use client'
import axios from 'axios'
import { type ReactNode, createContext, useState, useEffect } from 'react'

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

export interface FormContextData {
  nextPage: number
  email: string
  event: Event | null
  setEvent: (event: Event) => void
  setEmail: (email: string) => void
  setNextPage: (pageNumber: number) => void
}

export const FormContext = createContext({} as FormContextData)

export function FormProvider ({
  children
}: {
  children: ReactNode
}): React.ReactNode {
  const [nextPage, setNextPage] = useState<number>(0)
  const [email, setEmail] = useState<string>('')
  const [event, setEvent] = useState<Event | null>(null)

  return (
    <FormContext.Provider
      value={{ event, nextPage, email, setEvent, setEmail, setNextPage }}
    >
      {children}
    </FormContext.Provider>
  )
}
