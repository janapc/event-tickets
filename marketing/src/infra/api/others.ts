import { type RequestHandler, Router } from 'express'
import { register } from 'prom-client'

const router = Router()

router.get('/metrics', (async (req, res): Promise<void> => {
  /**
   * #swagger.ignore = true
   */
  const metrics = await register.metrics()
  res.setHeader('Content-Type', register.contentType).send(metrics)
}) as RequestHandler)

router.get('/healthcheck', (async (req, res): Promise<void> => {
  /**
   * #swagger.ignore = true
   */
  const healthcheck = {
    uptime: process.uptime(),
    message: 'OK',
    timestamp: Date.now(),
  }
  try {
    res.status(200).json(healthcheck)
  } catch (error) {
    healthcheck.message = error as string
    res.status(503)
  }
}) as RequestHandler)

export default router
