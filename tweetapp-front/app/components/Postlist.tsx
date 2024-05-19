import React from 'react'
import { PostData } from '../types/type'
import Post from './Post'
import axios from 'axios'


interface PostAllDataProps {
    postAllData: PostData[];
}

const deletePost = (id: number) => {
  axios.delete(`${process.env.API_BASE_URL}/posts/${id}`)
  .catch((error) => alert(error));
}

const PostList = ({ postAllData }: PostAllDataProps) => {
  return (
    <div>
      {postAllData.map((post) => (
        <Post post={post} key={post.ID} onDelete={deletePost} />
      ))}
    </div>
  )
}

export default PostList
