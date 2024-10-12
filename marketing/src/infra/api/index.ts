import express from 'express'
import morgan from 'morgan'
import type * as http from 'http'
import bodyParser from 'body-parser'
import others from './others'
import leads from './leads'
import { logger } from '@infra/logger/logger'
import cors from 'cors'

const app = express()
app.use(bodyParser.json())
app.use(morgan('tiny'))
app.use(cors())
app.use('/', others)
app.use('/leads', leads)

let server: http.Server

export function init(): void {
  server = app.listen(process.env.PORT, () => {
    logger.info(`HTTP server running in port ${process.env.PORT}`)
  })
}

export function close(): void {
  server.close((error) => {
    if (error) logger.error(`HTTP server close error ${String(error)}`)
    logger.info('HTTP server closed')
  })
}
