import { FetchResponseType } from "@/models/response-models"
import { loginUsuarioDTO } from "@/schemas/usuarios"

export const fetchLogin = async (usuario: loginUsuarioDTO): Promise<FetchResponseType<null>> =>  {
    const response = await fetch("http://localhost:8080/usuarios/login", {
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