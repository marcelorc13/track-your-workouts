'use client'

import { fetchGetTreinos } from "@/services/treino/treinos";
import { useQuery } from "@tanstack/react-query";
import Link from "next/link";

const Treinos: React.FC= () => {
    const { data, isPending } = useQuery({ queryKey: ['get-treinos'], queryFn: fetchGetTreinos })

    return (
        <section>
                <h2 className="text-xl font-semibold">Meus treinos:</h2>
                {isPending ? <div>...</div> :data?.data?.map((treino) => (
                    <div key={treino._id}>
                        <Link href={`treinos/${treino._id}`}>{treino.nome}</Link>
                    </div>
                ))}
            </section>
    )
}

export default Treinos;