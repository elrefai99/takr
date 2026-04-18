import express, { type Application, type NextFunction, type Request, type Response } from 'express'
import cors from 'cors'
import cookieParser from 'cookie-parser'
import morgan from 'morgan'
import helmet from 'helmet'
import { createWriteStream } from 'node:fs'
import { resolve } from 'node:path'

const morganLogStream = createWriteStream(resolve(process.cwd(), 'logs', 'app.log'), { flags: 'a' })

export const allowedOrigins: string[] = [
  process.env.SITE_URL_TEST as string,
  process.env.SITE_URL_LIVE as string,
]
export default (app: Application) => {
  const corsOptions: object = {
    origin: (origin: any, callback: any) => {
      if (!origin || origin === 'null' || allowedOrigins.includes(origin)) {
        callback(null, true)
      } else {
        callback(new Error('Not allowed by CORS'))
      }
    },
    credentials: true,
    optionsSuccessState: 200,
  }

  app.use(
    helmet({
      contentSecurityPolicy: {
        directives: {
          defaultSrc: ["'self'"],
          styleSrc: ["'self'", "'unsafe-inline'"],
          scriptSrc: ["'self'"],
          imgSrc: ["'self'", 'data:', 'https:'],
          connectSrc: ["'self'", process.env.SITE_API_URL as string],
        },
      },
      hsts: {
        maxAge: 31536000, // 1 year
        includeSubDomains: true,
        preload: true,
      },
      referrerPolicy: { policy: 'strict-origin-when-cross-origin' },
      permittedCrossDomainPolicies: { permittedPolicies: 'none' },
    }),
  )
  app.use(
    express.json({
      limit: '75mb',
    }),
  )
  app.use(
    express.urlencoded({
      extended: true,
    }),
  )

  app.use('/v0/cdn', express.static('assets'))
  app.use('/v0/public', express.static('public'))
  app.use(cors(corsOptions))
  app.use(cookieParser())
  app.use(
    morgan(process.env.NODE_ENV === 'development' ? 'dev' : 'combined', {
      stream: {
        write: (message) => {
          process.stdout.write(message)
          morganLogStream.write(message)
        },
      },
    }),
  )
  app.set('trust proxy', true)

  app.use(async (req: Request | any, _res: Response, next: NextFunction) => {
    req.clientIP =
      req.headers['cf-connecting-ip'] ||
      req.headers['x-real-ip'] ||
      req.headers['x-forwarded-for'] ||
      req.socket.remoteAddress ||
      ('' as string)

    if (req.path === '/metrics') {
      return next()
    }
    next()
  })
}
