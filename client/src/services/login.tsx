import { FetchResponseType } from "@/models/response-models"

export const fetchLogin = async (usuario: {email: string, senha: string}) =>  {
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