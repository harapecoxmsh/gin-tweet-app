import React from 'react'
import { useState, useEffect } from 'react'
import axios from 'axios'
import { PostData } from '../types/type'

const Postlist = () => {
    const [posts, setPosts] = useState<PostData[]>([])

    useEffect(() => {
        getPost();
    }, []);

    const getPost = () => {
        axios.get("/posts")
        .then((res) => res.data)
        .then((data) => {
            setPosts(data);
            console.log(data);
        })
    }
  return (
    <div>
        {posts.map((post: PostData) => (
            <ul key={post.id}>
                <li>{post.content}</li>
            </ul>
        ))}
    </div>
  )
}

export default Postlist
