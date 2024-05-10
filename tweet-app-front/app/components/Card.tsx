import React from 'react'
import {
    Card,
    CardContent,
    CardFooter,
  } from "@/components/ui/card"
  import PostData from '../types/type'

interface PostDataProps {
    post: PostData;
}
const ContentCard = ({post}: PostDataProps) => {
    
  return (
    <div>
    <Card>
        <CardContent>
            <p>{post.content}</p>
        </CardContent>
    </Card>
    </div>
  )
}

export default ContentCard
