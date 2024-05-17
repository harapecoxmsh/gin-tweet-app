import React from 'react'
import { PostData } from '../types/type'
import Post from './Post'


interface PostAllDataProps {
    postAllData: PostData[]
}
const PostList = ({ postAllData }: PostAllDataProps) => {
  return (
    <div>
      {postAllData.map((post) => (
        <Post post={post} key={post.ID} />
      ))}
    </div>
  )
}

export default PostList
