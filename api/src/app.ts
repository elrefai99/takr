import "./core/dotenv"
import { logger } from './utils/logger'
process
     .on('unhandledRejection', (reason, promise) => {
          console.error(reason, 'Unhandled Rejection at Promise', promise)
          logger.error({
               message: 'Unhandled Rejection',
               reason: reason instanceof Error ? reason.message : reason,
               stack: reason instanceof Error ? reason.stack : undefined,
               promise,
          })
     })
     .on('uncaughtException', (err) => {
          console.error(err, '\n Uncaught Exception thrown \n')
          logger.error({
               message: 'Uncaught Exception',
               error: err.message,
               stack: err.stack,
          })

          // allow logger to flush before exit
          setTimeout(() => {
               process.exit(1)
          }, 500)
     })
import express, { Application } from "express";
import { mongoDBConfig } from "./core/mongo";

const app: Application = express()

const startServer = async () => {
     try {
          mongoDBConfig().then(
               async () => {
                    app.listen(process.env.PORT, () => {
                         console.log('🌐 Server is running on:', process.env.API_LINK as string)
                    })
               },
          ).catch((err) => {
               logger.error({
                    message: 'MongoDB connection failed',
                    error: err.message,
                    stack: err.stack,
               })
               process.exit(1)
          })
     }
     catch (err) {
          console.error('Error starting server:', err);
     }
}
startServer()
