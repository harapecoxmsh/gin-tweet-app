'use client'
import { useState, useEffect } from "react";
import axios from "axios";
import { PostData } from "./types/type";
import PostList from "./components/PostList";

export default function Home() {
  const [postAllData, setPostAllData] = useState<PostData[]>([]);

  useEffect(() => {
    async function fetchPostData() {
      try {
        const response = await axios.get("http://localhost:8080/posts");
        const postDataArray: PostData[] = Object.values(response.data.data); 
        console.log(postDataArray);
        setPostAllData(postDataArray);
      } catch (error) {
        console.error("Error fetching posts:", error);
      }
    }

    fetchPostData();
  }, []);
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
        <PostList postAllData={postAllData}/>
    </main>
  );
}
