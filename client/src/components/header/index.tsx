'use client'

import { UseReload } from "@/hooks/useReload";
import { logout } from "@/services/logout";
import Link from "next/link";

const Header: React.FC = () => {

    const handleLogout = async () => {
        try {
            await logout()
            UseReload()
        } catch (err) {
            console.log(err)
        }
    }
    return (
        <header className="w-full flex justify-between border-b border-gray-500 py-4">
            <ul className="flex justify-center gap-4">
                <li><Link href={'/'}>Home</Link></li>
                <li><Link href={'/criar-treino'}>Treinos</Link></li>
            </ul>
            <span className="cursor-pointer" onClick={handleLogout}>logout</span>
        </header>
    );
};

export default Header;