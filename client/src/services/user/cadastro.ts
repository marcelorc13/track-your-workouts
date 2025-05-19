import { FetchResponseType } from "@/models/response-models"
import { cadastroUsuarioDTO } from "@/schemas/usuarios"

export const fetchCadastro = async (usuario: cadastroUsuarioDTO): Promise<FetchResponseType<null>> =>  {
    const response = await fetch("http://localhost:8080/usuarios/", {
        method: "POST", 
        headers: {
            'content-type': 'application/json'
        },
        credentials: "include",
        body: JSON.stringify(usuario)
    })

    const data: FetchResponseType<null> = await response.json()

    return data

}