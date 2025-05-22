'use client'

import { fetchGetUsuario } from "@/services/user/usuario";
import { useQuery } from "@tanstack/react-query";

const HomeClient: React.FC = () => {
    const { data, isPending } = useQuery({ queryKey: ['get-usuario'], queryFn: fetchGetUsuario })

    return (
        <div>
            {isPending ? <h1>...</h1> : (
                <div>
                    <h1>Seja bem-vindo {data?.data ? data?.data?.nome_completo.split(" ")[0] : "..."}</h1>
                </div>
            )}
            
        </div>
    );
};

export default HomeClient;