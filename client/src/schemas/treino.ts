import { z } from "zod"

export const exercicioSchema = z.object({
    id: z.number().int(),
    nome: z.string(),
    series: z.number().int()
})

export const treinoSchema = z.object({
    nome: z.string(),
    exercicios: z.array(exercicioSchema)
})

export type exercicioDTO = z.infer<typeof exercicioSchema>
export type treinoDTO = z.infer<typeof treinoSchema>
    