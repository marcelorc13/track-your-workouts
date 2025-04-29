import { z } from "zod"

export const loginUsuarioSchema = z.object({
    email: z.string().email().min(8).max(256),
    senha: z.string().min(6).max(30)
})
export const cadastroUsuarioSchema = z.object({
    nome_completo: z.string().min(6, 'O nome deve conter pelo menos 6 caracteres').max(100, 'O nome deve conter no máximo 100 caracteres'),
    username: z.string().min(3, 'O username deve contter pelo menos 3 caracteres').max(30, 'O username deve conter no máximo 30 caracteres'),
    email: z.string().email("Email inválido").max(256, 'O email deve conter no máximo 256 caracteres'),
    senha: z.string().min(6, 'A senha deve conter pelo menos 6 caracteres').max(30, 'A senha deve conter no máximo 30 caracteres'),
    confirmarSenha: z.string().min(6, 'A senha deve conter pelo menos 6 caracteres').max(30, 'A senha deve conter no máximo 30 caracteres').nullish(),
})

export type loginUsuarioDTO = z.infer<typeof loginUsuarioSchema>
export type cadastroUsuarioDTO = z.infer<typeof cadastroUsuarioSchema>
