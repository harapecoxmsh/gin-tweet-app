import React from 'react';
import { PostData } from '../types/type';

interface PostDataProps {
    post: PostData;
    onDelete: (id: number) => void;
}

const Post = ({ post, onDelete }: PostDataProps) => {
    const handleDelete = () => {
        onDelete(post.ID);
    };

    return (
        <div className='mb-3 margin-bottom: 0.75rem;'>
            {post.Content}
            <button className=
            "bg-red-500 hover:bg-red-700 text-white py-1 px-2 rounded mx-1	margin-left: 0.25rem"
            onClick={handleDelete}>消去</button>
        </div>
    );
};

export default Post;

