import { NextRequest, NextResponse } from "next/server"
import { verifyToken } from "./utils/jwt"

export const middleware = async (req: NextRequest) => {
    const urlAtual = req.nextUrl.pathname

    const token = req.cookies.get('token')

    if (!token && urlAtual != '/login' && urlAtual != '/cadastro') {
        return NextResponse.redirect(new URL('/login', req.url))
    }

    if ((urlAtual == '/login' || urlAtual == '/cadastro') && token) {
        const isValid = await verifyToken()
        if (!isValid) {
            req.cookies.clear()
            return
        }
        return NextResponse.redirect(new URL('/', req.url))
    }
    if (token) {
        const isValid = await verifyToken()
        if (!isValid) {
            req.cookies.clear()
            return NextResponse.redirect(new URL('/login', req.url))
        }
        return NextResponse.next()
    }
}

export const config = {
    matcher: ['/', '/:path'],
}