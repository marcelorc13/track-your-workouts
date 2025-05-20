'use server'

import { authTokenPayload } from '@/models/token-models'
import * as jose from 'jose'
import { cookies } from 'next/headers'

const getToken = async (): Promise<string | undefined> => {
    return (await cookies()).get("token")?.value
}

export const verifyToken = async (): Promise<boolean> => {
    const token: string | undefined = await getToken()
    const segredo = new TextEncoder().encode(process.env.NEXT_PUBLIC_JWT_SECRET) || ""

    if (token === undefined) {
        return false
    }

    try {
        await jose.jwtVerify(token, segredo)
        return true
    } catch (err) {
        console.error(err)
        return false
    }
}

export const getTokenUserID = async (): Promise<string> => {
    const token: string | undefined = await getToken()

    if (token === undefined) {
        return ""
    }

    const segredo = new TextEncoder().encode(process.env.NEXT_PUBLIC_JWT_SECRET) || ""
    const { payload } = await jose.jwtVerify<authTokenPayload>(token, segredo)

    return payload.id
}