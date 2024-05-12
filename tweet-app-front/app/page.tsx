'use client'
import {
  Card,
  CardContent,
  CardFooter,

} from "@/components/ui/card"
import CardList from "./components/CardList";
import { useState, useEffect } from "react";
import axios from "axios";
import PostData from "./types/type";



export default function Home() {
  const [posts, setPosts] = useState([])

  const API_URL = process.env.API_URL
  useEffect(() => {
    getPosts();
  }, []);

  const getPosts = () => {
    axios.get(`${API_URL}/posts`)
    .then((res) => res.data)
    .then((data) => {
        setPosts(data);
    })
    .catch((err) => alert(err));
  };


  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <CardList posts={posts}/>
    </main>
  );
}
