import express from 'express'
import morgan from 'morgan'
import bodyParser from 'body-parser'
import router from './routes'

const app = express()
app.use(bodyParser.json())
app.use(morgan('tiny'))
app.use('/leads', router)

export function server(): void {
  app.listen(process.env.PORT, () => {
    console.info('\x1b[32m', `Server running in port ${process.env.PORT}`)
  })
}
