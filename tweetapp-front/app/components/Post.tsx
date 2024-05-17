import React from 'react'
import { PostData } from '../types/type'

interface PostDataProps {
    post: PostData
}
const Post = ({ post }: PostDataProps) => {
  return (
    <div>
      {post.Content}
    </div>
  )
}

export default Post
