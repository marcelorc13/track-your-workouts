'use client'

import Link from "next/link";

const Header: React.FC = () => {
    return (
        <header className="w-full border-b border-gray-500">
            <ul className="flex justify-center gap-4">
                <li><Link href={'/'}>Home</Link></li>
                <li><Link href={'/criar-treino'}>Treinos</Link></li>
            </ul>
        </header>
    );
};

export default Header;