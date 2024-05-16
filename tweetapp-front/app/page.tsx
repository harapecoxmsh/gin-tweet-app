'use client'
import Image from "next/image";
import Postlist from "./components/Postlist";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
     <Postlist />
    </main>
  );
}
