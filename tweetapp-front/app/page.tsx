'use client'
import { useState, useEffect } from "react";
import axios from "axios";
import { PostData } from "./types/type";
import PostList from "./components/PostList";
import React, { FormEvent } from 'react';

export default function Home() {


  const [postAllData, setPostAllData] = useState<PostData[]>([]);
  const [postContent, setPostcontent] = useState("");

  useEffect(() => {
    fetchPostData();
  }, []);

async function fetchPostData() {
    try {
      const response = await axios.get(`${process.env.API_BASE_URL}/posts`);
      const postDataArray: PostData[] = Object.values(response.data.data);
      console.log(postDataArray);
      setPostAllData(postDataArray);
    } catch (error) {
      console.error("Error fetching posts:", error);
    }
  }

  const createPost = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
     const postData = {
       Content: postContent
   };

    axios.post(`${process.env.API_BASE_URL}/posts`, postData)
    .catch((error) => alert(error));
    fetchPostData();
  }


  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24 bg-cyan-400	background-color: rgb(34 211 238);">
        <PostList postAllData={postAllData} />
        <footer>
          <form onSubmit={createPost}>
            <input type="text" value={postContent} onChange={(e) => setPostcontent(e.target.value) } />
            <input className="mx-1	margin-left: 0.25rem" type="submit" value="投稿" />
          </form>
        </footer>
    </main>
  );
}
