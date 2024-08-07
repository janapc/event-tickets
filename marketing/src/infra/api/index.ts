import express from 'express'
import type * as http from 'http'
import morgan from 'morgan'
import bodyParser from 'body-parser'
import router from './routes'
import { logger } from '@infra/logger/logger'
import cors from 'cors'

const app = express()
app.use(bodyParser.json())
app.use(morgan('tiny'))
app.use(cors())
app.use('/leads', router)

export function server(): http.Server {
  return app.listen(process.env.PORT, () => {
    logger.info(`Server running in port ${process.env.PORT}`)
  })
}
