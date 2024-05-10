import React, { useState } from 'react'
import {
    Card,
    CardContent,
    CardFooter,

  } from "@/components/ui/card"
import ContentCard from './Card'
import PostData from '../types/type'

interface PostAllDataProps {
    posts: PostData[]
}

const CardList = ({posts}: PostAllDataProps) => {

  return (
    <div>
    {posts.map((post: PostData) => (
        <CardContent key={post.id} post={post}/>
    ))}
    </div>
  )
}

export default CardList
