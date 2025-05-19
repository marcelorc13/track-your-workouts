'use client'

import { UseRedirect } from "@/hooks/useRedirect";
import { UseReload } from "@/hooks/useReload";
import { cadastroUsuarioDTO, cadastroUsuarioSchema } from "@/schemas/usuarios";
import { fetchCadastro } from "@/services/user/cadastro";
import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import toast from "react-hot-toast";
import { LuEye, LuEyeClosed } from "react-icons/lu";

const CadastroComponent: React.FC = () => {
    const [verSenha, setVerSenha] = useState<boolean>(false)
    const [verConfirmarSenha, setVerConfirmarSenha] = useState<boolean>(false)
    const { register, handleSubmit } = useForm<cadastroUsuarioDTO>()

    const { mutate, isPending } = useMutation({
        mutationFn: fetchCadastro,
        onSuccess: (data) => {
            if (data.status != 201) {
                console.log(data.message)
                return
            }
            toast.success(data.message)
            UseRedirect('/login')
        },
        onError: (error) => {
            console.error('Erro no login', error);
        }
    })

    const handleCadastro: SubmitHandler<cadastroUsuarioDTO> = (usuario: cadastroUsuarioDTO) => {
        const result = cadastroUsuarioSchema.safeParse(usuario)
        if (!result.success) {
            result.error.issues.forEach((err) => {
                toast.error(err.message)
                return
            })
        }
        if (usuario.senha != usuario.confirmarSenha) {
            toast.error("As senhas não conferem")
            return
        }

        mutate({ nome_completo: usuario.nome_completo, username: usuario.username, email: usuario.email, senha: usuario.senha })
    }

    return (
        <main>
            <h2>Cadastro</h2>
            <form onSubmit={handleSubmit(handleCadastro)} className="flex flex-col gap-2">
                <input {...register("nome_completo")} type="text" name="nomeCompleto" id="nomeCompleto" placeholder="Nome Completo" />
                <input {...register("username")} type="text" name="username" id="username" placeholder="Username" />
                <input {...register("email")} type="email" name="email" id="email " placeholder="Email" />
                <input {...register("senha")} type={!verSenha ? "password" : "text"} name="senha" id="senha" placeholder="Senha" />
                <span className='cursor-pointer' onClick={() => setVerSenha(prev => !prev)}>{verSenha ? <LuEye /> : <LuEyeClosed />}</span>
                <input {...register("confirmarSenha")} type={!verConfirmarSenha ? "password" : "text"} name="confirmarSenha" id="confirmarSenha" placeholder="Confirmar Senha" />
                <span className='cursor-pointer' onClick={() => setVerConfirmarSenha(prev => !prev)}>{verConfirmarSenha ? <LuEye /> : <LuEyeClosed />}</span>
                <input type="submit" value={!isPending ? "Entrar" : "..."} />
            </form>
        </main>
    );
};

export default CadastroComponent;