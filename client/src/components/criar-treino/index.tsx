'use client'

import { UseReload } from "@/hooks/useReload";
import { treinoDTO, treinoSchema } from "@/schemas/treino";
import { fetchCreateTreino } from "@/services/treino/criar-treino";
import { useMutation } from "@tanstack/react-query";
import { useState } from "react";
import toast from "react-hot-toast";
import Treinos from "../treinos";

const CriarTreinoComponent: React.FC = () => {

    const [treino, setTreino] = useState<treinoDTO>({
        nome: '',
        exercicios: []
    })

    const { mutate, isPending } = useMutation({
        mutationFn: fetchCreateTreino,
        onSuccess: (data) => {
            if (data.status != 201) {
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

    const addExercicio = () => {
        const len: number = treino.exercicios.length
        setTreino((prev) => ({ ...prev, exercicios: [...prev.exercicios, { id: len + 1, nome: '', series: 0 }] }))
    }

    const handleChange = (id: number | null, field: string, value: string | number) => {
        if (id === null) {
            setTreino(prev => ({ ...prev, [field]: value }))
            return
        }

        setTreino(prevTreino => ({
            ...prevTreino, exercicios: prevTreino.exercicios.map(exercicio => exercicio.id === id ?
                { ...exercicio, [field]: field === 'series' ? parseInt(value as string) || 0 : value }
                : exercicio
            )
        }));
    }

    const handleLogin = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        const validData = treinoSchema.safeParse(treino)
        if (!validData.success) {
            validData.error.issues.forEach((err) => {
                toast.error(err.message)
            })
            return
        }

        mutate(treino)
    }

    return (
        <main className="flex flex-col gap-4">
            <Treinos />
            <form onSubmit={handleLogin} className="flex flex-col">
                <h2 className="text-xl font-semibold">Criar Treino</h2>
                <input onChange={(e) => handleChange(null, 'nome', e.target.value)} type="text" name="nome" id="nome" placeholder="nome" />
                {treino.exercicios.map(exercicio => (
                    <div key={exercicio.id}>
                        <h3>Exercicio {exercicio.id}</h3>
                        <input onChange={(e) => handleChange(exercicio.id, 'nome', e.target.value)} type="text" placeholder="Nome do Exercício" />
                        <input onChange={(e) => handleChange(exercicio.id, 'series', e.target.value)} type="number" placeholder="N de Séries" min={1} />
                    </div>
                ))}
                <button type="button" onClick={addExercicio}>Adicionar exercicio</button>
                <input type="submit" value={!isPending ? "Criar Treino" : "..."} />
            </form>
        </main>
    )
}

export default CriarTreinoComponent;