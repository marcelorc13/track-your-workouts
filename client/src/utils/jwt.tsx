import * as jose from 'jose'

export const verifyToken = async (token: string): Promise<boolean> => {
    const segredo = new TextEncoder().encode(process.env.NEXT_PUBLIC_JWT_SECRET) || ""

    try {
        await jose.jwtVerify(token, segredo)
        return true
    } catch (err) {
        console.error(err)
        return false
    }
}