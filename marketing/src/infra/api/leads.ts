import { type RequestHandler, Router } from 'express'
import swaggerUi from 'swagger-ui-express'
import swaggerOutput from '../docs/swagger_output.json'
import { GetLeads } from '@application/get_leads'
import { LeadPrismaRepository } from '@infra/dabatase/lead_prisma_repository'
import Database from '@infra/dabatase'
import { CreateLead, type InputCreateLead } from '@application/create_lead'
import { GetByEmail } from '@application/get_by_email'
import { requestTotalCounter } from '@infra/metrics/prometheus_metrics'

const router = Router()
const database = Database.getInstance()
const repository = new LeadPrismaRepository(database.connection)

router.get('/', (async (req, res): Promise<void> => {
  /*  #swagger.responses[200] = {
             schema: { $ref: "#/definitions/getLeadsResponse" }
  } */
  /* #swagger.responses[500] = { message: 'internal server error' }
   */
  try {
    const getLeads = new GetLeads(repository)
    const result = await getLeads.execute()
    res.status(200).json(result)
    requestTotalCounter
      .labels({
        route: req.url,
        method: req.method,
        statusCode: 200,
      })
      .inc()
  } catch (e) {
    res.status(500).json({ message: 'internal server error' })
  }
}) as RequestHandler)

router.get('/search', (async (req, res): Promise<void> => {
  /*  #swagger.responses[200] = {
             schema: { $ref: "#/definitions/getByEmailResponse" }
  } */
  /* #swagger.parameters['email'] = {
        in: 'query',                            
        description: 'email of lead',                   
        required: true
} */
  /* #swagger.responses[500] = { message: 'internal server error' }
   */

  /* #swagger.responses[404] = { message: 'lead is not found' }
   */

  try {
    const { email } = req.query as { email: string }
    const application = new GetByEmail(repository)
    const result = await application.execute(email)
    res.status(200).json(result)
    requestTotalCounter
      .labels({
        route: req.url,
        method: req.method,
        statusCode: 200,
      })
      .inc()
  } catch (e) {
    if (e instanceof Error && e.message === 'lead is not found') {
      res.status(404).json({ message: e.message })
    } else {
      res.status(500).json({ message: 'internal server error' })
    }
  }
}) as RequestHandler)

router.post('/', (async (req, res): Promise<void> => {
  /*  #swagger.requestBody = {
            required: true,
            schema: { $ref: "#/definitions/createLeadRequest" }
    } */

  /*  #swagger.responses[201] = {
             schema: { $ref: "#/definitions/lead" }
  } */
  /*
	#swagger.responses[500] = { message: 'internal server error' }
	 */
  try {
    const createLead = new CreateLead(repository)
    const body = req.body as InputCreateLead
    const result = await createLead.execute(body)
    res.status(201).json(result)
    requestTotalCounter
      .labels({
        route: req.url,
        method: req.method,
        statusCode: 201,
      })
      .inc()
  } catch (e) {
    res.status(500).json({ message: 'internal server error' })
  }
}) as RequestHandler)

router.use('/docs', swaggerUi.serve, swaggerUi.setup(swaggerOutput))

export default router
