import type { Metadata } from "next";
import "./globals.css";
import ReactQueryProvider from "@/utils/reactQueryProvider";
import { Toaster } from "react-hot-toast";

export const metadata: Metadata = {
  title: "Track Your Workouts"
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="pt-BR">
      <body>
        <ReactQueryProvider>
          <Toaster/>
          {children}
        </ReactQueryProvider>
      </body>
    </html>
  );
}
