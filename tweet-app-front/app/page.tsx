
import {
  Card,
  CardContent,
  CardFooter,

} from "@/components/ui/card"

import PostData from "./types/type";
import CardList from "./components/CardList";


const API_URL = process.env.API_URL

async function getAllPosts() {
  const response = await fetch('http://localhost:8080/posts',{
    cache: "no-store",
  });
  const AllPostData: PostData[] = await response.json();
  console.log(AllPostData)
  return AllPostData
}

export default async function Home() {
  const AllPostData = await getAllPosts();
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <CardList AllPostData={AllPostData}/>
    </main>
  );

  }
