import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import { FormProvider } from './context/form.context'
import './globals.css'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'App',
  description: 'buy tickets'
}

export default function RootLayout ({
  children
}: Readonly<{
  children: React.ReactNode
}>): React.ReactNode {
  return (
    <html lang="en">
      <body className={inter.className}>
        <FormProvider>{children}</FormProvider>
      </body>
    </html>
  )
}
