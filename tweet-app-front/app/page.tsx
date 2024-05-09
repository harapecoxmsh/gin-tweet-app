
import {
  Card,
  CardContent,
  CardFooter,

} from "@/components/ui/card"
import CardList from "./components/CardList";


export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <CardList />
    </main>
  );
}
