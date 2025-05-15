import { FetchResponseType } from "@/models/response-models";
import { treinoDTO } from "@/schemas/treino";

export const fetchCreateTreino = async (treino: treinoDTO) => {
    const response = await fetch("http://localhost:8080/treinos/", {
        method: "POST",
        headers: {
            'content-type': 'application/json'
        },
        credentials: "include",
        body: JSON.stringify(treino)
    })

    const data: FetchResponseType<null> = await response.json()

    return data

}