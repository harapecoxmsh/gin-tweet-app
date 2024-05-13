import React from 'react';
import PostData from '../types/type';
import ContentCard from './Card';

interface AllPostDataProps {
    AllPostData: PostData[];
}


const CardList = ({AllPostData}: AllPostDataProps) => {
  return (
    <div>
        {AllPostData.map((post: PostData) => (
            <ContentCard key={post.id} post={post}/>
        ))}

    </div>
  )
}

export default CardList


