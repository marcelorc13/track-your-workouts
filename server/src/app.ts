import express, { Application } from "express";
import cors from "cors";
import { appRouter } from "./routes";


export const app: Application = express()

app.use(express.json())
app.use(cors())

app.use("/api", appRouter)