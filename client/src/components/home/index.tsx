'use client'

import { fecthGetUsuario } from "@/services/user/usuario";
import { useQuery } from "@tanstack/react-query";

const HomeClient: React.FC = () => {
    const { data, isPending } = useQuery({ queryKey: ['get-usuario'], queryFn: fecthGetUsuario })

    return (
        <div>
            {isPending ? <h1>...</h1> : <h1>Seja bem-vindo {data?.data?.nome_completo.split(" ")[0]}</h1>}
        </div>
    );
};

export default HomeClient;