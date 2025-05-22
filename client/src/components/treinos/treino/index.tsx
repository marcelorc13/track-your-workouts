'use client'

import { fetchGetTreino } from "@/services/treino/treinos";
import { useQuery } from "@tanstack/react-query";
import Link from "next/link";

interface Props {
    id: string
}

const TreinoSelecionado: React.FC<Props> = ({ id }) => {
    const { data, isPending } = useQuery({ queryKey: ['get-usuario'], queryFn: () => fetchGetTreino(id) })

    return (
        <main>
            <Link href={"/treinos"}>voltar</Link>
            {isPending ? <span>...</span> : (
                <div>
                    <h2 className="text-xl font-semibold">{data?.data?.nome}</h2>
                    {data?.data?.exercicios.map((exercicio, key) => (
                        <div key={key} className="flex gap-4">
                            <p>{exercicio.nome} </p>
                            <p>{exercicio.series} {exercicio.series == 1 ? "Série" : "Séries"}</p>
                        </div>
                    ))}
                </div>
            )}
        </main>
    );
};

export default TreinoSelecionado;