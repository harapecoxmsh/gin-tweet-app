import React from 'react';
import PostData from '../types/type';
import ContentCard from './Card'; // CardContentのインポートが必要かもしれません

interface PostAllDataProps {
    posts: PostData[];
}

const CardList: React.FC<PostAllDataProps> = (props) => {
    const { posts } = props;

    return (
        <div>
            {posts.map((post: PostData) => (
                <ContentCard key={post.id} post={post}/>
            ))}
        </div>
    );
}

export default CardList;

