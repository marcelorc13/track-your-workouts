import { FetchResponseType } from "@/models/response-models"
import { treinoDTO } from "@/schemas/treino"
import { getTokenUserID } from "@/utils/jwt"

export const fetchGetTreinos = async () => {
    const id = await getTokenUserID()
    if (id === "") {
        return { status: 400, message: "Token não encontrado" }
    }

    const response = await fetch(`http://localhost:8080/treinos/usuario/${id}`, {
        method: "GET",
        headers: {
            'content-type': 'application/json'
        },
        credentials: "include",
    })

    const data: FetchResponseType<treinoDTO[]> = await response.json()

    return data
}

export const fetchGetTreino = async (id: string) => {
    const response = await fetch(`http://localhost:8080/treinos/${id}`, {
        method: "GET",
        headers: {
            'content-type': 'application/json'
        },
        credentials: "include",
    })

    const data: FetchResponseType<treinoDTO> = await response.json()

    return data
}