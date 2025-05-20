import { FetchResponseType } from "@/models/response-models"
import { usuarioDTO } from "@/models/usuario-models"
import { getTokenUserID } from "@/utils/jwt"

export const fecthGetUsuario = async (): Promise<FetchResponseType<usuarioDTO>> => {
    const id = await getTokenUserID()
    if (id === "") {
        return {status: 400, message: "Token não encontrado"}
    }

    const response = await fetch(`http://localhost:8080/usuarios/${id}`, {
        method: "GET",
        headers: {
            'content-type': 'application/json'
        },
        credentials: "include",
    })

    const data: FetchResponseType<usuarioDTO> = await response.json()

    return data

}