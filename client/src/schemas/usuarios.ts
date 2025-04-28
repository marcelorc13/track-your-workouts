import { z } from "zod"

export const loginUsuarioSchema = z.object({
    email: z.string().email().min(8).max(256),
    senha: z.string().min(6).max(30)
})

export type loginUsuarioDTO = z.infer<typeof loginUsuarioSchema>
