'use client'

import { UseReload } from '@/hooks/useReload';
import { loginUsuarioDTO, loginUsuarioSchema } from '@/schemas/usuarios';
import { fetchLogin } from '@/services/login';
import { useMutation } from '@tanstack/react-query';
import { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { LuEye, LuEyeClosed } from "react-icons/lu";
import toast from 'react-hot-toast';

const LoginComponent: React.FC = () => {
    const [verSenha, setVerSenha] = useState<boolean>(false)
    const { register, handleSubmit } = useForm<loginUsuarioDTO>()

    const { mutate, isPending } = useMutation({
        mutationFn: fetchLogin,
        onSuccess: (data) => {
            if (data.status != 200) {
                toast.error(data.message)
                return
            }
            toast.success(data.message)
            UseReload()
        },
        onError: (error) => {
            console.error('Erro no login', error);
        }
    })

    const handleLogin: SubmitHandler<loginUsuarioDTO> = (usuario: loginUsuarioDTO) => {
        const isValid = loginUsuarioSchema.safeParse(usuario)
        if (!isValid.success) {
            toast.error("Usuário ou senha incorreta")
            return
        }
        mutate({ email: usuario.email, senha: usuario.senha })
    }

    return (
        <main>
            <h2>Login</h2>
            <form onSubmit={handleSubmit(handleLogin)}>
                <input {...register("email")} type="email" placeholder='Email' />
                <input {...register("senha")} type={!verSenha ? "password" : "text"} placeholder='Senha' />
                <span className='cursor-pointer' onClick={() => setVerSenha(prev => !prev)}>{verSenha ? <LuEye/> : <LuEyeClosed/>}</span>
                <input type="submit" value={!isPending ? "Entrar" : "..."} />
            </form>
        </main>
    );
};

export default LoginComponent;